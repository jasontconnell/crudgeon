package process

import (
	"fmt"
	"lpgagen/data"
	"strings"
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
				return flags, fmt.Errorf("Xml root flag must provide xml root name (+xmlroot XmlRootName)")
			}
			flags.XmlWrapperElement = flds[1]
		default:
			return flags, fmt.Errorf("Invalid flag: %s", flds)
		}
	}
	return flags, nil
}

func parseGenFlags(flagstr string) (data.GenFlags, error) {
	flags := data.GenFlags{}
	ss := strings.Split(flagstr, ",")
	for _, s := range ss {
		flg := s[0] == '+'
		if !flg && s[0] != '-' {
			return flags, fmt.Errorf("Need + or - as first character for flag, %s ... %s", flagstr, s)
		}

		flds := strings.Fields(string(s[1:]))

		switch flds[0] {
		case "id":
			flags.Id = flg
		case "fields":
			flags.Fields = flg
		case "collections":
			flags.Collections = flg
		case "constructor":
			flags.Constructor = flg
		case "concretes":
			flags.Concretes = flg
		case "keys":
			flags.Keys = flg
		case "sqlignore":
			flags.SqlIgnore = flg
		case "csignore":
			flags.CsIgnore = flg
		case "jsonignore":
			flags.JsonIgnore = flg
		case "xmlignore":
			flags.XmlIgnore = flg
		case "xmlroot":
			flags.XmlRoot = flg
			if len(flds) == 1 {
				return flags, fmt.Errorf("Xml root flag must provide xml root name (+xmlroot XmlRootName)")
			}
			flags.XmlRootName = flds[1]
		default:
			return flags, fmt.Errorf("Invalid flags: %s", flds)
		}
	}
	return flags, nil
}
