package logic

import (
	"context"

	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManagerLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManagerLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManagerLoginLogic {
	return &ManagerLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManagerLoginLogic) ManagerLogin(req *types.ManagerLoginReq) (resp *types.ManagerLoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
