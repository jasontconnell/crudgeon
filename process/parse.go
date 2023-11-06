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

var fldreg *regexp.Regexp = regexp.MustCompile(`^\W*(?:private|public) (.*?) (.*?) +{.*?}( *//[0-9a-zA-Z\+\-,\."\/_ ]+)?$`)

// var genericreg *regexp.Regexp = regexp.MustCompile(`([a-zA-Z\.]*?)<(.*?)>`)
var globalflagsreg *regexp.Regexp = regexp.MustCompile(`^/{2}([\+\-a-zA-Z_,0-9\/_ ]*?)$`)

type ParsedFile struct {
	Path     string
	Fields   []data.Field
	GenFlags data.GenFlags
}

type parsedField struct {
	t            string
	name         string
	codenullable bool
	sqlnullable  bool
	collection   bool
	isInterface  bool
	flags        string
}

func ParseFile(file string, baseTypes map[string]string, genericreg string) (ParsedFile, error) {
	contents, err := os.ReadFile(file)
	parsed := ParsedFile{Path: file}

	greg := regexp.MustCompile(genericreg)

	if err != nil {
		return parsed, err
	}

	flags, fields, err := getParsed(string(contents), baseTypes, greg)
	if err != nil {
		return parsed, err
	}

	flds := []data.Field{}
	for _, p := range fields {
		if p.t == "" {
			return parsed, fmt.Errorf("Type not found for %v %v", p.t, p.name)
		}

		codenullable := p.codenullable
		baseType := isBaseType(p.t, baseTypes)
		if !p.codenullable && !baseType {
			codenullable = true
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

		sqlType := getSqlType(p.t, baseTypes)

		flds = append(flds, data.Field{Type: p.t, ConcreteType: concreteType, Name: name, FieldName: field, Nullable: codenullable, Collection: p.collection, IsInterface: p.isInterface, SqlType: sqlType, IsBaseType: baseType, Flags: fieldFlags})
	}

	parsed.Fields = flds
	parsed.GenFlags = flags

	return parsed, nil
}

func getParsed(c string, baseTypes map[string]string, greg *regexp.Regexp) (data.GenFlags, []parsedField, error) {
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

			tmatches := greg.FindAllStringSubmatch(t, -1)
			var nullable, collection, baseType bool
			baseType = isBaseType(t, baseTypes)
			if len(tmatches) > 0 {
				nullable = true
				collection = isCollection(tmatches[0][1])
				t = tmatches[0][2]
			} else {
				nullable = !baseType
			}

			flagstr := strings.TrimPrefix(m[3], " //")

			isInterface := nullable && strings.HasPrefix(t, "I") && !baseType
			nullable = nullable || collection || isInterface
			sqlnullable := nullable || t == "string"

			p := parsedField{
				t:            t,
				name:         strings.TrimSuffix(m[2], "Field"),
				codenullable: nullable,
				sqlnullable:  sqlnullable,
				collection:   collection,
				isInterface:  isInterface,
				flags:        flagstr,
			}

			plist = append(plist, p)
		}
	}

	return genflags, plist, nil
}

func isCollection(t string) bool {
	return strings.HasPrefix(t, "List")
}

func isBaseType(t string, baseTypes map[string]string) bool {
	_, ok := baseTypes[strings.ToLower(t)]
	return ok
}

func getSqlType(t string, baseTypes map[string]string) string {
	s, ok := baseTypes[t]
	if ok {
		return s
	}
	return ""
}
