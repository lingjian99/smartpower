package spserver

import (
	"smartpower/pkg/protocol"

	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
)
type SpServer struct {

	ProtoAddr string
	Multicore bool

	EventHandler *EventHandler
}

func (l *SpServer) Start() {
	
	err := gnet.Run(l.EventHandler, 
		l.ProtoAddr,
		gnet.WithMulticore(l.Multicore),
		gnet.WithTicker(true),
		//gnet.WithReuseAddr(true),
		gnet.WithReusePort(true),
	)
	if err != nil {
		logx.Error(err)
		panic(err)
	}
}

func (l *SpServer) Stop() {
	l.EventHandler.Stop()
}

func (l *SpServer) Send(bid string, in *protocol.Input) error {
	return nil

}