package db

import (
	"errors"

	"gorm.io/gorm"
)

// Client 表示第三方应用
type Client struct {
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
	// Developer.ID 表示这个应用属于哪一个开发者
	DeveloperID string     `json:"-" gorm:""`
	Developer   *Developer `json:"-" gorm:"foreignKey:DeveloperID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// GetClient 查询单个
func GetClient(id string) (*Client, error) {
	m := new(Client)
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

// AddClient 添加单个
func AddClient(m *Client) (int64, error) {
	db := _db.Create(m)
	return db.RowsAffected, db.Error
}

// UpdateClient 修改单个
func UpdateClient(m *Client, DeveloperID string) (int64, error) {
	db := _db.
		Where("`ID` = ?", m.ID).
		Where("`DeveloperID` = ?", DeveloperID).
		Updates(m)
	return db.RowsAffected, db.Error
}

// DeleteClient 删除单个
func DeleteClient(id, DeveloperID string) (int64, error) {
	db := _db.
		Where("`DeveloperID` = ?", DeveloperID).
		Delete(&Client{
			ID: id,
		})
	return db.RowsAffected, db.Error
}
