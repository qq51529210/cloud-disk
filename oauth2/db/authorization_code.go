package db

import (
	"context"
	"oauth2/cfg"

	"github.com/qq51529210/uuid"
)

// redis key 的前缀
const (
	AuthorizationCodePrefix = "authorization_code:"
)

// AuthorizationCode 表示授权码，使用 redis 来保存
type AuthorizationCode struct {
	// 码
	ID string
	// 范围
	Scope string
	// 重定向
	RedirectURI string
	// 应用
	ClientID string
	// 用户
	UserID string
}

// PutAuthorizationCodeWithContext 创建授权码
func PutAuthorizationCodeWithContext(ctx context.Context, code *AuthorizationCode) error {
	code.ID = uuid.LowerV1WithoutHyphen()
	return PutWithContext(ctx, AuthorizationCodePrefix+code.ID, code, cfg.Cfg.OAuth2.AuthorizationCodeExpires)
}

// PutAuthorizationCode 创建授权码
func PutAuthorizationCode(code *AuthorizationCode) error {
	// 超时
	ctx, cancel := newRedisTimeout()
	defer cancel()
	//
	return PutAuthorizationCodeWithContext(ctx, code)
}

// GetAuthorizationCodeWithContext 查询授权码
func GetAuthorizationCodeWithContext(ctx context.Context, code string) (*AuthorizationCode, error) {
	return GetWithContext[AuthorizationCode](ctx, AuthorizationCodePrefix+code)
}

// GetAuthorizationCode 查询授权码
func GetAuthorizationCode(code string) (*AuthorizationCode, error) {
	return Get[AuthorizationCode](AuthorizationCodePrefix + code)
}
