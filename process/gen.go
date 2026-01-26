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
		path = filepath.Join(path, pkg.Object.Name)
	}

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	outfile := pkg.Object.Name + "." + pkg.Ext
	if pkg.FilenameTemplate != "" {
		p := struct{ Name, NameLower string }{Name: pkg.Object.Name, NameLower: pkg.Object.NameLower}
		outfile = processTemplate(pkg.FilenameTemplate, p) + "." + pkg.Ext
	}

	output := filepath.Join(path, outfile)
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

func GetGenPackage(name, path string, pfile ParsedFile, db bool, tmplFile, ns, outputTmpl, folder, ext, flagstr, colltmpl string, conditionFlag string) (data.GenPackage, error) {
	flags := data.GenFlags{}
	err := flags.MergeParse(flagstr)
	if err != nil {
		return data.GenPackage{}, fmt.Errorf("parsing flags %s. %w", flagstr, err)
	}

	gf := data.MergeGenFlags(flags, pfile.GenFlags)
	gf.Database = db
	gf.CollectionTemplate = colltmpl

	if db && gf.DbIgnore || (!db && gf.CodeIgnore) {
		return data.GenPackage{Generate: false}, nil
	}

	if conditionFlag != "" && !getFlagValue(flags, gf, conditionFlag) {
		return data.GenPackage{Generate: false}, nil
	}

	obj, err := getGenObject(pfile, gf)
	if err != nil {
		return data.GenPackage{Generate: false}, fmt.Errorf("get gen object %s. %w", name, err)
	}
	pkg := data.GenPackage{Generate: true, Object: obj, Path: filepath.Join(path, folder), TemplateFile: tmplFile, Ext: ext, FilenameTemplate: outputTmpl, Flags: gf, Namespace: ns, Imports: getImports(obj.Fields)}
	return pkg, nil
}

func GetAllGenPackage(path string, pfiles []ParsedFile, db bool, tmplFile, ns, outputTmpl, folder, ext, flagstr, colltmpl, conditionFlag string) (data.GenPackage, error) {
	flags := data.GenFlags{}
	err := flags.MergeParse(flagstr)
	if err != nil {
		return data.GenPackage{}, fmt.Errorf("parsing flags %s. %w", flagstr, err)
	}

	flags.Database = db
	flags.CollectionTemplate = colltmpl

	imports := []string{}
	impmap := make(map[string]bool)

	objs := []data.GenObject{}
	for _, pfile := range pfiles {
		if conditionFlag != "" && !getFlagValue(flags, pfile.GenFlags, conditionFlag) {
			continue
		}

		if db && pfile.GenFlags.DbIgnore || (!db && pfile.GenFlags.CodeIgnore) {
			continue
		}

		gf := data.MergeGenFlags(flags, pfile.GenFlags)

		obj, err := getGenObject(pfile, gf)
		if err != nil {
			return data.GenPackage{}, fmt.Errorf("get gen object %s. %w", pfile.GenFlags.ClassName, err)
		}

		localImports := getImports(obj.Fields)
		for _, imp := range localImports {
			if _, ok := impmap[imp]; !ok {
				imports = append(imports, imp)
				impmap[imp] = true
			}
		}

		objs = append(objs, obj)
	}

	pkg := data.GenPackage{Generate: true, Objects: objs, Path: filepath.Join(path, folder), TemplateFile: tmplFile, Ext: ext, FilenameTemplate: outputTmpl, Flags: flags, Namespace: ns, Imports: imports}
	return pkg, nil
}

func getGenObject(pfile ParsedFile, genflags data.GenFlags) (data.GenObject, error) {

	if pfile.GenFlags.ClassName == "" {
		return data.GenObject{}, fmt.Errorf("No object name provided.")
	}

	name := pfile.GenFlags.ClassName
	if name == "" {
		name = genflags.ClassName
	}

	lname := strings.ToLower(name)

	skipFlags := make(map[string]bool)
	for _, flg := range genflags.Custom {
		if !flg.Flag {
			skipFlags[flg.Name] = true
		}
	}

	obj := data.GenObject{Name: name, NameLower: lname, Namespace: pfile.GenFlags.Namespace}

	if genflags.Fields || genflags.Constructor || genflags.Keys || genflags.PrimaryKeys || genflags.Updates {
		for _, f := range pfile.Fields {
			if f.Collection && !genflags.Collections {
				continue
			}

			if len(skipFlags) > 0 {
				skip := false
				for _, flg := range f.Flags.Custom {
					if _, ok := skipFlags[flg.Name]; ok && flg.Flag {
						skip = true
					}
				}
				if skip {
					continue
				}
			}

			field := f.FieldName
			cname := strings.Title(f.Name)
			lname := strings.ToLower(f.Name)
			if len(cname) < 3 {
				cname = strings.ToUpper(cname)
			}

			if genflags.ExactName {
				cname = f.Name
			}

			dbignore := f.Flags.DbIgnore
			dbtype := f.DbType

			if f.Flags.ForceDb {
				dbignore = false
				dbtype = f.Flags.ForceDbType
			}

			if genflags.Database && (dbtype == "" || dbignore) {
				continue
			}

			codeignore := f.Flags.CodeIgnore
			if !genflags.Database && codeignore {
				continue
			}

			xmlignore := f.Flags.XmlIgnore || genflags.XmlIgnore
			hashIgnore := f.Flags.HashIgnore || genflags.HashIgnore

			typeName, elementType := f.Type, f.CollectionType
			if f.Collection {
				typeName = processTemplate(genflags.CollectionTemplate, struct{ Name string }{Name: elementType})
			}

			if genflags.Database {
				typeName = dbtype
			}

			nullable := f.Nullable

			isBase := false
			if !genflags.Database {
				isBase = f.IsBaseType
				if !isBase {
					nullable = false
				} else if !f.Collection {
					typeName = f.CodeType
				}
			}

			dbignore = (f.Collection || dbignore || !isBase) && !f.Flags.ForceDb
			jsonIgnore := f.Flags.JsonIgnore || genflags.JsonIgnore

			xmlwrapper := f.Flags.XmlWrapper && !genflags.Database

			gf := data.GenField{
				FieldName:         field,
				Name:              cname,
				NameLower:         lname,
				Type:              typeName,
				ElementType:       elementType,
				XmlWrapper:        xmlwrapper,
				XmlWrapperElement: f.Flags.XmlWrapperElement,
				Nullable:          nullable,
				CodeIgnore:        f.Flags.CodeIgnore,
				DbIgnore:          dbignore,
				JsonIgnore:        jsonIgnore,
				XmlIgnore:         xmlignore,
				HashIgnore:        hashIgnore,
				Collection:        f.Collection,
				Key:               f.Flags.Key,
				ForeignKey:        f.Flags.ForeignKey,
				IsBaseType:        isBase,
				Flags:             f.Flags,
				CodeType:          f.CodeType,
				CodeDefault:       f.CodeDefault,
				DbType:            f.DbType,
				DbDefault:         f.DbDefault,
				Include:           f.Include,
			}
			obj.Fields = append(obj.Fields, gf)
		}
	}

	if genflags.Constructor {
		for _, f := range obj.Fields {
			if !f.Collection && !f.Flags.ReadOnly {
				obj.ConstructorFields = append(obj.ConstructorFields, f)
			}
		}
	}

	if genflags.Keys {
		for _, f := range obj.Fields {
			if f.Key || f.ForeignKey {
				obj.KeyFields = append(obj.KeyFields, f)
			}
		}
	}

	if genflags.PrimaryKeys {
		for _, f := range obj.Fields {
			if f.Key {
				obj.PrimaryKeyFields = append(obj.PrimaryKeyFields, f)
			}
		}
	}

	if genflags.Updates {
		for _, f := range obj.Fields {
			if !f.Flags.Auto && !f.Key {
				obj.UpdateFields = append(obj.UpdateFields, f)
			}
		}
	}

	if genflags.Id {
		pkfld := data.GenField{
			FieldName:  "",
			Name:       "ID",
			Type:       "int",
			JsonIgnore: true,
			DbIgnore:   true,
			Id:         true,
		}

		obj.Fields = append([]data.GenField{pkfld}, obj.Fields...)
	}

	return obj, nil
}

func getImports(flds []data.GenField) []string {
	mm := make(map[string]bool)
	list := []string{}
	for _, f := range flds {
		if len(f.Include) == 0 {
			continue
		}
		if _, ok := mm[f.Include]; !ok {
			mm[f.Include] = true
			list = append(list, f.Include)
		}
	}
	return list
}
