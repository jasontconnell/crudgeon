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

func GetGenPackage(name, path string, flds []data.Field, fileType, tmplFile, prefix, folder string, instruction string) (data.GenPackage, error) {
	inst, err := parseInstructions(instruction)
	if err != nil {
		return data.GenPackage{}, err
	}

	pkg := data.GenPackage{Name: name, Path: filepath.Join(path, folder), TemplateFile: tmplFile, Prefix: prefix, OutputFile: prefix + name + "." + fileType}
	if inst.Fields {
		for _, f := range flds {
			if f.Collection && !inst.Collections {
				continue
			}

			field := f.FieldName
			cname := strings.Title(f.Name)
			if len(cname) < 3 {
				cname = strings.ToUpper(cname)
			}

			if cname == "ID" {
				cname = name + "ID"
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
			sqlignore := sqltype == "" || f.Collection || f.Instructions.SqlIgnore

			if !isBaseType(typeName) {
				nullable = false
			}

			concreteProperty := ""
			if isInterface {
				concreteProperty = cname + "_Concrete"
				concreteTypeName = f.ConcreteType
				if f.Collection {
					concreteTypeName = fmt.Sprintf("List<%s>", concreteTypeName)
				}
			}

			ignore := (sqlignore && fileType == "sql")

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
					JsonIgnore:       f.JsonIgnore,
					IsInterface:      isInterface,
					Collection:       f.Collection,
				}
				pkg.Fields = append(pkg.Fields, gf)
			}
		}
	}

	if inst.Constructor {
		for _, f := range pkg.Fields {
			if !f.IsInterface {
				pkg.ConstructorFields = append(pkg.ConstructorFields, f)
			}
		}
	}

	if inst.Concretes {
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

	if inst.Id {
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

func parseInstructions(instructions string) (data.GenInstruct, error) {
	inst := data.GenInstruct{}
	ss := strings.Split(instructions, ",")
	for _, s := range ss {
		flg := s[0] == '+'
		if !flg && s[0] != '-' {
			return inst, fmt.Errorf("Need + or - as first character for instructions, %s ... %s", instructions, s)
		}

		p := string(s[1:])

		switch p {
		case "id":
			inst.Id = flg
		case "fields":
			inst.Fields = flg
		case "collections":
			inst.Collections = flg
		case "constructor":
			inst.Constructor = flg
		case "concretes":
			inst.Concretes = flg
		default:
			return inst, fmt.Errorf("Invalid instruction: %s", p)
		}
	}
	return inst, nil
}
