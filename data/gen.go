package data

import (
	"fmt"
	"maps"
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

	maps.Copy(ngf.SpecifiedFlags, gf.SpecifiedFlags)
	maps.Copy(ngf.Skip, gf.Skip)
	maps.Copy(ngf.Custom, gf.Custom)

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
	ss := strings.Split(flagstr, ",")
	for _, s := range ss {
		err := fs.SetFlag(s)
		if err != nil {
			return GenFlags{}, fmt.Errorf("parsing flag %s. %w", s, err)
		}
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
