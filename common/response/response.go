package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int `json:"code"`
	Data any `json:"data"`
	Msg  any `json:"msg"`
}

func APIResponse(w http.ResponseWriter, code int, res any, err error) {
	body := Body{Code: code, Data: res}
	if err != nil {
		body.Msg = err.Error()
	}
	httpx.OkJson(w, body)
}
