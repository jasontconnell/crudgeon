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

var fldreg *regexp.Regexp = regexp.MustCompile(`^([a-zA-Z0-9\*\[\]_]+) (.*?)( *//[0-9a-zA-Z\+\-,\."\/_\(\) ]+)?$`)
var globalflagsreg *regexp.Regexp = regexp.MustCompile(`^/{2}([\+\-a-zA-Z_,0-9\/_ ]*?)$`)

type ParsedFile struct {
	Path     string
	Fields   []data.Field
	GenFlags data.GenFlags
	Imports  []string
}

type parsedField struct {
	t              string
	name           string
	codenullable   bool
	dbnullable     bool
	collection     bool
	flags          string
	codeType       string
	collectionType string
	codeDefault    string
	dbType         string
	dbDefault      string
	include        string
}

func ParseFile(file string, baseTypes map[string]data.MappedType, null, dbnull string) (ParsedFile, error) {
	contents, err := os.ReadFile(file)
	parsed := ParsedFile{Path: file}

	mimp := make(map[string]string)

	if err != nil {
		return parsed, err
	}

	flags, fields, err := getParsed(string(contents), baseTypes, null, dbnull)
	if err != nil {
		return parsed, err
	}

	flds := []data.Field{}
	for _, p := range fields {
		if p.t == "" {
			return parsed, fmt.Errorf("Type not found for %v %v", p.t, p.name)
		}

		if _, ok := mimp[p.include]; !ok && p.include != "" {
			mimp[p.include] = p.include
		}

		codenullable := p.codenullable
		_, isBaseType := baseTypes[p.t]
		if !p.codenullable && !isBaseType {
			codenullable = true
		}

		name, field := p.name, p.name
		names := strings.Split(p.name, "|")
		if len(names) == 2 {
			name = names[1]
			field = names[0]
		}

		var fieldFlags data.FieldFlags
		if p.flags != "" {
			fieldFlags, err = data.ParseFieldFlags(p.flags)
			if err != nil {
				return parsed, fmt.Errorf("error parsing %s. %w", p.flags, err)
			}
		}

		flds = append(flds,
			data.Field{
				Type:           p.t,
				Name:           name,
				FieldName:      field,
				Nullable:       codenullable,
				Collection:     p.collection,
				CollectionType: p.collectionType,
				DbType:         p.dbType,
				IsBaseType:     isBaseType,
				CodeType:       p.codeType,
				CodeDefault:    p.codeDefault,
				DbDefault:      p.dbDefault,
				Flags:          fieldFlags,
				Include:        p.include,
			})
	}

	parsed.Fields = flds
	parsed.GenFlags = flags
	if len(mimp) > 0 {
		for k := range mimp {
			parsed.Imports = append(parsed.Imports, k)
		}
	}

	return parsed, nil
}

func getParsed(c string, baseTypes map[string]data.MappedType, null, dbnull string) (data.GenFlags, []parsedField, error) {
	plist := []parsedField{}

	fs := data.NewFlagSetter()

	s := bufio.NewScanner(bytes.NewBufferString(c))

	for s.Scan() {
		line := s.Text()
		globflags := globalflagsreg.FindAllStringSubmatch(line, -1)
		for _, m := range globflags {
			ss := strings.Split(m[1], ",")
			for _, s := range ss {
				perr := fs.SetFlag(s)
				// genflags, perr = data.ParseFlags(m[1])
				if perr != nil {
					return data.GenFlags{}, nil, perr
				}
			}
		}

		matches := fldreg.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			t := m[1]

			var nullable, collection bool

			if strings.HasSuffix(t, "*") {
				nullable = true
				t = strings.TrimSuffix(t, "*")
			}

			if strings.HasSuffix(t, "[]") {
				collection = true
				t = strings.TrimSuffix(t, "[]")
			}

			var codeType, codeDefault, dbType, dbDefault, colltype, include string
			if baseType, ok := baseTypes[t]; ok && !collection {
				codeType = baseType.CodeType
				codeDefault = baseType.CodeDefault
				dbType = baseType.DbType
				dbDefault = baseType.DbDefault
				include = baseType.Import
			} else if collection {
				colltype = t
			}
			flagstr := strings.TrimPrefix(m[3], " //")

			dbnullable := nullable || t == "string"

			if nullable {
				codeDefault = null
				// codeType = fmt.Sprintf(nullableFormat, codeType)
			}

			if dbnullable {
				dbDefault = dbnull
			}

			p := parsedField{
				t:              t,
				name:           strings.TrimSuffix(m[2], "Field"),
				codenullable:   nullable,
				dbnullable:     dbnullable,
				collection:     collection,
				collectionType: colltype,
				flags:          flagstr,
				codeType:       codeType,
				codeDefault:    codeDefault,
				dbType:         dbType,
				dbDefault:      dbDefault,
				include:        include,
			}

			plist = append(plist, p)
		}
	}

	return fs.GetFlags(), plist, nil
}
