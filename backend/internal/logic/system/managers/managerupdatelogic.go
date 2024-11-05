package managers

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"
	"next.com/next/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManagerUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManagerUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManagerUpdateLogic {
	return &ManagerUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManagerUpdateLogic) ManagerUpdate(req *types.ManagerUpdateReq) (resp *types.NilMessage, err error) {

	data := models.Managers{
		RoleId: req.RoleId,
		Status: req.Status,
	}

	if req.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		data.Password = string(password)
	}

	db := l.svcCtx.BackendDB.Model(&models.Managers{})
	if err = db.Model(&models.Managers{Id: req.Id}).Updates(&data).Error; err != nil {
		return nil, err
	}

	return resp, nil
}
