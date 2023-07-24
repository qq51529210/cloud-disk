package db

import "github.com/qq51529210/util"

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
