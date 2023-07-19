package db

import (
	"errors"

	"gorm.io/gorm"
)

// App 表示第三方应用
type App struct {
	ID string `gorm:"type:varchar(40);primayKey"`
	// 密码，SHA1 格式
	Secret *string `gorm:"type:varchar(40);not null"`
	// 名称
	Name *string `gorm:"type:varchar(64);not null"`
	// 描述
	Description *string `gorm:"type:varchar(255);"`
	// 是否启用，0/1
	Enable *int8 `gorm:"not null;default:0"`
	// 重定向 url 列表，';' 隔开
	URL *string `gorm:"type:text;"`
	// User.ID
	UserID string `json:"-" gorm:""`
	User   *User  `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// GetApp 查询单个
func GetApp(id string) (*App, error) {
	m := new(App)
	err := _db.
		Where("`ID` = ?", id).
		First(m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}

// AddApp 添加单个
func AddApp(m *App) (int64, error) {
	db := _db.Create(m)
	return db.RowsAffected, db.Error
}

// UpdateApp 修改单个
func UpdateApp(m *App, userID string) (int64, error) {
	db := _db.
		Where("`ID` = ?", m.ID).
		Where("`UserID` = ?", userID).
		Updates(m)
	return db.RowsAffected, db.Error
}

// DeleteApp 删除单个
func DeleteApp(id, userID string) (int64, error) {
	db := _db.
		Where("`UserID` = ?", userID).
		Delete(&App{
			ID: id,
		})
	return db.RowsAffected, db.Error
}
