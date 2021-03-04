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
		default:
			return flags, fmt.Errorf("Invalid flag: %s", flds)
		}
	}
	return flags, nil
}
