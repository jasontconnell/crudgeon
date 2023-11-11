package process

import (
	"fmt"
	"strings"

	"github.com/jasontconnell/crudgeon/data"
)

func parseFieldFlags(instructions string) (data.FieldFlags, error) {
	flags := data.FieldFlags{}
	ss := smartSplit(instructions, ',')
	for _, s := range ss {
		flg := s[0] == '+'
		if !flg && s[0] != '-' {
			return flags, fmt.Errorf("Need + or - as first character for flags, %s: %s", instructions, s)
		}

		flds := strings.Fields(string(s[1:]))
		switch flds[0] {
		case "dbignore":
			flags.DbIgnore = flg
		case "jsonignore":
			flags.JsonIgnore = flg
		case "codeignore":
			flags.CodeIgnore = flg
		case "key":
			flags.Key = flg
		case "foreignkey":
			flags.ForeignKey = flg
		case "auto":
			flags.Auto = flg
		case "index":
			flags.Index = flg
		case "hashignore":
			flags.HashIgnore = flg
		case "nomap":
			flags.NoMap = flg
		case "xmlignore":
			flags.XmlIgnore = flg
		case "xmlwrapper":
			flags.XmlWrapper = flg
			if len(flds) == 1 {
				return flags, fmt.Errorf("Xml wrapper flag must provide xml wrapper name (+xmlwrapper XmlWrapperName)")
			}
			flags.XmlWrapperElement = flds[1]
		case "parsefromstring":
			flags.ParseFromString = flg
			if len(flds) == 1 {
				return flags, fmt.Errorf("parse from string flag must provide string property name (+parsefromstring StringProperty DefaultVal Format)")
			}

			flags.ParseFromStringProperty = flds[1]
			if len(flds) > 2 {
				flags.ParseFromStringDefault = flds[2]
			}
			if len(flds) > 3 {
				flags.ParseFromStringFormat = flds[3]
			}

			flags.ReadOnly = true
		case "forcedb":
			flags.ForceDb = flg
			if len(flds) == 1 {
				return flags, fmt.Errorf("forcedb flag must provide db type (+forcedb dbtype)")
			}
			flags.ForceDbType = flds[1]
		default:
			if flags.Custom == nil {
				flags.Custom = make(map[string]data.CustomFlag)
			}
			val := ""
			if len(flds) > 1 {
				val = strings.Join(flds[1:], " ")
			}
			cf := data.CustomFlag{Name: flds[0], Value: val, Flag: flg}
			flags.Custom[cf.Name] = cf
		}
	}
	return flags, nil
}

func smartSplit(str string, sep rune) []string {
	list := []string{}
	cur := ""
	level := 0
	for i, c := range str {
		switch c {
		case '(':
			level++
		case ')':
			level--
		}

		if (c == sep) && level == 0 {
			list = append(list, cur)
			cur = ""
		} else if i == len(str)-1 {
			list = append(list, cur+string(c))
		} else {
			cur = cur + string(c)
		}
	}
	return list
}
