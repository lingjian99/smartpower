package logic

import (
	"context"

	"smartpower/diagram/internal/svc"
	"smartpower/diagram/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CabDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCabDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CabDataLogic {
	return &CabDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CabDataLogic) CabData(req *types.IdReq) (resp *types.JsonResp, err error) {

	return
}
