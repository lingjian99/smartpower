package devicelogic

import (
	"context"

	"smartpower/device/internal/svc"
	"smartpower/device/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EMTimeSlotPollingPushLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEMTimeSlotPollingPushLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EMTimeSlotPollingPushLogic {
	return &EMTimeSlotPollingPushLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 服务器下发时段配置数据 载体固定96字节
func (l *EMTimeSlotPollingPushLogic) EMTimeSlotPollingPush(in *pb.EMTimeSlot) (*pb.Empty, error) {
	// todo: add your logic here and delete this line

	return &pb.Empty{}, nil
}
