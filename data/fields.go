package data

type Field struct {
	FieldName    string `json:"field"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	ConcreteType string `json:"concreteType"`
	Nullable     bool   `json:"nullable"`
	Collection   bool   `json:"collection"`
	JsonIgnore   bool   `json:"jsonIgnore"`
	IsInterface  bool   `json:"-"`
}
