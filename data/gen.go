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
	FieldName        string
	Name             string
	Type             string
	ElementType      string
	ConcreteType     string
	ConcreteProperty string
	Nullable         bool
	CsIgnore         bool
	SqlIgnore        bool
	JsonIgnore       bool
	IsInterface      bool
	Collection       bool
	Key              bool
	IsBaseType       bool
	Flags            FieldFlags
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
	`, gf.Id, gf.Fields, gf.Collections, gf.Concretes, gf.Constructor, gf.Keys, gf.SqlIgnore, gf.CsIgnore)
}
