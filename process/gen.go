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
	if !pkg.Generate {
		return nil
	}
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

func GetGenPackage(name, path string, flds []data.Field, fileType, tmplFile, ns, prefix, folder string, flagstr string, fileflags data.GenFlags, usefieldname bool) (data.GenPackage, error) {
	flags, err := parseGenFlags(flagstr)

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

			if usefieldname {
				cname = field
			}

			sqlignore := f.Flags.SqlIgnore
			sqltype := getSqlType(f.Type)
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

			sqlignore = f.Collection || sqlignore || !isBase
			jsonIgnore := f.JsonIgnore || f.Flags.JsonIgnore || flags.JsonIgnore

			concreteProperty := ""
			concreteElementType := ""
			if isInterface {
				concreteProperty = cname + "_Concrete"
				concreteTypeName = f.ConcreteType
				concreteElementType = f.ConcreteType
				if f.Collection {
					concreteTypeName = fmt.Sprintf("List<%s>", concreteTypeName)
				}
			}

			xmlwrapper := f.Flags.XmlWrapper && fileType == "cs"

			gf := data.GenField{
				Access: "public",
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

				field := f.FieldName
				typeName := f.ConcreteType

				xmlwrappertype, xmlwrappername, xmlwrapperelement := f.ConcreteType, "", ""
				if !f.XmlIgnore && f.XmlWrapper && f.Collection {
					xmlwrappertype = typeName
					xmlwrappername = field
					xmlwrapperelement = f.XmlWrapperElement
					typeName = field + "Wrapper"
				}

				ngfld := data.GenField{
					Access: "public",
					FieldName:           field,
					Name:                f.Name + "_Concrete",
					Type:                typeName,
					ConcreteType:        typeName,
					ConcreteElementType: f.ConcreteElementType,
					ConcreteProperty:    "",
					XmlWrapper:          f.XmlWrapper,
					XmlWrapperType:      xmlwrappertype,
					XmlWrapperName:      xmlwrappername,
					XmlWrapperElement:   xmlwrapperelement,
					Nullable:            f.Nullable,
					CsIgnore:            f.CsIgnore,
					SqlIgnore:           f.SqlIgnore,
					XmlIgnore:           f.XmlIgnore,
					JsonIgnore:          f.JsonIgnore,
					IsInterface:         false,
					Collection:          f.Collection,
				}
				confields = append(confields, ngfld)
			}
		}
		pkg.Fields = append(pkg.Fields, confields...)
	}

	if flags.Id {
		pkfld := data.GenField{
			Access: "public",
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
