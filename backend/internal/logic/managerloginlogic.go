package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"next.com/next/models"

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

	manager := models.Managers{}
	if err = l.svcCtx.DB.Model(&models.Managers{}).Where("username = ?", req.Username).First(&manager).Error; err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	if bcrypt.CompareHashAndPassword([]byte(manager.Password), []byte(req.Password)) != nil {
		return nil, errors.New("用户名或密码错误")
	}

	return
}
