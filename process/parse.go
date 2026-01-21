package process

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/jasontconnell/crudgeon/data"
)

var fldreg *regexp.Regexp = regexp.MustCompile(`^([a-zA-Z0-9\[\]_]+) (.*?)( *//[0-9a-zA-Z\+\-,\."\/_\(\) ]+)?$`)
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

func ParseFile(file string, baseTypes map[string]data.MappedType, nullableFormat, null, dbnull, nullablereg string) (ParsedFile, error) {
	contents, err := os.ReadFile(file)
	parsed := ParsedFile{Path: file}

	var nreg *regexp.Regexp

	if nullablereg != "" {
		nreg = regexp.MustCompile(nullablereg)
	}

	if err != nil {
		return parsed, err
	}

	flags, fields, err := getParsed(string(contents), baseTypes, nullableFormat, null, dbnull, nreg)
	if err != nil {
		return parsed, err
	}

	flds := []data.Field{}
	for _, p := range fields {
		if p.t == "" {
			return parsed, fmt.Errorf("Type not found for %v %v", p.t, p.name)
		}

		codenullable := p.codenullable
		_, isBaseType := baseTypes[p.t]
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

		field = strings.Trim(field, "[]")

		var fieldFlags data.FieldFlags
		if p.flags != "" {
			fieldFlags, err = data.ParseFieldFlags(p.flags)
			if err != nil {
				return parsed, fmt.Errorf("error parsing %s. %w", p.flags, err)
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
				DbType:       p.dbType,
				IsBaseType:   isBaseType,
				CodeType:     p.codeType,
				CodeDefault:  p.codeDefault,
				DbDefault:    p.dbDefault,
				Flags:        fieldFlags})
	}

	parsed.Fields = flds
	parsed.GenFlags = flags

	return parsed, nil
}

func getParsed(c string, baseTypes map[string]data.MappedType, nullableFormat, null, dbnull string, nreg *regexp.Regexp) (data.GenFlags, []parsedField, error) {
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

			var nmatches [][]string
			if nreg != nil {
				nmatches = nreg.FindAllStringSubmatch(t, -1)
			}
			var nullable, collection bool

			if len(nmatches) > 0 {
				nullable = true
				t = nmatches[0][1]
			}

			log.Println(t)
			if strings.HasSuffix(t, "[]") {
				collection = true
				t = strings.TrimRight(t, "[]")
			}

			// if len(tmatches) > 0 && len(tmatches[0]) > 0 && !nullable {
			// 	log.Println("collection", tmatches, tmatches[0])
			// 	collection = isCollection(tmatches[0][1], conccoll, abstcoll)
			// 	t = tmatches[0][2]
			// }

			baseType, isBaseType := baseTypes[t]
			flagstr := strings.TrimPrefix(m[3], " //")

			isInterface := !isBaseType && strings.HasPrefix(t, "I")
			dbnullable := nullable || t == "string"

			codeType := baseType.CodeType
			codeDefault := baseType.CodeDefault
			dbType := baseType.DbType
			dbDefault := baseType.DbDefault
			if nullable {
				codeDefault = null
				codeType = fmt.Sprintf(nullableFormat, codeType)
			}

			if dbnullable {
				dbDefault = dbnull
			}

			p := parsedField{
				t:            t,
				name:         strings.TrimSuffix(m[2], "Field"),
				codenullable: nullable,
				dbnullable:   dbnullable,
				collection:   collection,
				isInterface:  isInterface,
				flags:        flagstr,
				codeType:     codeType,
				codeDefault:  codeDefault,
				dbType:       dbType,
				dbDefault:    dbDefault,
			}

			plist = append(plist, p)
		}
	}

	return genflags, plist, nil
}
