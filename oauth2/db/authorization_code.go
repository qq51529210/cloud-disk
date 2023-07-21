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
	Code string
	// 范围
	Scope string
	// 客户端 ID
	ClientID string
	// 客户端密钥
	ClientSecret string
	// 重定向
	RedirectURI string
}

// NewAuthorizationCodeWithContext 创建授权码
func NewAuthorizationCodeWithContext(ctx context.Context, code *AuthorizationCode) error {
	code.Code = uuid.LowerV1WithoutHyphen()
	return Put(ctx, AuthorizationCodePrefix+code.Code, code, cfg.Cfg.OAuth2.AuthorizationCodeExpires)
}

// NewAuthorizationCode 创建授权码
func NewAuthorizationCode(code *AuthorizationCode) error {
	// 超时
	ctx, cancel := newRedisTimeout()
	defer cancel()
	//
	return NewAuthorizationCodeWithContext(ctx, code)
}

// GetAuthorizationCodeWithContext 查询授权码
func GetAuthorizationCodeWithContext(ctx context.Context, code string) (*AuthorizationCode, error) {
	return GetWithContext[AuthorizationCode](ctx, AuthorizationCodePrefix+code)
}

// GetAuthorizationCode 查询授权码
func GetAuthorizationCode(code string) (*AuthorizationCode, error) {
	return Get[AuthorizationCode](AuthorizationCodePrefix + code)
}
