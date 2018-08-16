package process

import (
	"fmt"
	"io/ioutil"
	"lpgagen/data"
	"regexp"
	"strings"
)

var fldreg *regexp.Regexp = regexp.MustCompile(`\W*(?:private|public) (.*?) (.*?) +{.*?}( *//[a-z\+\-]+)?`)
var genericreg *regexp.Regexp = regexp.MustCompile(`([a-zA-Z\.]*?)<(.*?)>`)
var flagsreg *regexp.Regexp = regexp.MustCompile(`//(.*?)$`)

type parsed struct {
	t           string
	name        string
	csnullable  bool
	sqlnullable bool
	collection  bool
	isInterface bool
	flags       string
}

func ParseFields(file string) ([]data.Field, error) {
	contents, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	parsed := getParsed(string(contents))

	flds := []data.Field{}
	for _, p := range parsed {
		if p.t == "" {
			return nil, fmt.Errorf("Type not found for %v %v", p.t, p.name)
		}

		csnull := p.csnullable
		if !p.csnullable && !isBaseType(p.t) {
			csnull = true
		}

		jsonIgnore := false

		concreteType := p.t
		if p.isInterface {
			concreteType = string(concreteType[1:])
			jsonIgnore = true
		}

		name, field := p.name, p.name
		names := strings.Split(p.name, "|")
		if len(names) == 2 {
			name = names[1]
			field = names[0]
		}

		var flags data.FieldFlags
		if p.flags != "" {
			flags, err = parseFieldFlags(p.flags)
			if err != nil {
				return nil, err
			}
		}

		flds = append(flds, data.Field{Type: p.t, ConcreteType: concreteType, Name: name, FieldName: field, Nullable: csnull, Collection: p.collection, JsonIgnore: jsonIgnore, IsInterface: p.isInterface, Flags: flags})
	}

	return flds, nil
}

func getParsed(c string) []parsed {
	plist := []parsed{}

	matches := fldreg.FindAllStringSubmatch(c, -1)
	for _, m := range matches {
		t := m[1]

		tmatches := genericreg.FindAllStringSubmatch(t, -1)
		var nullable, collection bool
		if len(tmatches) > 0 {
			nullable = true
			collection = isCollection(tmatches[0][1])
			t = tmatches[0][2]
		} else {
			nullable = !isBaseType(t)
		}

		flagstr := ""
		imatches := flagsreg.FindAllStringSubmatch(m[0], -1)
		if len(imatches) > 0 {
			flagstr = imatches[0][1]
		}

		isInterface := nullable && strings.HasPrefix(t, "I")
		nullable = nullable || collection || isInterface
		sqlnullable := nullable || t == "string"

		p := parsed{
			t:           t,
			name:        strings.TrimSuffix(m[2], "Field"),
			csnullable:  nullable,
			sqlnullable: sqlnullable,
			collection:  collection,
			isInterface: isInterface,
			flags:       flagstr,
		}

		plist = append(plist, p)
	}

	return plist
}

func isCollection(t string) bool {
	return strings.HasPrefix(t, "List") || strings.HasPrefix(t, "IEnumerable")
}

func isBaseType(t string) bool {
	switch t {
	case "int", "short", "string", "decimal", "double", "long", "DateTime", "bool":
		return true
	}
	return false
}

func getSqlType(t string) string {
	st := ""
	switch t {
	case "int":
		st = "int"
	case "short":
		st = "smallint"
	case "string":
		st = "varchar(150)"
	case "decimal":
		st = "decimal(18,2)"
	case "double":
		st = "double"
	case "long":
		st = "bigint"
	case "DateTime":
		st = "datetime"
	case "bool":
		st = "bit"
	default: // don't ignore in C# but don't allow saving in sql
		st = ""
	}

	return st
}

func parseFieldFlags(instructions string) (data.FieldFlags, error) {
	flags := data.FieldFlags{}
	ss := strings.Split(instructions, ",")
	for _, s := range ss {
		flg := s[0] == '+'
		if !flg && s[0] != '-' {
			return flags, fmt.Errorf("Need + or - as first character for flags, %s: %s", instructions, s)
		}

		p := string(s[1:])

		switch p {
		case "sqlignore":
			flags.SqlIgnore = flg
		case "jsonignore":
			flags.JsonIgnore = flg
		case "csignore":
			flags.CsIgnore = flg
		case "hashkey":
			flags.HashKey = flg
		default:
			return flags, fmt.Errorf("Invalid flag: %s", p)
		}
	}
	return flags, nil
}
