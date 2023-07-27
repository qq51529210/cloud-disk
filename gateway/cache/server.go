package cache

import (
	"errors"
	"gateway/db"
	"net/http/httputil"
	"sync/atomic"
	"time"
)

var (
	// ErrServiceNotFound 服务组不存在
	ErrServiceNotFound = errors.New("service not found")
	// ErrServerNotFound 服务不存在
	ErrServerNotFound = errors.New("server not found")
)

// Server 表示一个服务器
type Server struct {
	// key
	k string
	// 负载
	load int64
	// 地址
	BaseURL string
	// 访问控制
	Limite time.Duration
	// 是否需要身份验证
	Auth bool
	// 代理
	*httputil.ReverseProxy
}

// Minus 减小负载
func (s *Server) Minus() {
	atomic.AddInt64(&s.load, -1)
}

// GetMinLoadServer 获取最小负载的服务
// 内部自动增加负载，完事后需要调用 Minus 恢复
func GetMinLoadServer(service string) *Server {
	ss := _services.get(service)
	if ss == nil {
		return nil
	}
	s := ss.getMinLoadServer()
	if s == nil {
		return nil
	}
	atomic.AddInt64(&s.load, 1)
	return s
}

// AddServer 添加一个服务
func AddServer(ser *db.Server) {
	ss := _services.add(ser.ServiceID)
	ss.add(ser)
}

// DelServer 删除指定服务
func DelServer(service, server string) {
	ss := _services.add(service)
	ss.del(server)
}
