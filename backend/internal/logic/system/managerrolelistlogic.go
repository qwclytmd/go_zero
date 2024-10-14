package system

import (
	"context"

	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManagerRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManagerRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManagerRoleListLogic {
	return &ManagerRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManagerRoleListLogic) ManagerRoleList(req *types.ManagerRoleListReq) (resp *types.ManagerRoleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
