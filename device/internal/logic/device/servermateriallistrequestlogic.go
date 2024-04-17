package devicelogic

import (
	"context"

	"smartpower/device/internal/svc"
	"smartpower/device/pb"
	"smartpower/pkg/protocol"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServerMaterialListRequestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewServerMaterialListRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServerMaterialListRequestLogic {
	return &ServerMaterialListRequestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}


// 服务器请求材料清单数据
func (l *ServerMaterialListRequestLogic) Server_MaterialList_Request(in *pb.BidReq) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	ss := l.svcCtx.SpServer

	err := (*ss).Send(in.Bid, &protocol.Input{
		Head: protocol.Head{
			BodyLen: 0,
			Code:    0x0EFF,
			Flag:    0x01,
			DataDst: 0x10,
		},
		Data: protocol.RawData16(0xff),
	})

	if err != nil {
		logx.Errorf("Send to %s Code: %x err: %+v", in.Bid, 0x0EFF, err)
	}
	return &pb.Empty{}, nil
}
