package models

type Menus struct {
	Id        int64  `gorm:"column:id;primaryKey;autoIncrement;type:bigint;comment:id" json:"id"`
	Pid       int64  `gorm:"not null;type:bigint;comment:上级节点" json:"pid"`
	Type      int    `gorm:"not null;type:tinyint;comment:类型,1:菜单,2:api" json:"is_menu"`
	Title     string `gorm:"not null;type:varchar(256);comment:标题" json:"title"`
	Path      string `gorm:"not null;default:'';type:varchar(256);comment:路径" json:"path"`
	Component string `gorm:"not null;default:'';type:varchar(256);comment:组件" json:"component"`
	Sort      int    `gorm:"not null;type:int;comment:排序" json:"sort"`
	Status    int    `gorm:"not null;type:tinyint;comment:状态,1:正常，2：禁用" json:"status"`
}

func (*Menus) TableName() string {
	return "menus"
}
