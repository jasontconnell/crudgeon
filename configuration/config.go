package configuration

import (
	"github.com/jasontconnell/conf"
)

type Config struct {
	Generations        []Generation `json:"generations"`
	ConcreteCollection string       `json:"concreteCollection"`
	AbstractCollection string       `json:"abstractCollection"`
	GenericRegex       string       `json:"genericRegex"`
}

type Generation struct {
	File         string `json:"file"`
	Alias        string `json:"alias"`
	Extension    string `json:"ext"`
	Database     bool   `json:"db"`
	OutputPrefix string `json:"outputPrefix"`
	OutputSuffix string `json:"outputSuffix"`
	Folder       string `json:"folder"`
	Flags        string `json:"flags"`
	CreateObjDir bool   `json:"objdir"`
}

func LoadConfig(filename string) Config {
	cfg := Config{}
	conf.LoadConfig(filename, &cfg)
	return cfg
}
