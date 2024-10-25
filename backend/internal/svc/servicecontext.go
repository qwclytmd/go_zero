package svc

import (
	"gorm.io/gorm"
	"next.com/next/backend/internal/config"
	"next.com/next/models"
	"next.com/next/package/mysql"
)

type ServiceContext struct {
	Config    config.Config
	BackendDB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	backendTables := map[string]any{
		"后台管理员": models.Managers{},
	}

	return &ServiceContext{
		Config:    c,
		BackendDB: mysql.Connection(c.MySQL.Backend, backendTables),
	}
}
