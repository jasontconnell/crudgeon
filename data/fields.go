package data

import (
	"fmt"
	"strings"
)

type Field struct {
	FieldName    string     `json:"field"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	ConcreteType string     `json:"concreteType"`
	Nullable     bool       `json:"nullable"`
	Collection   bool       `json:"collection"`
	IsInterface  bool       `json:"interface"`
	IsBaseType   bool       `json:"baseType"`
	Flags        FieldFlags `json:"flags"`
	CodeType     string     `json:"codeType"`
	CodeDefault  string     `json:"codeDefault"`
	DbType       string     `json:"dbType"`
	DbDefault    string     `json:"dbDefault"`
}

func (f Field) String() string {
	return fmt.Sprintf(`
		Field Name     :  %s
		Name           :  %s
		Type           :  %s
		Concrete Type  :  %s
		Nullable       :  %v
		Collection     :  %v
		IsInterface    :  %v
		IsBaseType     :  %v
		DbType        :  %s
	`, f.FieldName, f.Name, f.Type, f.ConcreteType, f.Nullable, f.Collection, f.IsInterface, f.IsBaseType, f.DbType)
}

type FieldFlags struct {
	IsId       bool
	DbIgnore   bool
	JsonIgnore bool
	CodeIgnore bool
	XmlIgnore  bool
	Key        bool
	ForeignKey bool
	Auto       bool
	Index      bool
	NoMap      bool

	HashIgnore bool

	XmlWrapper        bool
	XmlWrapperElement string

	ParseFromString         bool
	ParseFromStringProperty string
	ParseFromStringFormat   string
	ParseFromStringDefault  string

	ForceDb     bool
	ForceDbType string

	ReadOnly bool

	Custom         map[string]CustomFlag
	SpecifiedFlags map[string]bool
}

func (f FieldFlags) GetFlagSpecified(name string) bool {
	if f.SpecifiedFlags == nil {
		return false
	}
	_, ok := f.SpecifiedFlags[name]
	return ok
}

func (f FieldFlags) GetFlagValue(name string) bool {
	if f.SpecifiedFlags == nil {
		return false
	}
	return f.SpecifiedFlags[name]
}

func ParseFieldFlags(instructions string) (FieldFlags, error) {
	flags := FieldFlags{SpecifiedFlags: make(map[string]bool)}
	ss := smartSplit(instructions, ',')
	for _, s := range ss {
		flg := s[0] == '+'
		if !flg && s[0] != '-' {
			return flags, fmt.Errorf("Need + or - as first character for flags, %s: %s", instructions, s)
		}

		flds := strings.Fields(string(s[1:]))

		flags.SpecifiedFlags[flds[0]] = flg

		switch flds[0] {
		case DbIgnoreFlag:
			flags.DbIgnore = flg
		case JsonIgnoreFlag:
			flags.JsonIgnore = flg
		case CodeIgnoreFlag:
			flags.CodeIgnore = flg
		case KeyFlag:
			flags.Key = flg
		case ForeignKeyFlag:
			flags.ForeignKey = flg
		case AutoFlag:
			flags.Auto = flg
		case IndexFlag:
			flags.Index = flg
		case HashIgnoreFlag:
			flags.HashIgnore = flg
		case NoMapFlag:
			flags.NoMap = flg
		case XmlIgnoreFlag:
			flags.XmlIgnore = flg
		case XmlWrapperFlag:
			flags.XmlWrapper = flg
			if len(flds) == 1 {
				return flags, fmt.Errorf("Xml wrapper flag must provide xml wrapper name (+xmlwrapper XmlWrapperName)")
			}
			flags.XmlWrapperElement = flds[1]
		case ParseFromStringFlag:
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
		case ForceDbFlag:
			flags.ForceDb = flg
			if len(flds) == 1 {
				return flags, fmt.Errorf("forcedb flag must provide db type (+forcedb dbtype)")
			}
			flags.ForceDbType = flds[1]
		default:
			if flags.Custom == nil {
				flags.Custom = make(map[string]CustomFlag)
			}
			val := ""
			if len(flds) > 1 {
				val = strings.Join(flds[1:], " ")
			}
			cf := CustomFlag{Name: flds[0], Value: val, Flag: flg}
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

type MappedType struct {
	CodeType    string
	DbType      string
	CodeDefault string
	DbDefault   string
}
