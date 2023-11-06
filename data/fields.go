package data

type Field struct {
	FieldName    string     `json:"field"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	ConcreteType string     `json:"concreteType"`
	Nullable     bool       `json:"nullable"`
	Collection   bool       `json:"collection"`
	IsInterface  bool       `json:"interface"`
	IsBaseType   bool       `json:"baseType"`
	SqlType      string     `json:"sqlType"`
	Flags        FieldFlags `json:"flags"`
}

type FieldFlags struct {
	IsId       bool
	SqlIgnore  bool
	JsonIgnore bool
	CodeIgnore bool
	XmlIgnore  bool
	Key        bool
	Index      bool
	NoMap      bool

	HashIgnore bool

	XmlWrapper        bool
	XmlWrapperElement string

	ParseFromString         bool
	ParseFromStringProperty string
	ParseFromStringFormat   string
	ParseFromStringDefault  string

	ForceSql     bool
	ForceSqlType string

	ReadOnly bool

	Custom map[string]CustomFlag
}
