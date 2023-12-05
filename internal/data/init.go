package data

import (
	"go-project-model/config"
	"go-project-model/internal/biz"
	"go-project-model/pkg/logs"
)

type data struct {
	logger logs.Logger
	config *config.Config
}

func NewData(cfg *config.Config, log logs.Logger) (biz.Repo, error) {
	var v = new(data)
	v.config = cfg
	v.logger = log
	return v, nil
}
