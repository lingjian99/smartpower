package logic

import (
	"context"

	"smartpower/diagram/internal/svc"
	"smartpower/diagram/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StructDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStructDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StructDataLogic {
	return &StructDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StructDataLogic) StructData(req *types.IdReq) (resp *types.JsonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
