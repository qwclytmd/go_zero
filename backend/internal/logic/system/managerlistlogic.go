package system

import (
	"context"

	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManagerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManagerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManagerListLogic {
	return &ManagerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManagerListLogic) ManagerList(req *types.ManagerListReq) (resp *types.ManagerListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
