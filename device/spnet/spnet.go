package spnet

import (
	"smartpower/device/spnet/internal/spserver"
	"smartpower/device/spnet/types"
	"smartpower/pkg/kqueue"
	"smartpower/pkg/protocol"
	"sync"
)


type IServer interface {
	Start()
	Stop()
	Send(bid string, in *protocol.Input) error
} 


var once sync.Once
func NewServer(brokers []string, topic string, conf *types.NetConfig) (*IServer) {
	
	var s IServer
	producer := kqueue.NewProducer(brokers, topic)
	once.Do(func() {
		ss := spserver.SpServer{
			EventHandler: spserver.NewLisServer(conf, producer),
			ProtoAddr:  conf.Network + "://" + conf.ListenOn,
			Multicore:  conf.Multicore,

		}
		
		s = &ss
	})
	return &s
}

