package proxy

import (
	"apigateway/cfg"
	"net/http"
)

// Serve 开始服务
func Serve() error {
	// 监听
	return http.ListenAndServe(cfg.Cfg.ProxyAddr, http.HandlerFunc(handle))
}

// handle 处理所有请求
func handle(w http.ResponseWriter, r *http.Request) {
	// 找出第一层
	path := r.URL.Path
	for i := 1; i < len(path); i++ {
		if i == '/' {
			path = path[1:i]
			break
		}
	}
	// 获取服务
}
