package svc

import (
	"smartpower/diagram/internal/config"
	"smartpower/diagram/internal/middleware"
	"smartpower/pkg/def/defcache"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config  config.Config
	AppAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	cc := cache.New(c.Cache, syncx.NewSingleFlight(), cache.NewStat("apic"), defcache.ErrCacheErrNotFound)

	return &ServiceContext{
		Config:  c,
		AppAuth: middleware.NewAppAuthMiddleware(cc).Handle,
	}
}
