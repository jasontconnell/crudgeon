package data

import (
	"fmt"
	"strings"
)

type CustomFlag struct {
	Name  string
	Flag  bool
	Value string
}

type GenPackage struct {
	Generate          bool
	Name              string
	Path              string
	OutputFile        string
	Namespace         string
	Fields            []GenField
	ConstructorFields []GenField
	KeyFields         []GenField
	PrimaryKeyFields  []GenField
	UpdateFields      []GenField
	TemplateFile      string
	Prefix            string
	Suffix            string
	Flags             GenFlags
}

type GenField struct {
	Access              string
	FieldName           string
	Name                string
	Type                string
	ElementType         string
	ConcreteType        string
	ConcreteElementType string
	ConcreteProperty    string
	Nullable            bool
	CodeIgnore          bool
	DbIgnore            bool
	JsonIgnore          bool
	XmlIgnore           bool
	IsInterface         bool
	Collection          bool
	Key                 bool
	ForeignKey          bool
	IsBaseType          bool

	CodeType    string
	CodeDefault string
	DbType      string
	DbDefault   string

	HashIgnore bool

	Flags             FieldFlags
	XmlWrapper        bool
	XmlWrapperType    string
	XmlWrapperName    string
	XmlWrapperElement string
	Id                bool
}

type GenFlags struct {
	Id             bool
	Fields         bool
	Collections    bool
	Concretes      bool
	Constructor    bool
	Keys           bool
	PrimaryKeys    bool
	Updates        bool
	DbIgnore       bool
	Merge          bool
	CodeIgnore     bool
	JsonIgnore     bool
	XmlIgnore      bool
	XmlRoot        bool
	XmlRootName    string
	HashIgnore     bool
	Class          bool
	ClassName      string
	Table          bool
	TableName      string
	HasNamespace   bool
	Namespace      string
	ExactName      bool
	HasSkip        bool
	Skip           map[string]bool
	Custom         map[string]CustomFlag
	SpecifiedFlags map[string]bool
}

func (gf *GenFlags) MergeParse(flagstr string) error {
	if gf.SpecifiedFlags == nil {
		gf.SpecifiedFlags = make(map[string]bool)
	}
	ss := strings.Split(flagstr, ",")
	for _, s := range ss {
		flg := s[0] == '+'
		if !flg && s[0] != '-' {
			return fmt.Errorf("Need + or - as first character for flag, %s ... %s", flagstr, s)
		}

		flds := strings.Fields(string(s[1:]))

		gf.SpecifiedFlags[flds[0]] = flg

		switch flds[0] {
		case IdFlag:
			gf.Id = flg
		case FieldsFlag:
			gf.Fields = flg
		case CollectionsFlag:
			gf.Collections = flg
		case ConstructorFlag:
			gf.Constructor = flg
		case ConcretesFlag:
			gf.Concretes = flg
		case KeysFlag:
			gf.Keys = flg
		case PrimaryKeysFlag:
			gf.PrimaryKeys = flg
		case UpdatesFlag:
			gf.Updates = flg
		case DbIgnoreFlag:
			gf.DbIgnore = flg
		case MergeFlag:
			gf.Merge = flg
		case CodeIgnoreFlag:
			gf.CodeIgnore = flg
		case JsonIgnoreFlag:
			gf.JsonIgnore = flg
		case XmlIgnoreFlag:
			gf.XmlIgnore = flg
		case HashIgnoreFlag:
			gf.HashIgnore = flg
		case XmlRootFlag:
			gf.XmlRoot = flg
			if len(flds) == 1 {
				return fmt.Errorf("Xml root flag must provide xml root name (+xmlroot XmlRootName)")
			}
			gf.XmlRootName = flds[1]
		case NamespaceFlag:
			gf.HasNamespace = flg
			if len(flds) == 1 {
				return fmt.Errorf("Namespace flag needs a namespace (+namespace LocalNamespace)")
			}
			gf.Namespace = flds[1]
		case ClassFlag:
			gf.Class = flg
			if len(flds) == 1 {
				return fmt.Errorf("Class root flag must provide the class name (+class ClassName)")
			}
			gf.ClassName = flds[1]
		case TableFlag:
			gf.Table = flg
			if len(flds) == 1 {
				return fmt.Errorf("Table root flag must provide the class name (+table TableName)")
			}
			gf.TableName = strings.Join(flds[1:], " ") // tables can have spaces...
		case ExactFlag:
			gf.ExactName = flg
		case SkipFlag:
			gf.HasSkip = flg
			if gf.Skip == nil {
				gf.Skip = make(map[string]bool)
			}
			gf.Skip[flds[1]] = flg
		default:
			if gf.Custom == nil {
				gf.Custom = make(map[string]CustomFlag)
			}
			val := ""
			if len(flds) > 1 {
				val = flds[1]
			}
			cf := CustomFlag{Name: flds[0], Value: val, Flag: flg}
			gf.Custom[cf.Name] = cf
		}
	}
	return nil
}

func (gf GenFlags) IsFlagSpecified(name string) bool {
	if gf.SpecifiedFlags == nil {
		return false
	}
	_, ok := gf.SpecifiedFlags[name]
	return ok
}

func (gf GenFlags) GetFlagValue(name string) bool {
	if gf.SpecifiedFlags == nil {
		return false
	}
	return gf.SpecifiedFlags[name]
}

func (gf GenFlags) String() string {
	return fmt.Sprintf(`
		Id:          %v
		Fields:      %v
		Collections: %v
		Concretes:   %v
		Constructor: %v
		Keys:        %v
		DbIgnore:   %v
		CodeIgnore:    %v
		JsonIgnore:  %v
		XmlIgnore:   %v
		XmlRoot:     %v (%v)
		Class:       %s
		Exact:       %v
	`, gf.Id, gf.Fields, gf.Collections, gf.Concretes, gf.Constructor, gf.Keys, gf.DbIgnore, gf.CodeIgnore, gf.JsonIgnore, gf.XmlIgnore, gf.XmlRoot, gf.XmlRootName, gf.ClassName, gf.ExactName)
}
