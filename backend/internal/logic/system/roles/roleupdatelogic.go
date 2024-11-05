package roles

import (
	"context"
	"errors"
	"next.com/next/models"

	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUpdateLogic) RoleUpdate(req *types.RoleUpdateReq) (resp *types.NilMessage, err error) {
	var isset models.Roles
	if err = l.svcCtx.BackendDB.Model(&models.Roles{}).Where("name = ?", req.Name).Select("name").Limit(1).Find(&isset).Error; err != nil {
		return nil, err
	}
	if isset.Name != "" {
		return nil, errors.New("角色名已存在")
	}

	data := models.Roles{
		Name:       req.Name,
		Permission: req.Permission,
		Status:     req.Status,
	}

	db := l.svcCtx.BackendDB.Model(&models.Roles{})
	if err = db.Model(&models.Roles{Id: req.Id}).Updates(&data).Error; err != nil {
		return nil, err
	}

	return resp, nil
}
