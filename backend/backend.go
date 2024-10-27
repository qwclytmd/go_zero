package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"next.com/next/backend/internal/config"
	"next.com/next/backend/internal/handler"
	"next.com/next/backend/internal/svc"
	"next.com/next/package/jwt"
	"os"
)

var configFile = flag.String("f", "etc/backend.yaml", "the config file")

func main() {
	flag.Parse()

	//解析配置
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//添加日志
	logx.MustSetup(c.Log)
	logx.AddWriter(logx.NewWriter(os.Stdout))

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

	fmt.Printf("%s Starting Server %s:%d...\n", c.Name, c.Host, c.Port)
	server.Start()
}
