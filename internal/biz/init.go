package biz

import (
	"go-project-model/config"
	"go-project-model/pkg/http_log"
	"go-project-model/pkg/logs"
)

type Business struct {
	logger  logs.Logger
	data    Repo
	config  *config.Config
	httpLog http_log.HttpLog
}

func NewBusiness(logger logs.Logger, cfg *config.Config, data Repo) (*Business, error) {
	var b = new(Business)
	b.logger = logger
	b.config = cfg
	b.data = data
	var opt = http_log.LogOutputToTerminal()
	if cfg.HttpLogToFile {
		opt = http_log.LogOutputToFile(cfg.HttpLogFileDir)
	}
	httpLog, err := http_log.InitHttpLogger(opt)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	b.httpLog = httpLog
	return b, nil
}

type Repo interface {
}
