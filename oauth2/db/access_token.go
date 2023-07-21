package db

// AccessToken 表示访问令牌，使用 redis 来保存
type AccessToken struct {
	// uuid
	ID string
	// 过期时间
	Expires int64
	// 应用 id
	ClientID string
	// 用户 id
	UserID string
	// 范围
	Scope string
}

// NewAccessToken 创建新的访问令牌
func NewAccessToken(user *User) (*AccessToken, error) {
	return nil, nil
}

// GetAccessToken 获取会话
func GetAccessToken(id string) (*AccessToken, error) {
	return nil, nil
}
