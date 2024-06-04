package process

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/google/uuid"

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

	m["fldbitflag"] = func(fld data.GenField, k string) bool {
		val := false
		if fld.Flags.Custom == nil {
			return val
		}
		if v, ok := fld.Flags.Custom[k]; ok {
			val = v.Flag
		}
		return val
	}

	m["fldstringflag"] = func(fld data.GenField, k string) string {
		val := ""
		if fld.Flags.Custom == nil {
			return val
		}
		if v, ok := fld.Flags.Custom[k]; ok {
			val = v.Value
		}
		return val
	}

	m["newguid"] = func() string {
		v := uuid.New()

		return strings.ToLower(v.String())
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

func getFlagValue(flags data.Flags, fileflags data.Flags, name string) bool {
	if flags.IsFlagSpecified(name) && fileflags.IsFlagSpecified(name) {
		return fileflags.GetFlagValue(name)
	} else if fileflags.IsFlagSpecified(name) {
		return fileflags.GetFlagValue(name)
	}
	return flags.GetFlagValue(name)
}

func GetGenPackage(name, path string, flds []data.Field, db bool, tmplFile, ns, prefix, suffix, folder, ext, flagstr, coll, icoll string, fileflags data.GenFlags, usefieldname bool, conditionFlag string) (data.GenPackage, error) {
	flags := data.GenFlags{}
	err := flags.MergeParse(flagstr)
	if err != nil {
		return data.GenPackage{}, err
	}

	flags.Id = getFlagValue(flags, fileflags, data.IdFlag)
	flags.Fields = getFlagValue(flags, fileflags, data.FieldsFlag)
	flags.Collections = getFlagValue(flags, fileflags, data.CollectionsFlag)
	flags.Concretes = getFlagValue(flags, fileflags, data.ConcretesFlag)
	flags.Constructor = getFlagValue(flags, fileflags, data.ConstructorFlag)
	flags.Keys = getFlagValue(flags, fileflags, data.KeysFlag)
	flags.DbIgnore = getFlagValue(flags, fileflags, data.DbIgnoreFlag)
	flags.CodeIgnore = getFlagValue(flags, fileflags, data.CodeIgnoreFlag)
	flags.XmlIgnore = getFlagValue(flags, fileflags, data.XmlIgnoreFlag)
	flags.JsonIgnore = getFlagValue(flags, fileflags, data.JsonIgnoreFlag)
	flags.HashIgnore = getFlagValue(flags, fileflags, data.HashIgnoreFlag)
	flags.XmlRoot = getFlagValue(flags, fileflags, data.XmlRootFlag)
	flags.Merge = getFlagValue(flags, fileflags, data.MergeFlag)

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

	flags.Table = fileflags.Table
	flags.TableName = fileflags.TableName
	if !flags.Table && !flags.Class {
		flags.TableName = name
	} else if !flags.Table && flags.Class {
		flags.TableName = flags.ClassName
	}

	if name == "" && flags.ClassName == "" {
		return data.GenPackage{Generate: false}, fmt.Errorf("No object name provided.")
	}

	if name == "" {
		name = flags.ClassName
	}

	if fileflags.XmlRootName != "" {
		flags.XmlRootName = fileflags.XmlRootName
	}

	if db && flags.DbIgnore || (!db && flags.CodeIgnore) {
		return data.GenPackage{Generate: false}, nil
	}

	if conditionFlag != "" && !getFlagValue(flags, fileflags, conditionFlag) {
		return data.GenPackage{Generate: false}, nil
	}

	pkg := data.GenPackage{Generate: true, Name: name, Namespace: ns, Path: filepath.Join(path, folder), TemplateFile: tmplFile, Prefix: prefix, Suffix: suffix, OutputFile: prefix + name + suffix + "." + ext, Flags: flags}
	if flags.Fields || flags.Constructor || flags.Keys || flags.Concretes || flags.PrimaryKeys || flags.Updates {
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

			dbignore := f.Flags.DbIgnore
			dbtype := f.DbType

			if f.Flags.ForceDb {
				dbignore = false
				dbtype = f.Flags.ForceDbType
			}

			if db && (dbtype == "" || dbignore) {
				continue
			}

			codeignore := f.Flags.CodeIgnore
			if !db && codeignore {
				continue
			}

			xmlignore := f.Flags.XmlIgnore || flags.XmlIgnore
			hashIgnore := f.Flags.HashIgnore || flags.HashIgnore

			isInterface := f.Type != f.ConcreteType
			typeName, concreteTypeName, elementType := f.Type, f.Type, f.Type
			if f.Collection {
				listType := coll
				if isInterface {
					listType = icoll
				}
				typeName = fmt.Sprintf("%s<%s>", listType, typeName)
			}

			if db {
				typeName = dbtype
			}

			nullable := f.Nullable

			isBase := false
			if !db {
				isBase = f.IsBaseType
				if !isBase {
					nullable = false
				}
			}

			dbignore = (f.Collection || dbignore || !isBase) && !f.Flags.ForceDb
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

			xmlwrapper := f.Flags.XmlWrapper && !db

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
				CodeIgnore:          f.Flags.CodeIgnore,
				DbIgnore:            dbignore,
				JsonIgnore:          jsonIgnore,
				XmlIgnore:           xmlignore,
				HashIgnore:          hashIgnore,
				IsInterface:         isInterface,
				Collection:          f.Collection,
				Key:                 f.Flags.Key,
				ForeignKey:          f.Flags.ForeignKey,
				IsBaseType:          isBase,
				Flags:               f.Flags,
				CodeType:            f.CodeType,
				CodeDefault:         f.CodeDefault,
				DbType:              f.DbType,
				DbDefault:           f.DbDefault,
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
			if f.Key || f.ForeignKey {
				pkg.KeyFields = append(pkg.KeyFields, f)
			}
		}
	}

	if flags.PrimaryKeys {
		for _, f := range pkg.Fields {
			if f.Key {
				pkg.PrimaryKeyFields = append(pkg.PrimaryKeyFields, f)
			}
		}
	}

	if flags.Updates {
		for _, f := range pkg.Fields {
			if !f.Flags.Auto && !f.Key {
				pkg.UpdateFields = append(pkg.UpdateFields, f)
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
					CodeIgnore:          f.CodeIgnore,
					DbIgnore:            f.DbIgnore,
					XmlIgnore:           f.XmlIgnore,
					JsonIgnore:          flags.JsonIgnore,
					HashIgnore:          f.HashIgnore,
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
			DbIgnore:   true,
			Id:         true,
		}

		pkg.Fields = append([]data.GenField{pkfld}, pkg.Fields...)
	}

	return pkg, nil
}
