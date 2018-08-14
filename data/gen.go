package data

type GenPackage struct {
	Name              string
	Path              string
	OutputFile        string
	Fields            []GenField
	ConstructorFields []GenField
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
}
