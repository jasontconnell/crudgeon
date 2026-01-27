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
	Generate bool
	Object   GenObject
	Objects  []GenObject

	Namespace        string
	Imports          []string
	Path             string
	Ext              string
	TemplateFile     string
	FilenameTemplate string
	Flags            GenFlags
	OneFile          bool
}

type GenObject struct {
	Name              string
	NameLower         string
	Fields            []GenField
	ConstructorFields []GenField
	KeyFields         []GenField
	PrimaryKeyFields  []GenField
	UpdateFields      []GenField
	Namespace         string
}

type GenField struct {
	FieldName   string
	Name        string
	NameLower   string
	Type        string
	ElementType string
	Nullable    bool
	CodeIgnore  bool
	DbIgnore    bool
	JsonIgnore  bool
	XmlIgnore   bool
	Collection  bool
	Key         bool
	ForeignKey  bool
	IsBaseType  bool

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

	Include string
}

type GenFlags struct {
	Id                 bool   `flag:"id"`
	IdUpdate           bool   `flag:"idupdate"`
	Fields             bool   `flag:"fields"`
	Collections        bool   `flag:"collections"`
	CollectionTemplate string `flag:"collectionTemplate"`
	Constructor        bool   `flag:"constructor"`
	Keys               bool   `flag:"keys"`
	PrimaryKeys        bool   `flag:"primarykeys"`
	Updates            bool   `flag:"updates"`
	DbIgnore           bool   `flag:"dbignore"`
	Merge              bool   `flag:"merge"`
	CodeIgnore         bool   `flag:"codeignore"`
	JsonIgnore         bool   `flag:"jsonignore"`
	XmlIgnore          bool   `flag:"xmlignore"`
	XmlRoot            bool   `flag:"xmlroot" value:"XmlRootName"`
	XmlRootName        string
	HashIgnore         bool `flag:"hashignore"`
	Class              bool `flag:"class" value:"ClassName"`
	ClassName          string
	Table              bool `flag:"table" value:"TableName"`
	TableName          string
	Database           bool `flag:"database"`
	HasNamespace       bool `flag:"namespace" value:"Namespace"`
	Namespace          string
	Exact              bool `flag:"exact"`
	HasSkip            bool `flag:"hasskip"`
	Skip               map[string]bool
	Custom             map[string]CustomFlag
	SpecifiedFlags     map[string]bool
}

func mergeString(left, right string) string {
	if left == "" && right != "" {
		return right
	} else if left != "" && right == "" {
		return left
	}
	return left
}

func mergeBool(gf GenFlags, f GenFlags, prop string, def bool) bool {
	gfval, ingf := gf.SpecifiedFlags[prop]
	fval, inf := f.SpecifiedFlags[prop]

	val := def

	if ingf && !inf {
		val = gfval
	} else if !ingf && inf {
		val = fval
	}

	return val
}

func MergeGenFlags(gf GenFlags, f GenFlags) GenFlags {
	ngf := GenFlags{
		Skip:           make(map[string]bool),
		SpecifiedFlags: make(map[string]bool),
		Custom:         make(map[string]CustomFlag),
	}

	for k, v := range gf.SpecifiedFlags {
		ngf.SpecifiedFlags[k] = v
	}

	for k, v := range f.SpecifiedFlags {
		ngf.SpecifiedFlags[k] = v
	}

	for k, v := range gf.Skip {
		ngf.Skip[k] = v
	}

	for k, v := range f.Skip {
		ngf.Skip[k] = v
	}

	for k, v := range gf.Custom {
		ngf.Custom[k] = v
	}

	for k, v := range f.Custom {
		ngf.Custom[k] = v
	}

	ngf.Id = mergeBool(gf, f, IdFlag, false)
	ngf.IdUpdate = mergeBool(gf, f, IdUpdateFlag, false)
	ngf.Fields = mergeBool(gf, f, FieldsFlag, false)
	ngf.Collections = mergeBool(gf, f, CollectionsFlag, false)
	ngf.CollectionTemplate = gf.CollectionTemplate
	ngf.Constructor = mergeBool(gf, f, ConstructorFlag, false)
	ngf.Keys = mergeBool(gf, f, KeysFlag, false)
	ngf.PrimaryKeys = mergeBool(gf, f, PrimaryKeysFlag, false)
	ngf.Updates = mergeBool(gf, f, UpdatesFlag, false)
	ngf.DbIgnore = mergeBool(gf, f, DbIgnoreFlag, false)
	ngf.Merge = mergeBool(gf, f, MergeFlag, false)
	ngf.CodeIgnore = mergeBool(gf, f, CodeIgnoreFlag, false)
	ngf.JsonIgnore = mergeBool(gf, f, JsonIgnoreFlag, false)
	ngf.XmlIgnore = mergeBool(gf, f, XmlIgnoreFlag, false)
	ngf.XmlRoot = mergeBool(gf, f, XmlRootFlag, false)
	ngf.XmlRootName = mergeString(gf.XmlRootName, f.XmlRootName)
	ngf.HashIgnore = mergeBool(gf, f, HashIgnoreFlag, false)
	ngf.Class = mergeBool(gf, f, ClassFlag, false)
	ngf.ClassName = mergeString(gf.ClassName, f.ClassName)
	ngf.Table = mergeBool(gf, f, TableFlag, false)
	ngf.TableName = mergeString(gf.TableName, f.TableName)
	ngf.Database = mergeBool(gf, f, DatabaseFlag, gf.Database || f.Database)
	ngf.HasNamespace = mergeBool(gf, f, HasNamespaceFlag, false)
	ngf.Namespace = mergeString(gf.Namespace, f.Namespace)
	ngf.Exact = mergeBool(gf, f, ExactFlag, false)
	ngf.HasSkip = mergeBool(gf, f, SkipFlag, false)

	return ngf
}

func ParseFlags(flagstr string) (GenFlags, error) {
	fs := NewFlagSetter()
	// if gf.SpecifiedFlags == nil {
	// 	gf.SpecifiedFlags = make(map[string]bool)
	// }
	ss := strings.Split(flagstr, ",")
	for _, s := range ss {
		err := fs.SetFlag(s)
		if err != nil {
			return GenFlags{}, fmt.Errorf("parsing flag %s. %w", s, err)
		}
		// flg := s[0] == '+'
		// if !flg && s[0] != '-' {
		// 	return fmt.Errorf("Need + or - as first character for flag, %s ... %s", flagstr, s)
		// }

		// flds := strings.Fields(s[1:])

		// // if FlagTypes[flds[0]] == String {
		// // 	sval := strings.Join(flds[1:], " ")
		// // }

		// gf.SpecifiedFlags[flds[0]] = flg

		// switch flds[0] {
		// case IdFlag:
		// 	gf.Id = flg
		// case IdUpdateFlag:
		// 	gf.IdUpdate = flg
		// case FieldsFlag:
		// 	gf.Fields = flg
		// case CollectionsFlag:
		// 	gf.Collections = flg
		// case ConstructorFlag:
		// 	gf.Constructor = flg
		// case KeysFlag:
		// 	gf.Keys = flg
		// case PrimaryKeysFlag:
		// 	gf.PrimaryKeys = flg
		// case UpdatesFlag:
		// 	gf.Updates = flg
		// case DbIgnoreFlag:
		// 	gf.DbIgnore = flg
		// case MergeFlag:
		// 	gf.Merge = flg
		// case CodeIgnoreFlag:
		// 	gf.CodeIgnore = flg
		// case JsonIgnoreFlag:
		// 	gf.JsonIgnore = flg
		// case XmlIgnoreFlag:
		// 	gf.XmlIgnore = flg
		// case HashIgnoreFlag:
		// 	gf.HashIgnore = flg
		// case XmlRootFlag:
		// 	gf.XmlRoot = flg
		// 	if len(flds) == 1 {
		// 		return fmt.Errorf("Xml root flag must provide xml root name (+xmlroot XmlRootName)")
		// 	}
		// 	gf.XmlRootName = flds[1]
		// 	gf.SpecifiedFlags[XmlRootNameFlag] = flg
		// case NamespaceFlag:
		// 	gf.HasNamespace = flg
		// 	if len(flds) == 1 {
		// 		return fmt.Errorf("Namespace flag needs a namespace (+namespace LocalNamespace)")
		// 	}
		// 	gf.Namespace = flds[1]
		// 	gf.SpecifiedFlags[HasNamespaceFlag] = flg
		// case ClassFlag:
		// 	gf.Class = flg
		// 	if len(flds) == 1 {
		// 		return fmt.Errorf("Class root flag must provide the class name (+class ClassName)")
		// 	}
		// 	gf.ClassName = flds[1]
		// 	gf.SpecifiedFlags[ClassNameFlag] = flg
		// case TableFlag:
		// 	gf.Table = flg
		// 	if len(flds) == 1 {
		// 		return fmt.Errorf("Table root flag must provide the class name (+table TableName)")
		// 	}
		// 	gf.TableName = strings.Join(flds[1:], " ") // tables can have spaces...
		// 	gf.SpecifiedFlags[TableNameFlag] = flg
		// case ExactFlag:
		// 	gf.Exact = flg
		// case SkipFlag:
		// 	gf.HasSkip = flg
		// 	if gf.Skip == nil {
		// 		gf.Skip = make(map[string]bool)
		// 	}
		// 	gf.Skip[flds[1]] = flg
		// default:
		// 	if gf.Custom == nil {
		// 		gf.Custom = make(map[string]CustomFlag)
		// 	}
		// 	val := ""
		// 	if len(flds) > 1 {
		// 		val = flds[1]
		// 	}
		// 	cf := CustomFlag{Name: flds[0], Value: val, Flag: flg}
		// 	gf.Custom[cf.Name] = cf
		// }
	}
	return fs.GetFlags(), nil
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
		Id:                    %v
		Fields:                %v
		Collections:           %v
		Collection Template:   %v
		Constructor:           %v
		Keys:                  %v
		DbIgnore:              %v
		CodeIgnore:            %v
		JsonIgnore:            %v
		XmlIgnore:             %v
		XmlRoot:               %v (%v)
		Class:                 %s
		Exact:                 %v
	`, gf.Id, gf.Fields, gf.Collections, gf.CollectionTemplate, gf.Constructor, gf.Keys, gf.DbIgnore, gf.CodeIgnore, gf.JsonIgnore, gf.XmlIgnore, gf.XmlRoot, gf.XmlRootName, gf.ClassName, gf.Exact)
}
