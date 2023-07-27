package proxy

import (
	"gateway/cfg"
	"net/http"
	"net/url"
)

// authorize 身份验证
func authorize(w http.ResponseWriter, r *http.Request) bool {
	// 没有配置
	if cfg.Cfg.AuthService.TokenURL == "" {
		return true
	}
	// 获取 token
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	_u, _ := url.Parse(cfg.Cfg.AuthService.TokenURL)
	q := _u.Query()
	q.Set("token", token)
	_u.RawQuery = q.Encode()
	res, err := http.Get(_u.String())
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return false
	}
	defer res.Body.Close()
	// 是否 200 ok
	return res.StatusCode == http.StatusOK
}
