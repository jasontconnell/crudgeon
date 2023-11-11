package data

import "fmt"

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

	Custom map[string]CustomFlag
}

type MappedType struct {
	CodeType    string
	DbType      string
	CodeDefault string
	DbDefault   string
}
