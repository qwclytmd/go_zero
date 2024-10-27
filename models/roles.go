package models

import (
	"gorm.io/gorm"
	"time"
)

type Roles struct {
	Id          int64          `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:id" json:"id"`
	Name        string         `gorm:"not null;type:varchar(256);comment:角色名" json:"name"`
	Permission  string         `gorm:"not null;type:varchar(256);comment:权限" json:"permission"`
	Status      int            `gorm:"not null;type:tinyint;comment:状态,1:正常，2：禁用" json:"status"`
	CreatedTime time.Time      `gorm:"not null;autoCreateTime;comment:创建时间" json:"created_time"`
	UpdatedTime time.Time      `gorm:"not null;autoUpdateTime;comment:更新时间" json:"updated_time"`
	DeletedTime gorm.DeletedAt `gorm:"comment:删除时间" json:"deleted_time"`
}

func (*Roles) TableName() string {
	return "roles"
}
