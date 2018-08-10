package process

import (
	"fmt"
	"io/ioutil"
	"lpgagen/data"
	"regexp"
	"strings"
)

var fldreg *regexp.Regexp = regexp.MustCompile(`\W*(?:private|public) (.*?) (.*?) *{`)
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
		ct, st := getTypes(p.t)
		if ct == data.CNone || st == data.SNone {
			return nil, fmt.Errorf("Type not found for %v", p.name)
		}

		csnull := p.csnullable
		if ct == data.CCustom {
			csnull = true
		}

		name, rawname := p.name, p.name
		names := strings.Split(p.name, "|")
		if len(names) == 2 {
			name = names[1]
			rawname = names[0]
		}

		fld := data.Field{Type: p.t, Name: name, RawName: rawname, CsNullable: csnull, SqlNullable: p.sqlnullable}

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
	case "bool":
		ct = data.CBool
		st = data.SBit
	default: // don't ignore in C# but don't allow saving in sql
		ct = data.CCustom
		st = data.SIgnore
	}

	return ct, st
}
