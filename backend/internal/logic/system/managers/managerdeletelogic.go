package managers

import (
	"context"
	"next.com/next/models"

	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManagerDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManagerDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManagerDeleteLogic {
	return &ManagerDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManagerDeleteLogic) ManagerDelete(req *types.ManagerDeleteReq) (resp *types.NilMessage, err error) {
	if err = l.svcCtx.BackendDB.Delete(&models.Managers{}, req.Id).Error; err != nil {
		return nil, err
	}

	return
}
