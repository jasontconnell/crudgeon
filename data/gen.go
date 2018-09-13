package data

import (
	"fmt"
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
	//Flags               FieldFlags
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
