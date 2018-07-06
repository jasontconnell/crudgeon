package data

type GenPackage struct {
	Name         string
	Path         string
	OutputFile   string
	Fields       []GenField
	TemplateFile string
	Prefix       string
}

type GenField struct {
	Name     string
	Type     string
	Nullable bool
}
