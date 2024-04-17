package kqserver

import (
	"context"
	"encoding/json"
	"errors"

	"smartpower/pkg/kqueue"
	"smartpower/pkg/protocol"
	"smartpower/receiver/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
)

type KqServer struct {
	svcCtx *svc.ServiceContext
	queue  queue.MessageQueue

}
 
func MustNewKqServer(svcCtx *svc.ServiceContext) *KqServer {
	q := kqueue.MustNewQueue(svcCtx.KqConf, kqueue.WithHandle(NewConsumeHandler(svcCtx)))
	return &KqServer{
		svcCtx: svcCtx,
		queue:  q,
	}
}

func (s *KqServer) Start() {
	s.queue.Start()
}

func (s *KqServer) Stop() {
	s.queue.Stop()
}

func NewConsumeHandler(svcCtx *svc.ServiceContext) (kqueue.ConsumeHandle) {
	return func(k string, value []byte) error {
		var msg protocol.KqMessage
		if err := json.Unmarshal(value, &msg); err != nil {
			return err
		}

		logx.Debugf("%s ===> %+v", k, value)
		return KqMessageHandler(svcCtx, &msg)
	}
}
var ErrUnsupportCode = errors.New("unsupport code")

func KqMessageHandler(svcCtx *svc.ServiceContext, msg *protocol.KqMessage) error {
	if msg.Bid == "" {
		logx.Error("BID是无效的")
		return nil
	}
	ctx := context.Background()
	handler, ok := svcCtx.Handlers[msg.Code]
	if !ok {
		return ErrUnsupportCode
	}
	return handler(ctx, svcCtx)(msg.Bid, msg.Data)
}
