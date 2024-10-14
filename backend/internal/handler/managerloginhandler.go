package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"next.com/next/backend/internal/logic"
	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"
)

func ManagerLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ManagerLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewManagerLoginLogic(r.Context(), svcCtx)
		resp, err := l.ManagerLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
