package conf

import (
	"github.com/jasontconnell/conf"
)

type Config struct {
}

func LoadConfig(filename string) Config {
	cfg := Config{}
	conf.LoadConfig(filename, &cfg)
	return cfg
}
