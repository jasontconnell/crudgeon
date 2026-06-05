package data

import (
	"fmt"
)

type Field struct {
	FieldName      string     `json:"field"`
	Name           string     `json:"name"`
	Type           string     `json:"type"`
	Nullable       bool       `json:"nullable"`
	Collection     bool       `json:"collection"`
	CollectionType string     `json:"collectionType"`
	IsBaseType     bool       `json:"baseType"`
	Flags          FieldFlags `json:"flags"`
	CodeType       string     `json:"codeType"`
	CodeDefault    string     `json:"codeDefault"`
	DbType         string     `json:"dbType"`
	DbDefault      string     `json:"dbDefault"`
	Include        string     `json:"include"`
}

func (f Field) String() string {
	return fmt.Sprintf(`
		Field Name      :  %s
		Name            :  %s
		Type            :  %s
		Nullable        :  %v
		Collection      :  %v
		Collection Type :  %v
		IsBaseType      :  %v
		DbType          :  %s
	`, f.FieldName, f.Name, f.Type, f.Nullable, f.Collection, f.CollectionType, f.IsBaseType, f.DbType)
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
	pflags, err := ParseFlagsRaw(instructions)
	if err != nil {
		return flags, err
	}
	for _, f := range pflags {
		flags.SpecifiedFlags[f.Name] = true

		switch f.Name {
		case DbIgnoreFlag:
			flags.DbIgnore = f.Switch
		case JsonIgnoreFlag:
			flags.JsonIgnore = f.Switch
		case CodeIgnoreFlag:
			flags.CodeIgnore = f.Switch
		case KeyFlag:
			flags.Key = f.Switch
		case ForeignKeyFlag:
			flags.ForeignKey = f.Switch
		case AutoFlag:
			flags.Auto = f.Switch
		case IndexFlag:
			flags.Index = f.Switch
		case HashIgnoreFlag:
			flags.HashIgnore = f.Switch
		case NoMapFlag:
			flags.NoMap = f.Switch
		case XmlIgnoreFlag:
			flags.XmlIgnore = f.Switch
		case XmlWrapperFlag:
			flags.XmlWrapper = f.Switch
			if len(f.Values) == 0 {
				return flags, fmt.Errorf("Xml wrapper flag must provide xml wrapper name (+xmlwrapper XmlWrapperName)")
			}
			flags.XmlWrapperElement = f.Value
		case ParseFromStringFlag:
			flags.ParseFromString = f.Switch
			if len(f.Values) == 0 {
				return flags, fmt.Errorf("parse from string flag must provide string property name (+parsefromstring StringProperty DefaultVal Format)")
			}

			flags.ParseFromStringProperty = f.Values[1]
			if len(f.Values) > 2 {
				flags.ParseFromStringDefault = f.Values[2]
			}
			if len(f.Values) > 3 {
				flags.ParseFromStringFormat = f.Values[3]
			}

			flags.ReadOnly = true
		case ForceDbFlag:
			flags.ForceDb = f.Switch
			if len(f.Values) == 1 {
				return flags, fmt.Errorf("forcedb flag must provide db type (+forcedb dbtype)")
			}
			flags.ForceDbType = f.Values[1]
		default:
			if flags.Custom == nil {
				flags.Custom = make(map[string]CustomFlag)
			}

			cf := CustomFlag{Name: f.Name, Value: f.Value, Flag: f.Switch}
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
	Import      string
}
