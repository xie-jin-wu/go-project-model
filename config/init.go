package config

import (
	"go-project-model/consts"
	"gopkg.in/yaml.v3"
	"os"
)

func Init() (*Config, error) {
	bytes, err := os.ReadFile(consts.ConfigFile)
	if err != nil {
		panic(err)
	}
	var cfg = new(Config)
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}
