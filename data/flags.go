package data

type Flags interface {
	IsFlagSpecified(name string) bool
	GetFlagValue(name string) bool
}

const (
	IdFlag              = "id"
	IdUpdateFlag        = "idupdate"
	AutoFlag            = "auto"
	FieldsFlag          = "fields"
	CollectionsFlag     = "collections"
	ConstructorFlag     = "constructor"
	IndexFlag           = "index"
	KeyFlag             = "key"
	KeysFlag            = "keys"
	PrimaryKeysFlag     = "primarykeys"
	ForeignKeyFlag      = "foreignkey"
	UpdatesFlag         = "updates"
	DbIgnoreFlag        = "dbignore"
	ForceDbFlag         = "forcedb"
	CodeIgnoreFlag      = "codeignore"
	JsonIgnoreFlag      = "jsonignore"
	XmlIgnoreFlag       = "xmlignore"
	HashIgnoreFlag      = "hashignore"
	NoMapFlag           = "nomap"
	XmlRootFlag         = "xmlroot"
	XmlWrapperFlag      = "xmlwrapper"
	NamespaceFlag       = "namespace"
	ClassFlag           = "class"
	TableFlag           = "table"
	ExactFlag           = "exact"
	MergeFlag           = "merge"
	SkipFlag            = "skip"
	ParseFromStringFlag = "parsefromstring"
	DatabaseFlag        = "database"
)

var AllFlags []string = []string{
	IdFlag, IdUpdateFlag, AutoFlag, FieldsFlag, CollectionsFlag, ConstructorFlag, IndexFlag, KeyFlag, KeysFlag, PrimaryKeysFlag, ForeignKeyFlag,
	UpdatesFlag, DbIgnoreFlag, ForceDbFlag, CodeIgnoreFlag, JsonIgnoreFlag, XmlIgnoreFlag, HashIgnoreFlag, NoMapFlag, XmlRootFlag, XmlWrapperFlag,
	NamespaceFlag, ClassFlag, TableFlag, ExactFlag, MergeFlag, SkipFlag, ParseFromStringFlag, DatabaseFlag,
}
