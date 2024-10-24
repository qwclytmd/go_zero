package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"next.com/next/package/jwt"

	"next.com/next/backend/internal/config"
	"next.com/next/backend/internal/handler"
	"next.com/next/backend/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/backend-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		token := r.Header.Get("Authorization")
		claims, err := jwt.ParseToken(token, c.Auth.AccessSecret)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("权限验证失败"))
			return
		}

		context.WithValue(r.Context(), "claims", claims)
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
