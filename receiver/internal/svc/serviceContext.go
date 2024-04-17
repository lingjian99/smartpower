package svc

import (
	"context"

	"smartpower/pkg/kqueue"
	"smartpower/receiver/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	KqConf   kqueue.Config
	Handlers map[uint16]func(context.Context, *ServiceContext) func(string, []byte) error
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
	}
}
