package managers

import (
	"net/http"
	"next.com/next/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"next.com/next/backend/internal/logic/system/managers"
	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"
)

func ManagerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ManagerListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.APIResponse(w, 0, nil, err)
			return
		}

		l := managers.NewManagerListLogic(r.Context(), svcCtx)
		resp, err := l.ManagerList(&req)

		if err != nil {
			response.APIResponse(w, 0, resp, err)
			return
		}
		response.APIResponse(w, 1, resp, err)
	}
}
