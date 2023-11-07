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
	dbnullable   bool
	collection   bool
	isInterface  bool
	flags        string
	codeType     string
	codeDefault  string
	dbType       string
	dbDefault    string
}

func ParseFile(file string, baseTypes map[string]data.MappedType, genericreg string) (ParsedFile, error) {
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
		baseType, isBaseType := baseTypes[p.t]
		if !p.codenullable && !isBaseType {
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

		field = strings.TrimPrefix(field, "[")
		field = strings.TrimSuffix(field, "]")

		var fieldFlags data.FieldFlags
		if p.flags != "" {
			fieldFlags, err = parseFieldFlags(p.flags)
			if err != nil {
				return parsed, err
			}
		}

		flds = append(flds,
			data.Field{
				Type:         p.t,
				ConcreteType: concreteType,
				Name:         name,
				FieldName:    field,
				Nullable:     codenullable,
				Collection:   p.collection,
				IsInterface:  p.isInterface,
				DbType:       baseType.DbType,
				IsBaseType:   isBaseType,
				CodeType:     baseType.CodeType,
				CodeDefault:  p.codeDefault,
				DbDefault:    p.dbDefault,
				Flags:        fieldFlags})
	}

	parsed.Fields = flds
	parsed.GenFlags = flags

	return parsed, nil
}

func getParsed(c string, baseTypes map[string]data.MappedType, greg *regexp.Regexp) (data.GenFlags, []parsedField, error) {
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
			var nullable, collection, isBasetype bool

			fmt.Println(baseTypes)

			if len(tmatches) > 0 {
				nullable = true
				collection = isCollection(tmatches[0][1])
				t = tmatches[0][2]
			} else {
				nullable = !isBasetype
			}

			baseType, isBaseType := baseTypes[t]

			flagstr := strings.TrimPrefix(m[3], " //")

			isInterface := !isBasetype && strings.HasPrefix(t, "I")
			nullable = nullable || collection || isInterface
			dbnullable := nullable || t == "string"

			if nullable {
				baseType.CodeDefault = "null"
			}

			if dbnullable {
				baseType.DbDefault = "null"
			}

			if !isBaseType {
				baseType.CodeType = t
			}

			fmt.Println(m[2], baseType)

			p := parsedField{
				t:            t,
				name:         strings.TrimSuffix(m[2], "Field"),
				codenullable: nullable,
				dbnullable:   dbnullable,
				collection:   collection,
				isInterface:  isInterface,
				flags:        flagstr,
				codeType:     baseType.CodeType,
				codeDefault:  baseType.CodeDefault,
				dbType:       baseType.DbType,
				dbDefault:    baseType.DbDefault,
			}

			plist = append(plist, p)
		}
	}

	return genflags, plist, nil
}

func isCollection(t string) bool {
	return strings.HasPrefix(t, "List")
}
