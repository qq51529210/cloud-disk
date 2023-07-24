package db

import "github.com/qq51529210/util"

// Server 表示一个具体的服务
type Server struct {
	// 主键
	ID string `gorm:"type:varchar(40);primayKey"`
	// 所属的服务
	ServiceID string `json:"serviceID" gorm:"type:varchar(40);not null;uniqueIndex:ServerUnique"`
	// 所属的服务
	Service *Service `json:"-" gorm:"foreignKey:ServiceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// 基本路径，http(https)://hostname:port/
	BaseURL string `gorm:"type:varchar(128);not null;uniqueIndex:ServerUnique"`
	// 名称，好记
	Name *string `gorm:"type:varchar(40);not null;uniqueIndex"`
	// 是否启用，0/1
	Enable *int8 `gorm:"not null;default:0"`
	// 在线状态，0/1
	Online *int8 `gorm:"not null;default:0"`
	util.GORMTime
}
