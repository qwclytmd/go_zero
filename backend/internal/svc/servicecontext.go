package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"next.com/next/backend/internal/config"
	"next.com/next/backend/internal/middleware"
	"next.com/next/models"
	"next.com/next/package/mysql"
)

type ServiceContext struct {
	Config         config.Config
	CorsMiddleware rest.Middleware
	BackendDB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	backendTables := map[string]any{
		"后台用户": models.Managers{},
		"后台菜单": models.Menus{},
		"后台角色": models.Roles{},
	}

	return &ServiceContext{
		Config:         c,
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
		BackendDB:      mysql.Connection(c.MySQL.Backend, backendTables),
	}
}
