package devicelogic

import (
	"context"

	"smartpower/device/internal/svc"
	"smartpower/device/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoteControlRequestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoteControlRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoteControlRequestLogic {
	return &RemoteControlRequestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 请求远程控制
func (l *RemoteControlRequestLogic) RemoteControlRequest(in *pb.RemoteControlReq) (*pb.RemoteControlResp, error) {
	// todo: add your logic here and delete this line

	return &pb.RemoteControlResp{}, nil
}
