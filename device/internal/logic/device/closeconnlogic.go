package devicelogic

import (
	"context"

	"smartpower/device/internal/svc"
	"smartpower/device/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseConnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCloseConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseConnLogic {
	return &CloseConnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CloseConnLogic) CloseConn(in *pb.Conn) (*pb.Empty, error) {
	// todo: add your logic here and delete this line

	return &pb.Empty{}, nil
}
