package managers

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"
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

type ManagerListResponse struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	RoleId       int64  `json:"role_id"`
	Status       int    `json:"status"`
	RegisterTime string `json:"register_time"`
	Name         string `json:"name"`
}

func (l *ManagerListLogic) ManagerList(req *types.ManagerListReq) (resp *types.ManagerListResp, err error) {
	db := l.svcCtx.BackendDB.Table("managers as a").Joins("inner join roles as b on a.role_id = b.id")
	if req.Username != "" {
		db = db.Where("username = ?", req.Username)
	}

	if req.RoleName != "" {
		db = db.Where("role_name = ?", req.RoleName)
	}

	var record []ManagerListResponse
	if err = db.Select("a.*, b.name, b.permission").Find(&record).Error; err != nil {
		return nil, err
	}

	resp = &types.ManagerListResp{
		Total: 0,
		Data:  record,
	}

	return
}
