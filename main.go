package main

import (
	"github.com/xie-jin-wu/logs"
	"go-project-model/config"
	"go-project-model/internal/biz"
	"go-project-model/internal/data"
	"go-project-model/internal/service"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func initApp(cfg *config.Config, logger logs.Logger) (*service.Service, error) {
	repo, err := data.NewData(cfg, logger)
	if err != nil {
		return nil, err
	}
	business, err := biz.NewBusiness(logger, cfg, repo)
	if err != nil {
		return nil, err
	}
	return service.NewService(logger, cfg, business)
}

func main() {
	cfg, err := config.Init()
	if err != nil {
		return
	}
	var opt = logs.LogOutputToTerminal()
	if cfg.ProgramLogToFile {
		opt = logs.LogOutputToFile(cfg.ProgramLogFileDir)
	}
	logger, err := logs.NewLogger(cfg.ProgramLogLevel, opt)
	if err != nil {
		panic(err)
	}
	app, err := initApp(cfg, logger)
	if err != nil {
		logger.Error(err)
		return
	}
	app.Run()
}
