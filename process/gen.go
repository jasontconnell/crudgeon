package process

import (
	"bytes"
	"fmt"
	"github.com/jasontconnell/fileutil"
	"io/ioutil"
	"lpgagen/data"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var fns = template.FuncMap{
	"plus1": func(x int) int {
		return x + 1
	},
}

func Generate(pkg data.GenPackage, objdir bool) error {
	tmpl, err := template.New(pkg.TemplateFile).Funcs(fns).ParseFiles(pkg.TemplateFile)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	_, templateName := filepath.Split(pkg.TemplateFile)
	terr := tmpl.ExecuteTemplate(buffer, templateName, pkg)

	if terr != nil {
		return terr
	}

	path := pkg.Path
	if objdir {
		path = filepath.Join(path, pkg.Name)
	}

	err = fileutil.MakeDirs(path)
	if err != nil {
		return err
	}

	output := filepath.Join(path, pkg.OutputFile)
	return ioutil.WriteFile(output, buffer.Bytes(), os.ModePerm)
}

func GetGenPackage(name, path string, flds []data.Field, fileType, tmplFile, ns, prefix, folder string, flagstr string, usefieldname bool) (data.GenPackage, error) {
	flags, err := parseFlags(flagstr)
	if err != nil {
		return data.GenPackage{}, err
	}

	pkg := data.GenPackage{Name: name, Namespace: ns, Path: filepath.Join(path, folder), TemplateFile: tmplFile, Prefix: prefix, OutputFile: prefix + name + "." + fileType}
	if flags.Fields || flags.Constructor || flags.Keys || flags.Concretes {
		for _, f := range flds {
			if f.Collection && !flags.Collections {
				continue
			}

			field := f.FieldName
			cname := strings.Title(f.Name)
			if len(cname) < 3 {
				cname = strings.ToUpper(cname)
			}

			if usefieldname {
				cname = field
			}

			sqltype := getSqlType(f.Type)
			if fileType == "sql" && (sqltype == "" || f.Flags.SqlIgnore) {
				continue
			}

			if fileType == "cs" && f.Flags.CsIgnore {
				continue
			}

			isInterface := f.Type != f.ConcreteType
			typeName, concreteTypeName, elementType := f.Type, f.Type, f.Type
			if f.Collection {
				listType := "List"
				if isInterface {
					listType = "IEnumerable"
				}
				typeName = fmt.Sprintf("%s<%s>", listType, typeName)
			}
			if fileType == "sql" {
				typeName = sqltype
			}

			nullable := f.Nullable
			sqlignore := sqltype == "" || f.Collection || f.Flags.SqlIgnore

			isbase := isBaseType(typeName)
			if !isbase {
				nullable = false
			}

			jsonIgnore := f.JsonIgnore || f.Flags.JsonIgnore

			concreteProperty := ""
			if isInterface {
				concreteProperty = cname + "_Concrete"
				concreteTypeName = f.ConcreteType
				if f.Collection {
					concreteTypeName = fmt.Sprintf("List<%s>", concreteTypeName)
				}
			}

			gf := data.GenField{
				FieldName:        field,
				Name:             cname,
				Type:             typeName,
				ConcreteType:     concreteTypeName,
				ConcreteProperty: concreteProperty,
				ElementType:      elementType,
				Nullable:         nullable,
				CsIgnore:         f.Flags.CsIgnore,
				SqlIgnore:        sqlignore,
				JsonIgnore:       jsonIgnore,
				IsInterface:      isInterface,
				Collection:       f.Collection,
				Key:              f.Flags.Key,
				IsBaseType:       isbase,
			}
			pkg.Fields = append(pkg.Fields, gf)
		}
	}

	if flags.Constructor {
		for _, f := range pkg.Fields {
			if !f.IsInterface {
				pkg.ConstructorFields = append(pkg.ConstructorFields, f)
			}
		}
	}

	if flags.Keys {
		for _, f := range pkg.Fields {
			if f.Key {
				pkg.KeyFields = append(pkg.KeyFields, f)
			}
		}
	}

	if flags.Concretes {
		confields := []data.GenField{}
		for _, f := range pkg.Fields {
			if f.IsInterface {
				ngfld := data.GenField{
					FieldName:        f.FieldName,
					Name:             f.Name + "_Concrete",
					Type:             f.ConcreteType,
					ConcreteType:     f.ConcreteType,
					ConcreteProperty: "",
					Nullable:         f.Nullable,
					CsIgnore:         f.CsIgnore,
					SqlIgnore:        f.SqlIgnore,
					JsonIgnore:       false,
					IsInterface:      false,
					Collection:       f.Collection,
				}
				confields = append(confields, ngfld)
			}
		}
		pkg.Fields = append(pkg.Fields, confields...)
	}

	if flags.Id {
		pkfld := data.GenField{
			FieldName:  "",
			Name:       "ID",
			Type:       "int",
			JsonIgnore: true,
			SqlIgnore:  true,
		}

		pkg.Fields = append([]data.GenField{pkfld}, pkg.Fields...)
	}

	return pkg, nil
}

func parseFlags(flagstr string) (data.GenFlags, error) {
	flags := data.GenFlags{}
	ss := strings.Split(flagstr, ",")
	for _, s := range ss {
		flg := s[0] == '+'
		if !flg && s[0] != '-' {
			return flags, fmt.Errorf("Need + or - as first character for flag, %s ... %s", flagstr, s)
		}

		p := string(s[1:])

		switch p {
		case "id":
			flags.Id = flg
		case "fields":
			flags.Fields = flg
		case "collections":
			flags.Collections = flg
		case "constructor":
			flags.Constructor = flg
		case "concretes":
			flags.Concretes = flg
		case "keys":
			flags.Keys = flg
		default:
			return flags, fmt.Errorf("Invalid flags: %s", p)
		}
	}
	return flags, nil
}
