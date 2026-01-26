package data

type Flags interface {
	IsFlagSpecified(name string) bool
	GetFlagValue(name string) bool
}

type FlagType int

const (
	Boolean FlagType = iota
	String
)

const (
	IdFlag                 = "id"
	IdUpdateFlag           = "idupdate"
	AutoFlag               = "auto"
	FieldsFlag             = "fields"
	CollectionTemplateFlag = "collectionTemplate"
	CollectionsFlag        = "collections"
	ConstructorFlag        = "constructor"
	IndexFlag              = "index"
	KeyFlag                = "key"
	KeysFlag               = "keys"
	PrimaryKeysFlag        = "primarykeys"
	ForeignKeyFlag         = "foreignkey"
	UpdatesFlag            = "updates"
	DbIgnoreFlag           = "dbignore"
	ForceDbFlag            = "forcedb"
	CodeIgnoreFlag         = "codeignore"
	JsonIgnoreFlag         = "jsonignore"
	XmlIgnoreFlag          = "xmlignore"
	HashIgnoreFlag         = "hashignore"
	NoMapFlag              = "nomap"
	XmlRootFlag            = "xmlroot"
	XmlRootNameFlag        = "xmlrootname"
	XmlWrapperFlag         = "xmlwrapper"
	HasNamespaceFlag       = "hasnamespace"
	NamespaceFlag          = "namespace"
	ClassFlag              = "class"
	ClassNameFlag          = "classname"
	TableFlag              = "table"
	TableNameFlag          = "tablename"
	ExactFlag              = "exact"
	MergeFlag              = "merge"
	SkipFlag               = "skip"
	ParseFromStringFlag    = "parsefromstring"
	DatabaseFlag           = "database"
)

var AllFlags []string = []string{
	IdFlag, IdUpdateFlag, AutoFlag, FieldsFlag, CollectionTemplateFlag, CollectionsFlag, ConstructorFlag, IndexFlag, KeyFlag, KeysFlag, PrimaryKeysFlag, ForeignKeyFlag,
	UpdatesFlag, DbIgnoreFlag, ForceDbFlag, CodeIgnoreFlag, JsonIgnoreFlag, XmlIgnoreFlag, HashIgnoreFlag, NoMapFlag, XmlRootFlag, XmlRootNameFlag, XmlWrapperFlag,
	HasNamespaceFlag, NamespaceFlag, ClassFlag, ClassNameFlag, TableFlag, TableNameFlag, ExactFlag, MergeFlag, SkipFlag, ParseFromStringFlag, DatabaseFlag,
}

var FlagTypes map[string]FlagType

func init() {
	FlagTypes = make(map[string]FlagType)
	for _, t := range AllFlags {
		FlagTypes[t] = Boolean
	}

	FlagTypes[XmlRootNameFlag] = String
	FlagTypes[XmlWrapperFlag] = String
	FlagTypes[NamespaceFlag] = String
	FlagTypes[ClassNameFlag] = String
	FlagTypes[TableNameFlag] = String
	FlagTypes[CollectionTemplateFlag] = String
}
