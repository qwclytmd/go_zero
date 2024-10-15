package models

import "time"

type Managers struct {
	Id           int64     `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:id" json:"id"`
	Username     string    `gorm:"type:varchar(256);comment:密码" json:"username"`
	Password     string    `gorm:"type:varchar(256);comment:密码" json:"password"`
	RoleId       int64     `gorm:"type:int;comment:角色ID" json:"role_id"`
	Status       string    `gorm:"type:tinyint;comment:状态,1:正常，2：禁用" json:"status"`
	RegisterTime time.Time `gorm:"autoCreateTime;comment:注册时间" json:"register_time"`
}

func (*Managers) TableName() string {
	return "managers"
}
