package db

// User 表示用户
type User struct {
	ID string `gorm:"type:varchar(40);primayKey"`
	// 账号
	Account string `gorm:"type:varchar(40);uniqueIndex;not null"`
	// 密码，SHA1 格式
	Password *string `gorm:"type:varchar(40);not null"`
}

// GetUser 查询单个
func GetUser(id string) (*User, error) {
	return nil, nil
}

// AddUser 添加单个
func AddUser(model *User) (int64, error) {
	return 0, nil
}

// UpdateUser 修改单个
func UpdateUser(model *User) (int64, error) {
	return 0, nil
}

// DeleteUser 删除单个
func DeleteUser(id string) (int64, error) {
	return 0, nil
}
