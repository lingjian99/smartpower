package svc

import (
	"log"
	"smartpower/mqjob/internal/config"
	"smartpower/mqjob/model"

	"github.com/hibiken/asynq"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	//MiniProgram *miniprogram.MiniProgram

	AsynqClient *asynq.Client

	// Models
	AppDeviceTokenModel model.AppDeviceTokenModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := xorm.NewEngine("postgres", c.Pg.Datasource)
	if err != nil {
		log.Fatalf("xorm db err:%+v", err)
	}

	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		AsynqClient: newAsynqClient(c),

		AppDeviceTokenModel: model.NewAppDeviceTokenModel(db),
	}
}
