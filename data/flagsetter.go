package data

import (
	"fmt"
	"log"
	"reflect"
	"strings"
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

func (fs *FlagSetter) SetFlag(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("flag is 0-length")
	}

	flg := s[0] == '+'

	rest := strings.Fields(s[1:])
	flgname := rest[0]

	if flgname == "" {
		log.Println("got blank flag name", s)
	}

	fs.gf.SpecifiedFlags[flgname] = flg

	var val string
	if len(rest) > 1 {
		val = strings.Join(rest[1:], " ")
	}

	if prop, ok := fs.m[flgname]; ok {
		sprop := fs.vm[flgname]

		setobj := fs.v
		if fs.v.Kind() == reflect.Pointer {
			setobj = setobj.Elem()
		}

		rfld := setobj.FieldByName(prop)
		if rfld.CanSet() {
			rfld.SetBool(flg)
		} else {
			log.Println("can't set field", prop, s)
		}

		if sprop != "" && val != "" {
			vfld := setobj.FieldByName(sprop)
			if vfld.CanSet() {
				vfld.SetString(val)
			} else {
				log.Println("can't set field", sprop, s)
			}
		}
	} else {
		fs.gf.Custom[flgname] = CustomFlag{
			Name:  flgname,
			Flag:  flg,
			Value: val,
		}
	}

	return nil
}

func (fs *FlagSetter) GetFlags() GenFlags {
	return *fs.gf
}
