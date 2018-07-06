package process

import (
	"fmt"
	"io/ioutil"
	"lpgagen/data"
	"regexp"
	"strings"
)

var fldreg *regexp.Regexp = regexp.MustCompile(`\W*(?:private|public) (.*?) (.*?) {`)
var nilreg *regexp.Regexp = regexp.MustCompile(`System.Nullable<(.*?)>`)

type parsed struct {
	t           string
	name        string
	csnullable  bool
	sqlnullable bool
}

func Parse(file string) ([]data.Field, error) {
	contents, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	parsed := getParsed(string(contents))

	flds := []data.Field{}
	for _, p := range parsed {
		fld := data.Field{Name: p.name, CsNullable: p.csnullable, SqlNullable: p.sqlnullable}
		ct, st := getTypes(p.t)
		if ct == data.CNone || st == data.SNone {
			return nil, fmt.Errorf("Type not found for %v", p.name)
		}

		fld.CsType = ct
		fld.SqlType = st

		flds = append(flds, fld)
	}

	return flds, nil
}

func getParsed(c string) []parsed {
	plist := []parsed{}

	matches := fldreg.FindAllStringSubmatch(c, -1)
	for _, m := range matches {
		t := m[1]

		tmatches := nilreg.FindAllStringSubmatch(t, -1)
		nullable := len(tmatches) > 0
		if len(tmatches) > 0 {
			t = tmatches[0][1]
		}

		sqlnullable := nullable || t == "string"

		p := parsed{t: t, name: strings.TrimSuffix(m[2], "Field"), csnullable: nullable, sqlnullable: sqlnullable}

		plist = append(plist, p)
	}

	return plist
}

func getTypes(t string) (data.CsType, data.SqlType) {
	ct := data.CNone
	st := data.SNone
	switch t {
	case "int":
		ct = data.CInt
		st = data.SInt
	case "short":
		ct = data.CShort
		st = data.SShort
	case "string":
		ct = data.CString
		st = data.SString
	case "decimal":
		ct = data.CDecimal
		st = data.SDecimal
	case "double":
		ct = data.CDouble
		st = data.SDouble
	case "long":
		ct = data.CLong
		st = data.SLong
	case "DateTime", "System.DateTime":
		ct = data.CDateTime
		st = data.SDateTime
	}

	return ct, st
}
