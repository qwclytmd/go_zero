package roles

import (
	"net/http"
	"next.com/next/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"next.com/next/backend/internal/logic/system/roles"
	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"
)

func RoleUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			response.APIResponse(w, 0, nil, err)
			return
		}

		l := roles.NewRoleUpdateLogic(r.Context(), svcCtx)
		resp, err := l.RoleUpdate(&req)
		if err != nil {
			response.APIResponse(w, 0, resp, err)
			return
		}
		response.APIResponse(w, 1, resp, err)
	}
}
