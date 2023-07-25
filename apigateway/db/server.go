package db

import (
	"github.com/qq51529210/util"
	"gorm.io/gorm"
)

var (
	// ServerDA 数据访问
	ServerDA *util.GORMDB[string, *Server]
)

// Server 表示一个具体的服务
type Server struct {
	// 主键
	ID string `gorm:"type:varchar(40);primayKey"`
	// Service.ID
	ServiceID string `json:"serviceID" gorm:"type:varchar(40);not null;uniqueIndex:ServerUnique"`
	// 所属的服务组
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

// ServerQuery 是 Server 查询参数
type ServerQuery struct {
	// 所属的服务组，精确
	ServiceID *string `form:"enable" binding:"omitempty,max=40" gq:"eq"`
	// 基本路径，模糊
	BaseURL *string `form:"name" binding:"omitempty,max=40" gq:"like"`
	// 名称，模糊
	Name *string `form:"baseURL" binding:"omitempty,max=128" gq:"like"`
	// 是否启用，精确
	Enable *int8 `form:"enable" binding:"omitempty,oneof=0 1" gq:"eq"`
	// 在线状态，精确
	Online *int8 `form:"online" binding:"omitempty,oneof=0 1" gq:"eq"`
}

// Init 实现接口
func (m *ServerQuery) Init(db *gorm.DB) *gorm.DB {
	return util.GORMInitQuery(db, m)
}
