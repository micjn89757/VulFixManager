package svc

import (
	"github.com/micjn89757/VulFixMiner/server/basicService/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
