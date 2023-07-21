package db

import (
	"errors"

	"github.com/qq51529210/util"
	"gorm.io/gorm"
)

// Client 表示第三方应用
type Client struct {
	ID string `gorm:"type:varchar(40);primayKey"`
	// Developer.ID 表示这个应用属于哪一个开发者
	DeveloperID string     `json:"-" gorm:""`
	Developer   *Developer `json:"-" gorm:"foreignKey:DeveloperID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// 密码，SHA1 格式
	Secret *string `gorm:"type:varchar(40);not null"`
	// 重定向，好像没有什么卵用，除了验证是否与请求中的一致
	RedirectURI *string `gorm:"type:varchar(255)"`
	// 名称
	Name *string `gorm:"type:varchar(64);not null"`
	// 图片
	Image *string `gorm:"type:text"`
	// 描述
	Description *string `gorm:"type:varchar(255);"`
	// 是否启用，0/1
	Enable *int8 `gorm:"not null;default:0"`
	// 授权，value1:name1 value2:name2 ... 多个用空格分开
	Scope *string `gorm:"type:varchar(255);"`
	util.GORMTime
}

// GetClient 查询单个
func GetClient(id string) (*Client, error) {
	// todo 做缓存
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
