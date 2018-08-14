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

func GetGenPackage(name, path string, flds []data.Field, fileType, tmplFile, prefix, folder string, forInterface bool) data.GenPackage {
	pkg := data.GenPackage{Name: name, Path: filepath.Join(path, folder), TemplateFile: tmplFile, Prefix: prefix, OutputFile: prefix + name + "." + fileType}
	for _, f := range flds {
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
		sqlignore := sqltype == "" || f.Collection

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

	if !forInterface {
		confields := []data.GenField{}
		for _, f := range pkg.Fields {
			if !f.IsInterface {
				pkg.ConstructorFields = append(pkg.ConstructorFields, f)
			} else if f.IsInterface {
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

	pkfld := data.GenField{
		FieldName:  "",
		Name:       "ID",
		Type:       "int",
		JsonIgnore: true,
	}

	pkg.Fields = append([]data.GenField{pkfld}, pkg.Fields...)

	return pkg
}
