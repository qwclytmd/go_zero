package system

import (
	"context"

	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BackendLogic {
	return &BackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BackendLogic) Backend(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
