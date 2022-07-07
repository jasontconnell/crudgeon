package configuration

import (
	"github.com/jasontconnell/conf"
)

type Config struct {
	Generations        []Generation `json:"generations"`
	ConcreteCollection string       `json:"concreteCollection"`
	AbstractCollection string       `json:"abstractCollection"`
}

type Generation struct {
	File         string `json:"file"`
	FileType     string `json:"fileType"`
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
