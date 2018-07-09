package process

import (
	"bytes"
	"github.com/jasontconnell/fileutil"
	"io/ioutil"
	"lpgagen/data"
	"os"
	"path/filepath"
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

	err = fileutil.MakeDirs(filepath.Join(pkg.Path, pkg.Name))
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	_, templateName := filepath.Split(pkg.TemplateFile)
	terr := tmpl.ExecuteTemplate(buffer, templateName, pkg)

	if terr != nil {
		return terr
	}

	output := filepath.Join(pkg.Path, pkg.Name, pkg.OutputFile)
	return ioutil.WriteFile(output, buffer.Bytes(), os.ModePerm)
}

func GetGenPackage(name, path string, flds []data.Field, fileType, tmplFile, prefix string) data.GenPackage {
	pkg := data.GenPackage{Name: name, Path: path, TemplateFile: tmplFile, Prefix: prefix, OutputFile: prefix + name + "." + fileType}
	for _, f := range flds {
		nullable := f.CsNullable
		if fileType == "sql" {
			nullable = f.SqlNullable
		}
		gf := data.GenField{Name: f.Name, Type: getTypeName(f.CsType, f.SqlType, fileType), Nullable: nullable}

		pkg.Fields = append(pkg.Fields, gf)
	}

	return pkg
}

func getTypeName(cstype data.CsType, sqltype data.SqlType, fileType string) string {
	switch fileType {
	case "sql":
		return getSqlType(sqltype)
	case "cs":
		return getCsType(cstype)
	}
	return ""
}

func getSqlType(sqltype data.SqlType) string {
	var t string
	switch sqltype {
	case data.SInt:
		t = "int"
	case data.SString:
		t = "varchar(100)"
	case data.SDecimal:
		t = "decimal(18,2)"
	case data.SDouble:
		t = "decimal(18,2)"
	case data.SShort:
		t = "smallint"
	case data.SLong:
		t = "bigint"
	case data.SDateTime:
		t = "datetime"
	case data.SBit:
		t = "bit"
	}
	return t
}

func getCsType(cstype data.CsType) string {
	var t string
	switch cstype {
	case data.CInt:
		t = "int"
	case data.CString:
		t = "string"
	case data.CDecimal:
		t = "decimal"
	case data.CDouble:
		t = "double"
	case data.CShort:
		t = "short"
	case data.CLong:
		t = "long"
	case data.CDateTime:
		t = "DateTime"
	case data.CBool:
		t = "bool"
	}
	return t
}
