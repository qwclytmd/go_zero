package svc

import (
	"gorm.io/gorm"
	"next.com/next/backend/internal/config"
	"next.com/next/common/mysql"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		DB:     GetMySQLConn(c),
	}
}

// 数据库连接
func GetMySQLConn(c config.Config) *gorm.DB {
	db, err := mysql.Connection(c.MySQL.DSN)
	if err != nil {
		panic(err)
	}

	return db
}
