package data

type Field struct {
	FieldName    string     `json:"field"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	ConcreteType string     `json:"concreteType"`
	Nullable     bool       `json:"nullable"`
	Collection   bool       `json:"collection"`
	IsInterface  bool       `json:"-"`
	Flags        FieldFlags `json:"-"`
}

type FieldFlags struct {
	IsId       bool
	SqlIgnore  bool
	JsonIgnore bool
	CsIgnore   bool
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
