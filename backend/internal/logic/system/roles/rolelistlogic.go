package roles

import (
	"context"
	"next.com/next/models"

	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	db := l.svcCtx.BackendDB.Model(&models.Roles{})
	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
	}

	var record []models.Roles
	if err = db.Select("a.*, b.name, b.permission").Find(&record).Error; err != nil {
		return nil, err
	}

	resp = &types.RoleListResp{
		Total: 0,
		Data:  record,
	}

	return
}
