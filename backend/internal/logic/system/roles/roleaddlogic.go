package roles

import (
	"context"
	"errors"
	"next.com/next/models"
	"time"

	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req *types.RoleAddReq) (resp *types.NilMessage, err error) {
	var isset models.Roles
	if err = l.svcCtx.BackendDB.Model(&models.Roles{}).Where("name = ?", req.Name).Select("name").Limit(1).Find(&isset).Error; err != nil {
		return nil, err
	}
	if isset.Name != "" {
		return nil, errors.New("角色名已存在")
	}

	data := models.Roles{
		Name:        req.Name,
		Permission:  req.Permission,
		Status:      req.Status,
		CreatedTime: time.Now(),
	}

	if err = l.svcCtx.BackendDB.Model(&models.Roles{}).Create(&data).Error; err != nil {
		return nil, err
	}

	return resp, nil
}
