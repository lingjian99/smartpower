package devicelogic

import (
	"context"

	"smartpower/device/internal/svc"
	"smartpower/device/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnaDeviceConnectedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnaDeviceConnectedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnaDeviceConnectedLogic {
	return &AnaDeviceConnectedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnaDeviceConnectedLogic) AnaDeviceConnected(in *pb.Empty) (*pb.DeviceOnline, error) {
	// todo: add your logic here and delete this line

	return &pb.DeviceOnline{}, nil
}
