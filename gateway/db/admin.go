package db

import (
	"github.com/qq51529210/util"
	"gorm.io/gorm"
)

var (
	// AdminDA 数据访问
	AdminDA *util.GORMDB[string, *Admin]
)

// Admin 表示管理员
type Admin struct {
	// 主键
	ID string `gorm:"type:varchar(40);primayKey"`
	// 账号
	Account string `gorm:"type:varchar(40);not null;uniqueIndex"`
	// 密码，SHA1 格式
	Password string `gorm:"type:varchar(40);not null"`
	// 是否启用，0/1
	Enable *int8 `gorm:"not null;default:0"`
	util.GORMTime
}

// AdminQuery 是 Admin 查询参数
type AdminQuery struct {
	// 账号，模糊
	Account *string `form:"baseURL" binding:"omitempty,max=40" gq:"like"`
	// 是否启用，精确
	Enable *int8 `form:"enable" binding:"omitempty,oneof=0 1" gq:"eq"`
}

// Init 实现接口
func (m *AdminQuery) Init(db *gorm.DB) *gorm.DB {
	return util.GORMInitQuery(db, m)
}
