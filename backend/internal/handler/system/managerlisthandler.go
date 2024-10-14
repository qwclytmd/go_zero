package system

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"next.com/next/backend/internal/logic/system"
	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"
)

func ManagerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ManagerListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := system.NewManagerListLogic(r.Context(), svcCtx)
		resp, err := l.ManagerList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
