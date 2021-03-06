package process

import (
	"fmt"
	"strings"

	"github.com/jasontconnell/crudgeon/data"
)

func parseFieldFlags(instructions string) (data.FieldFlags, error) {
	flags := data.FieldFlags{}
	ss := strings.Split(instructions, ",")
	for _, s := range ss {
		flg := s[0] == '+'
		if !flg && s[0] != '-' {
			return flags, fmt.Errorf("Need + or - as first character for flags, %s: %s", instructions, s)
		}

		flds := strings.Fields(string(s[1:]))
		switch flds[0] {
		case "sqlignore":
			flags.SqlIgnore = flg
		case "jsonignore":
			flags.JsonIgnore = flg
		case "csignore":
			flags.CsIgnore = flg
		case "key":
			flags.Key = flg
		case "index":
			flags.Index = flg
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
		case "forcesql":
			flags.ForceSql = flg
			if len(flds) == 1 {
				return flags, fmt.Errorf("forcesql flag must provide sql type (+forcesql sqltype)")
			}
			flags.ForceSqlType = flds[1]
		default:
			return flags, fmt.Errorf("Invalid flag: %s", flds)
		}
	}
	return flags, nil
}
