package process

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/jasontconnell/crudgeon/data"
)

var fldreg *regexp.Regexp = regexp.MustCompile(`^\W*(?:private|public) (.*?) (.*?) +{.*?}( *//[a-zA-Z\+\-,\." ]+)?$`)
var genericreg *regexp.Regexp = regexp.MustCompile(`([a-zA-Z\.]*?)<(.*?)>`)
var globalflagsreg *regexp.Regexp = regexp.MustCompile(`^//([\+\-a-zA-Z_, ]*?)$`)

type ParsedFile struct {
	Fields   []data.Field
	GenFlags data.GenFlags
}

type parsedField struct {
	t           string
	name        string
	csnullable  bool
	sqlnullable bool
	collection  bool
	isInterface bool
	flags       string
}

func ParseFile(file string) (ParsedFile, error) {
	contents, err := os.ReadFile(file)

	parsed := ParsedFile{}

	if err != nil {
		return parsed, err
	}

	flags, fields, err := getParsed(string(contents))
	if err != nil {
		return parsed, err
	}

	flds := []data.Field{}
	for _, p := range fields {
		if p.t == "" {
			return parsed, fmt.Errorf("Type not found for %v %v", p.t, p.name)
		}

		csnull := p.csnullable
		if !p.csnullable && !isBaseType(p.t) {
			csnull = true
		}

		concreteType := p.t
		if p.isInterface {
			concreteType = string(concreteType[1:])
		}

		name, field := p.name, p.name
		names := strings.Split(p.name, "|")
		if len(names) == 2 {
			name = names[1]
			field = names[0]
		}

		var fieldFlags data.FieldFlags
		if p.flags != "" {
			fieldFlags, err = parseFieldFlags(p.flags)
			if err != nil {
				return parsed, err
			}
		}

		flds = append(flds, data.Field{Type: p.t, ConcreteType: concreteType, Name: name, FieldName: field, Nullable: csnull, Collection: p.collection, IsInterface: p.isInterface, Flags: fieldFlags})
	}

	pfile := ParsedFile{Fields: flds, GenFlags: flags}

	return pfile, nil
}

func getParsed(c string) (data.GenFlags, []parsedField, error) {
	plist := []parsedField{}
	genflags := data.GenFlags{}

	s := bufio.NewScanner(bytes.NewBufferString(c))

	for s.Scan() {
		line := s.Text()
		globflags := globalflagsreg.FindAllStringSubmatch(line, -1)
		for _, m := range globflags {
			var err error
			err = genflags.MergeParse(m[1])
			if err != nil {
				return genflags, nil, err
			}
		}

		matches := fldreg.FindAllStringSubmatch(line, -1)
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

			flagstr := strings.TrimPrefix(m[3], " //")

			isInterface := nullable && strings.HasPrefix(t, "I")
			nullable = nullable || collection || isInterface
			sqlnullable := nullable || t == "string"

			p := parsedField{
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
	}

	return genflags, plist, nil
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
	case "decimal", "double":
		st = "decimal(18,7)"
	// case "double":
	// 	st = "decimal(13,7)"
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
