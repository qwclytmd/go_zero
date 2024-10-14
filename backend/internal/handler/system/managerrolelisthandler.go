package system

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"next.com/next/backend/internal/logic/system"
	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"
)

func ManagerRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ManagerRoleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := system.NewManagerRoleListLogic(r.Context(), svcCtx)
		resp, err := l.ManagerRoleList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
