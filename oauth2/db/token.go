package db

// Token 表示访问令牌，使用 redis 来保存
type Token struct {
	// uuid
	ID string
	// 过期时间
	Expires int64
	// 应用 id
	ClientID string
	// 用户 id
	UserID string
}

// NewToken 创建新的访问令牌
func NewToken(user *User) (*Token, error) {
	return nil, nil
}

// GetToken 获取会话
func GetToken(id string) (*Token, error) {
	return nil, nil
}
