package data

import (
	"fmt"
	"strings"
)

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

	Flags             FieldFlags
	XmlWrapper        bool
	XmlWrapperType    string
	XmlWrapperName    string
	XmlWrapperElement string
	Id                bool
}

type GenFlags struct {
	Id          bool
	Fields      bool
	Collections bool
	Concretes   bool
	Constructor bool
	Keys        bool
	SqlIgnore   bool
	CsIgnore    bool
	JsonIgnore  bool
	XmlIgnore   bool
	XmlRoot     bool
	XmlRootName string
	Class       bool
	ClassName   string
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
		case "xmlroot":
			gf.XmlRoot = flg
			if len(flds) == 1 {
				return fmt.Errorf("Xml root flag must provide xml root name (+xmlroot XmlRootName)")
			}
			gf.XmlRootName = flds[1]
		case "class":
			gf.Class = flg
			if len(flds) == 1 {
				return fmt.Errorf("Class root flag must provide the class name (+class ClassName)")
			}
			gf.ClassName = flds[1]
		default:
			return fmt.Errorf("Invalid flags: %s", flds)
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
	`, gf.Id, gf.Fields, gf.Collections, gf.Concretes, gf.Constructor, gf.Keys, gf.SqlIgnore, gf.CsIgnore, gf.JsonIgnore, gf.XmlIgnore, gf.XmlRoot, gf.XmlRootName)
}
