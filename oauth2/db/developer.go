package db

import (
	"errors"

	"gorm.io/gorm"
)

// Developer 表示第三方应用的开发者
type Developer struct {
	ID string `gorm:"type:varchar(40);primayKey"`
	// 账号
	Account string `gorm:"type:varchar(40);uniqueIndex;not null"`
	// 密码，SHA1 格式
	Password *string `gorm:"type:varchar(40);not null"`
	// 是否启用，0/1
	Enable *int8 `gorm:"not null;default:0"`
	// Developer.ID 表示这个应用属于哪一个开发者
	DeveloperID string     `json:"-" gorm:""`
	Developer   *Developer `json:"-" gorm:"foreignKey:DeveloperID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// GetDeveloper 查询单个
func GetDeveloper(id string) (*Developer, error) {
	m := new(Developer)
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

// GetDeveloperByAccount 查询单个
func GetDeveloperByAccount(account string) (*Developer, error) {
	m := new(Developer)
	err := _db.
		Where("`Account` = ?", account).
		First(m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}

// AddDeveloper 添加单个
func AddDeveloper(m *Developer) (int64, error) {
	db := _db.Create(m)
	return db.RowsAffected, db.Error
}

// UpdateDeveloper 修改单个
func UpdateDeveloper(m *Developer) (int64, error) {
	db := _db.
		Where("`ID` = ?", m.ID).
		Updates(m)
	return db.RowsAffected, db.Error
}

// DeleteDeveloper 删除单个
func DeleteDeveloper(id string) (int64, error) {
	db := _db.
		Delete(&Client{
			ID: id,
		})
	return db.RowsAffected, db.Error
}
