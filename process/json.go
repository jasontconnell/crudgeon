package process

import (
	"bytes"
	"encoding/json"
	"github.com/jasontconnell/crudgeon/data"
	"io/ioutil"
	"os"
)

func GenerateJson(fields []data.Field, file string) error {
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	enc.SetIndent("", "   ")

	enc.Encode(fields)

	return ioutil.WriteFile(file, buffer.Bytes(), os.ModePerm)
}

func ParseJsonFields(file string) ([]data.Field, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(f)

	var flds []data.Field
	err = dec.Decode(&flds)
	if err != nil {
		return nil, err
	}

	return flds, nil
}
