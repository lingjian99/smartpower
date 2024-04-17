package svc

import (
	//"smartpower/device/connect"
	"smartpower/device/internal/config"
	"smartpower/device/spnet"
)

type ServiceContext struct {
	Config       config.Config
	//SocketServer *connect.HwServer
	SpServer *spnet.IServer
	}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:       c,
		SpServer: spnet.NewServer(c.KqConf.Brokers, c.KqConf.Topic, &c.HwServer),
	}
}
