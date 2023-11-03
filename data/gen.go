package data

import (
	"fmt"
	"strings"
)

type CustomFlag struct {
	Name  string
	Flag  bool
	Value string
}

type GenPackage struct {
	Generate          bool
	Name              string
	Path              string
	OutputFile        string
	Namespace         string
	Fields            []GenField
	ConstructorFields []GenField
	KeyFields         []GenField
	TemplateFile      string
	Prefix            string
	Suffix            string
	Flags             GenFlags
}

type GenField struct {
	Access              string
	FieldName           string
	Name                string
	Type                string
	ElementType         string
	ConcreteType        string
	ConcreteElementType string
	ConcreteProperty    string
	Nullable            bool
	CsIgnore            bool
	SqlIgnore           bool
	JsonIgnore          bool
	XmlIgnore           bool
	IsInterface         bool
	Collection          bool
	Key                 bool
	IsBaseType          bool

	HashIgnore bool

	Flags             FieldFlags
	XmlWrapper        bool
	XmlWrapperType    string
	XmlWrapperName    string
	XmlWrapperElement string
	Id                bool
}

type GenFlags struct {
	Id           bool
	Fields       bool
	Collections  bool
	Concretes    bool
	Constructor  bool
	Keys         bool
	SqlIgnore    bool
	CsIgnore     bool
	JsonIgnore   bool
	XmlIgnore    bool
	XmlRoot      bool
	XmlRootName  string
	HashIgnore   bool
	Class        bool
	ClassName    string
	HasNamespace bool
	Namespace    string
	ExactName    bool
	HasSkip      bool
	Skip         map[string]bool
	Custom       map[string]CustomFlag
}

func (gf *GenFlags) MergeParse(flagstr string) error {
	ss := strings.Split(flagstr, ",")
	for _, s := range ss {
		flg := s[0] == '+'
		if !flg && s[0] != '-' {
			return fmt.Errorf("Need + or - as first character for flag, %s ... %s", flagstr, s)
		}

		flds := strings.Fields(string(s[1:]))

		switch flds[0] {
		case "id":
			gf.Id = flg
		case "fields":
			gf.Fields = flg
		case "collections":
			gf.Collections = flg
		case "constructor":
			gf.Constructor = flg
		case "concretes":
			gf.Concretes = flg
		case "keys":
			gf.Keys = flg
		case "sqlignore":
			gf.SqlIgnore = flg
		case "csignore":
			gf.CsIgnore = flg
		case "jsonignore":
			gf.JsonIgnore = flg
		case "xmlignore":
			gf.XmlIgnore = flg
		case "hashignore":
			gf.HashIgnore = flg
		case "xmlroot":
			gf.XmlRoot = flg
			if len(flds) == 1 {
				return fmt.Errorf("Xml root flag must provide xml root name (+xmlroot XmlRootName)")
			}
			gf.XmlRootName = flds[1]
		case "namespace":
			gf.HasNamespace = flg
			if len(flds) == 1 {
				return fmt.Errorf("Namespace flag needs a namespace (+namespace LocalNamespace)")
			}
			gf.Namespace = flds[1]
		case "class":
			gf.Class = flg
			if len(flds) == 1 {
				return fmt.Errorf("Class root flag must provide the class name (+class ClassName)")
			}
			gf.ClassName = flds[1]
		case "exact":
			gf.ExactName = flg
		case "skip":
			gf.HasSkip = flg
			if gf.Skip == nil {
				gf.Skip = make(map[string]bool)
			}
			gf.Skip[flds[1]] = flg
		default:
			if gf.Custom == nil {
				gf.Custom = make(map[string]CustomFlag)
			}
			val := ""
			if len(flds) > 1 {
				val = flds[1]
			}
			cf := CustomFlag{Name: flds[0], Value: val, Flag: flg}
			gf.Custom[cf.Name] = cf
		}
	}
	return nil
}

func (gf GenFlags) String() string {
	return fmt.Sprintf(`
		Id:          %v
		Fields:      %v
		Collections: %v
		Concretes:   %v
		Constructor: %v
		Keys:        %v
		SqlIgnore:   %v
		CsIgnore:    %v
		JsonIgnore:  %v
		XmlIgnore:   %v
		XmlRoot:     %v (%v)
		Class:       %s
		Exact:       %v
	`, gf.Id, gf.Fields, gf.Collections, gf.Concretes, gf.Constructor, gf.Keys, gf.SqlIgnore, gf.CsIgnore, gf.JsonIgnore, gf.XmlIgnore, gf.XmlRoot, gf.XmlRootName, gf.ClassName, gf.ExactName)
}
