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

func Generate(pkg data.GenPackage) error {
	tmpl, err := template.New(pkg.TemplateFile).Funcs(fns).ParseFiles(pkg.TemplateFile)
	if err != nil {
		return err
	}

	err = fileutil.MakeDirs(pkg.Path)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	_, templateName := filepath.Split(pkg.TemplateFile)
	terr := tmpl.ExecuteTemplate(buffer, templateName, pkg)

	if terr != nil {
		return terr
	}

	output := filepath.Join(pkg.Path, pkg.OutputFile)
	return ioutil.WriteFile(output, buffer.Bytes(), os.ModePerm)
}

func GetGenPackage(name, path string, flds []data.Field, fileType, tmplFile, ns, prefix, folder string, flagstr string) (data.GenPackage, error) {
	flags, err := parseFlags(flagstr)
	if err != nil {
		return data.GenPackage{}, err
	}

	pkg := data.GenPackage{Name: name, Namespace: ns, Path: filepath.Join(path, folder), TemplateFile: tmplFile, Prefix: prefix, OutputFile: prefix + name + "." + fileType}
	if flags.Fields {
		for _, f := range flds {
			if f.Collection && !flags.Collections {
				continue
			}

			field := f.FieldName
			cname := strings.Title(f.Name)
			if len(cname) < 3 {
				cname = strings.ToUpper(cname)
			}

			sqltype := getSqlType(f.Type)
			if fileType == "sql" && sqltype == "" {
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

			if !isBaseType(typeName) {
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

			ignore := (sqlignore && fileType == "sql") || (f.Flags.CsIgnore && fileType == "cs")

			if !ignore {
				gf := data.GenField{
					FieldName:        field,
					Name:             cname,
					Type:             typeName,
					ConcreteType:     concreteTypeName,
					ConcreteProperty: concreteProperty,
					ElementType:      elementType,
					Nullable:         nullable,
					CsIgnore:         false,
					SqlIgnore:        sqlignore,
					JsonIgnore:       jsonIgnore,
					IsInterface:      isInterface,
					Collection:       f.Collection,
				}
				pkg.Fields = append(pkg.Fields, gf)
			}
		}
	}

	if flags.Constructor {
		for _, f := range pkg.Fields {
			if !f.IsInterface {
				pkg.ConstructorFields = append(pkg.ConstructorFields, f)
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
		default:
			return flags, fmt.Errorf("Invalid flags: %s", p)
		}
	}
	return flags, nil
}
