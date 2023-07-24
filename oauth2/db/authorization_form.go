package db

import (
	"context"
	"oauth2/cfg"

	"github.com/qq51529210/uuid"
)

// redis key 的前缀
const (
	AuthorizationFormPrefix = "authorization_form:"
)

// AuthorizationForm 用于保存授权页面的信息
type AuthorizationForm struct {
	ID     string
	Client *Client
	
}

// PutAuthorizationFormWithContext 创建授权码
func PutAuthorizationFormWithContext(ctx context.Context, form *AuthorizationForm) error {
	form.ID = uuid.LowerV1WithoutHyphen()
	return PutWithContext(ctx, AuthorizationFormPrefix+form.ID, form, cfg.Cfg.OAuth2.AuthorizationFormExpires)
}

// PutAuthorizationForm 创建授权码
func PutAuthorizationForm(form *AuthorizationForm) error {
	// 超时
	ctx, cancel := newRedisTimeout()
	defer cancel()
	//
	return PutAuthorizationFormWithContext(ctx, form)
}

// GetAuthorizationFormWithContext 查询授权码
func GetAuthorizationFormWithContext(ctx context.Context, form string) (*AuthorizationForm, error) {
	return GetWithContext[AuthorizationForm](ctx, AuthorizationFormPrefix+form)
}

// GetAuthorizationForm 查询授权码
func GetAuthorizationForm(form string) (*AuthorizationForm, error) {
	return Get[AuthorizationForm](AuthorizationFormPrefix + form)
}

// DelAuthorizationFormWithContext 查询授权码
func DelAuthorizationFormWithContext(ctx context.Context, form string) error {
	return rds.Del(ctx, AuthorizationFormPrefix+form).Err()
}

// DelAuthorizationForm 查询授权码
func DelAuthorizationForm(form string) error {
	return Del(AuthorizationFormPrefix + form)
}
