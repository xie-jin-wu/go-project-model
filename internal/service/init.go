package service

import (
	"go-project-model/config"
	"go-project-model/internal/biz"
	"go-project-model/pkg/logs"
)

type Service struct {
	biz    *biz.Business
	logger logs.Logger
	config *config.Config
}

func NewService(logger logs.Logger, cfg *config.Config, business *biz.Business) (*Service, error) {
	var s = new(Service)
	s.logger = logger
	s.config = cfg
	s.biz = business
	return s, nil
}

func (s *Service) Run() {
}
