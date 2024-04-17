package config

import (
	"smartpower/device/spnet/types"
	"smartpower/pkg/kqueue"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	HwServer types.NetConfig
	LogX     logx.LogConf
	KqConf   kqueue.Config
}
