package db

import "github.com/qq51529210/util"

// Admin 表示管理员
type Admin struct {
	ID string `gorm:"type:varchar(40);primayKey"`
	// 账号
	Account string `gorm:"type:varchar(40);uniqueIndex;not null"`
	// 密码，SHA1 格式
	Password *string `gorm:"type:varchar(40);not null"`
	// 是否启用，0/1
	Enable *int8 `gorm:"not null;default:0"`
	util.GORMTime
}
