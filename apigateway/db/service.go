package db

import (
	"github.com/qq51529210/util"
	"gorm.io/gorm"
)

var (
	// ServiceDA 数据访问
	ServiceDA *util.GORMDB[string, *Service]
)

// Service 表示服务，有多个服务器
type Service struct {
	// 主键
	ID string `gorm:"type:varchar(40);primayKey"`
	// 代理路径，/order 这样的
	Path string `gorm:"type:varchar(40);not null;uniqueIndex"`
	// 名称，好记
	Name *string `gorm:"type:varchar(40);not null;uniqueIndex"`
	// 是否启用，0/1
	Enable *int8 `gorm:"not null;default:0"`
	util.GORMTime
}

// ServiceQuery 是 Service 查询参数
type ServiceQuery struct {
	// 代理路径，精确
	Path *string `form:"path" binding:"omitempty,path,max=40" gq:"eq"`
	// 名称，模糊
	Name *string `form:"name" binding:"omitempty,max=40" gq:"like"`
	// 是否启用，精确
	Enable *int8 `form:"enable" binding:"omitempty,oneof=0 1" gq:"eq"`
}

// Init 实现接口
func (m *ServiceQuery) Init(db *gorm.DB) *gorm.DB {
	return util.GORMInitQuery(db, m)
}
