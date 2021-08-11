package process

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jasontconnell/crudgeon/data"
)

var fns = template.FuncMap{
	"plus1": func(x int) int {
		return x + 1
	},
}

func getPackageFunctions(pkg data.GenPackage) template.FuncMap {
	m := make(template.FuncMap)

	m["plus1"] = func(x int) int {
		return x + 1
	}

	m["stringflag"] = func(k string) string {
		val := ""
		if pkg.Flags.Custom == nil {
			return val
		}
		if v, ok := pkg.Flags.Custom[k]; ok {
			val = v.Value
		}
		return val
	}

	m["bitflag"] = func(k string) bool {
		val := false
		if pkg.Flags.Custom == nil {
			return val
		}

		if v, ok := pkg.Flags.Custom[k]; ok {
			val = v.Flag
		}

		return val
	}

	return m
}

func Generate(pkg data.GenPackage, objdir bool) error {
	if !pkg.Generate {
		return nil
	}
	tmpl, err := template.New(pkg.TemplateFile).Funcs(getPackageFunctions(pkg)).ParseFiles(pkg.TemplateFile)
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

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	output := filepath.Join(path, pkg.OutputFile)
	return os.WriteFile(output, buffer.Bytes(), os.ModePerm)
}

func GetGenPackage(name, path string, flds []data.Field, fileType, tmplFile, ns, prefix, folder, flagstr string, fileflags data.GenFlags, usefieldname bool) (data.GenPackage, error) {
	flags := data.GenFlags{}
	err := flags.MergeParse(flagstr)
	if err != nil {
		return data.GenPackage{}, err
	}

	flags.Id = flags.Id || fileflags.Id
	flags.Fields = flags.Fields || fileflags.Fields
	flags.Collections = flags.Collections || fileflags.Collections
	flags.Concretes = flags.Concretes || fileflags.Concretes
	flags.Constructor = flags.Constructor || fileflags.Constructor
	flags.Keys = flags.Keys || fileflags.Keys
	flags.SqlIgnore = flags.SqlIgnore || fileflags.SqlIgnore
	flags.CsIgnore = flags.CsIgnore || fileflags.CsIgnore
	flags.XmlIgnore = flags.XmlIgnore || fileflags.XmlIgnore
	flags.JsonIgnore = flags.JsonIgnore || fileflags.JsonIgnore
	flags.XmlRoot = flags.XmlRoot || fileflags.XmlRoot

	if flags.Custom == nil {
		flags.Custom = make(map[string]data.CustomFlag)
	}

	if flags.Custom != nil && fileflags.Custom != nil {
		for k, v := range fileflags.Custom {
			flags.Custom[k] = v
		}
	}

	// file specific flags
	flags.Class = fileflags.Class
	flags.ClassName = fileflags.ClassName
	flags.ExactName = fileflags.ExactName

	if name == "" && flags.ClassName == "" {
		return data.GenPackage{Generate: false}, fmt.Errorf("No object name provided.")
	}

	if name == "" {
		name = flags.ClassName
	}

	if fileflags.XmlRootName != "" {
		flags.XmlRootName = fileflags.XmlRootName
	}

	if (fileType == "sql") && flags.SqlIgnore || (fileType == "cs" && flags.CsIgnore) {
		return data.GenPackage{Generate: false}, nil
	}

	if err != nil {
		return data.GenPackage{}, err
	}

	pkg := data.GenPackage{Generate: true, Name: name, Namespace: ns, Path: filepath.Join(path, folder), TemplateFile: tmplFile, Prefix: prefix, OutputFile: prefix + name + "." + fileType, Flags: flags}
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

			if usefieldname || flags.ExactName {
				cname = f.Name
			}

			sqlignore := f.Flags.SqlIgnore
			sqltype := getSqlType(f.Type)

			if f.Flags.ForceSql {
				sqlignore = false
				sqltype = f.Flags.ForceSqlType
			}

			if fileType == "sql" && (sqltype == "" || sqlignore) {
				continue
			}

			csignore := f.Flags.CsIgnore
			if fileType == "cs" && csignore {
				continue
			}

			xmlignore := f.Flags.XmlIgnore || flags.XmlIgnore

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

			isBase := false
			if fileType == "cs" {
				isBase = isBaseType(typeName)
				if !isBase {
					nullable = false
				}
			}

			sqlignore = (f.Collection || sqlignore || !isBase) && !f.Flags.ForceSql
			jsonIgnore := f.Flags.JsonIgnore || flags.JsonIgnore

			concreteProperty := ""
			concreteElementType := ""
			if isInterface {
				concreteProperty = cname + "_Concrete"
				concreteTypeName = f.ConcreteType
				concreteElementType = f.ConcreteType
				if f.Collection {
					concreteTypeName = fmt.Sprintf("List<%s>", concreteTypeName)
				}
				jsonIgnore = true
			}

			xmlwrapper := f.Flags.XmlWrapper && fileType == "cs"

			gf := data.GenField{
				Access:              "public",
				FieldName:           field,
				Name:                cname,
				Type:                typeName,
				ConcreteType:        concreteTypeName,
				ConcreteProperty:    concreteProperty,
				ConcreteElementType: concreteElementType,
				ElementType:         elementType,
				XmlWrapper:          xmlwrapper,
				XmlWrapperElement:   f.Flags.XmlWrapperElement,
				Nullable:            nullable,
				CsIgnore:            f.Flags.CsIgnore,
				SqlIgnore:           sqlignore,
				JsonIgnore:          jsonIgnore,
				XmlIgnore:           xmlignore,
				IsInterface:         isInterface,
				Collection:          f.Collection,
				Key:                 f.Flags.Key,
				IsBaseType:          isBase,
				Flags:               f.Flags,
			}
			pkg.Fields = append(pkg.Fields, gf)
		}
	}

	if flags.Constructor {
		for _, f := range pkg.Fields {
			if !f.IsInterface && !f.Flags.ReadOnly {
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

				field := f.FieldName
				typeName := f.ConcreteType

				xmlwrappertype, xmlwrappername := f.ConcreteType, ""
				// if !f.XmlIgnore && f.XmlWrapper && f.Collection {
				// 	xmlwrappertype = typeName
				// 	xmlwrappername = field
				// 	xmlwrapperelement = f.XmlWrapperElement
				// 	typeName = pkg.Name + field + "Wrapper"
				// }

				ngfld := data.GenField{
					Access:              "public",
					FieldName:           field,
					Name:                f.Name + "_Concrete",
					Type:                typeName,
					ConcreteType:        typeName,
					ConcreteElementType: f.ConcreteElementType,
					ConcreteProperty:    "",
					XmlWrapper:          f.XmlWrapper,
					XmlWrapperType:      xmlwrappertype,
					XmlWrapperName:      xmlwrappername,
					XmlWrapperElement:   f.XmlWrapperElement,
					Nullable:            f.Nullable,
					CsIgnore:            f.CsIgnore,
					SqlIgnore:           f.SqlIgnore,
					XmlIgnore:           f.XmlIgnore,
					JsonIgnore:          flags.JsonIgnore,
					IsInterface:         false,
					Collection:          f.Collection,
					Flags:               f.Flags,
				}
				confields = append(confields, ngfld)
			}
		}
		pkg.Fields = append(pkg.Fields, confields...)
	}

	if flags.Id {
		pkfld := data.GenField{
			Access:     "public",
			FieldName:  "",
			Name:       "ID",
			Type:       "int",
			JsonIgnore: true,
			SqlIgnore:  true,
			Id:         true,
		}

		pkg.Fields = append([]data.GenField{pkfld}, pkg.Fields...)
	}

	return pkg, nil
}
