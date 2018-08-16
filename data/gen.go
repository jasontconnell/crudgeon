package data

type GenPackage struct {
	Name              string
	Path              string
	OutputFile        string
	Namespace         string
	Fields            []GenField
	ConstructorFields []GenField
	KeyFields         []GenField
	TemplateFile      string
	Prefix            string
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
}

type GenFlags struct {
	Id          bool
	Fields      bool
	Collections bool
	Concretes   bool
	Constructor bool
	Keys        bool
}
