package svc

import (
	"smartpower/mqjob/internal/config"

	"github.com/hibiken/asynq"
)

func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass})
}
