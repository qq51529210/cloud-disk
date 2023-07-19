package db

import (
	"errors"

	"gorm.io/gorm"
)

// User 表示用户
type User struct {
	ID string `gorm:"type:varchar(40);primayKey"`
	// 账号
	Account string `gorm:"type:varchar(40);uniqueIndex;not null"`
	// 密码，SHA1 格式
	Password *string `gorm:"type:varchar(40);not null"`
	// 是否启用，0/1
	Enable *int8 `gorm:"not null;default:0"`
}

// GetUser 查询单个
func GetUser(id string) (*User, error) {
	m := new(User)
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

// GetUserByAccount 查询单个
func GetUserByAccount(account string) (*User, error) {
	m := new(User)
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

// AddUser 添加单个
func AddUser(m *User) (int64, error) {
	db := _db.Create(m)
	return db.RowsAffected, db.Error
}

// UpdateUser 修改单个
func UpdateUser(m *User) (int64, error) {
	db := _db.
		Where("`ID` = ?", m.ID).
		Updates(m)
	return db.RowsAffected, db.Error
}

// DeleteUser 删除单个
func DeleteUser(id string) (int64, error) {
	db := _db.
		Delete(&Client{
			ID: id,
		})
	return db.RowsAffected, db.Error
}
