package devicelogic

import (
	"context"

	"smartpower/device/internal/svc"
	"smartpower/device/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendToDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendToDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendToDeviceLogic {
	return &SendToDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendToDeviceLogic) SendToDevice(in *pb.DeviceData) (*pb.Empty, error) {
	// todo: add your logic here and delete this line

	return &pb.Empty{}, nil
}
