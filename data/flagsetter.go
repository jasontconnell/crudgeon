package data

import (
	"log"
	"reflect"
)

type FlagSetter struct {
	gf *GenFlags
	t  reflect.Type
	v  reflect.Value
	m  map[string]string
	vm map[string]string
}

func NewFlagSetter() *FlagSetter {
	gf := new(GenFlags)
	return NewFlagSetterWithValue(gf)
}

func NewFlagSetterWithValue(gf *GenFlags) *FlagSetter {
	gf.SpecifiedFlags = make(map[string]bool)
	gf.Custom = make(map[string]CustomFlag)
	gf.Skip = make(map[string]bool)

	fs := &FlagSetter{gf: gf, t: reflect.TypeOf(gf).Elem(), v: reflect.ValueOf(gf)}
	fs.parseTags()
	return fs

}

func (fs *FlagSetter) parseTags() {
	fs.m = make(map[string]string)
	fs.vm = make(map[string]string)
	for i := range fs.t.NumField() {
		fld := fs.t.Field(i)
		flagtag := fld.Tag.Get("flag")
		if flagtag != "" {
			fs.m[flagtag] = fld.Name
		}

		valtag := fld.Tag.Get("value")
		if valtag != "" {
			fs.vm[flagtag] = valtag
		}
	}
}

func (fs *FlagSetter) SetFlag(f ParsedFlag) error {

	fs.gf.SpecifiedFlags[f.Name] = true

	if prop, ok := fs.m[f.Name]; ok {
		sprop := fs.vm[f.Name]

		setobj := fs.v
		if fs.v.Kind() == reflect.Pointer {
			setobj = setobj.Elem()
		}

		rfld := setobj.FieldByName(prop)
		if rfld.CanSet() {
			rfld.SetBool(f.Switch)
		} else {
			log.Println("can't set field", prop, f.Name)
		}

		if sprop != "" && f.Value != "" {
			vfld := setobj.FieldByName(sprop)
			if vfld.CanSet() {
				vfld.SetString(f.Value)
			} else {
				log.Println("can't set field", sprop, f.Name)
			}
		}
	} else {
		fs.gf.Custom[f.Name] = CustomFlag{
			Name:  f.Name,
			Flag:  f.Switch,
			Value: f.Value,
		}
	}

	return nil
}

func (fs *FlagSetter) GetFlags() GenFlags {
	return *fs.gf
}
