package managers

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"next.com/next/models"
	"time"

	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManagerAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManagerAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManagerAddLogic {
	return &ManagerAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManagerAddLogic) ManagerAdd(req *types.ManagerAddReq) (resp *types.NilMessage, err error) {

	var issetManager models.Managers
	if err = l.svcCtx.BackendDB.Model(&models.Managers{}).Where("username = ?", req.Username).Select("username").Limit(1).Find(&issetManager).Error; err != nil {
		return nil, err
	}

	if issetManager.Username != "" {
		return nil, errors.New("用户名已存在")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	data := models.Managers{
		Username:     req.Username,
		Password:     string(password),
		RoleId:       req.RoleId,
		Status:       req.Status,
		RegisterTime: time.Now(),
	}

	if err = l.svcCtx.BackendDB.Model(&models.Managers{}).Create(&data).Error; err != nil {
		return nil, err
	}

	return resp, nil
}
