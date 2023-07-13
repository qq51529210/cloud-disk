package db

// App 表示第三方应用
type App struct {
	ID string `gorm:"type:varchar(40);primayKey"`
	// 密码，SHA1 格式
	Secret *string `gorm:"type:varchar(40);not null"`
	// 名称
	Name *string `gorm:"type:varchar(64);not null"`
	// 描述
	Description *string `gorm:"type:varchar(255);not null"`
	// 是否启用，0/1
	Enable *int8 `gorm:"not null;default:0"`
	// 重定向 url 列表，';' 隔开
	URL *string `gorm:"type:text;not null;"`
	// User.ID
	UserID string `json:"-" gorm:""`
	User   *User  `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// GetApp 查询单个
func GetApp(id string) (*App, error) {
	return nil, nil
}

// AddApp 添加单个
func AddApp(model *App) (int64, error) {
	return 0, nil
}

// UpdateApp 修改单个
func UpdateApp(model *App, userID string) (int64, error) {
	return 0, nil
}

// DeleteApp 删除单个
func DeleteApp(id, userID string) (int64, error) {
	return 0, nil
}
