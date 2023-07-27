package proxy

import (
	"gateway/db"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

var (
	_services *services
)

// API
var (
	// 添加服务
	AddService = _services.add
	// 删除服务
	DelService = _services.del
)

func init() {
	_services = new(services)
	_services.init()
}

// service 表示服务组的缓存
type service struct {
	sync.RWMutex
	// key
	k string
	// 列表
	s map[string]*server
}

func (s *service) add(m *db.Server) {
	ser := &server{
		k:       m.ID,
		BaseURL: m.BaseURL,
		Auth:    *m.Authorization == db.True,
		Limite:  time.Second / time.Duration(*m.Limite),
	}
	_u, _ := url.Parse(ser.BaseURL)
	ser.ReverseProxy = httputil.NewSingleHostReverseProxy(_u)
	// 上锁
	s.Lock()
	defer s.Unlock()
	// 添加
	s.s[ser.k] = ser
}

func (s *service) del(key string) {
	// 上锁
	s.Lock()
	defer s.Unlock()
	// 删除
	delete(s.s, key)
}

// getMinLoadServer 获取最小负载服务
func (s *service) getMinLoadServer() *server {
	// 上锁
	s.RLock()
	defer s.RUnlock()
	// 查询
	var min *server
	for _, s := range s.s {
		if min == nil {
			min = s
			continue
		}
		if min.load < s.load {
			min = s
		}
	}
	return min
}

// services 用于管理 service
type services struct {
	// 锁
	sync.RWMutex
	// 列表
	s map[string]*service
}

// init 初始化
func (ss *services) init() {
	ss.s = make(map[string]*service)
}

// add 添加
func (ss *services) add(key string) *service {
	// 上锁
	ss.Lock()
	defer ss.Unlock()
	// 添加列表
	s := ss.s[key]
	if s == nil {
		s = new(service)
		s.k = key
		s.s = make(map[string]*server)
		ss.s[key] = s
	}
	//
	return s
}

// del 删除
func (ss *services) del(key string) {
	// 上锁
	ss.Lock()
	defer ss.Unlock()
	// 删除
	delete(ss.s, key)
}

// get 获取
func (ss *services) get(key string) *service {
	// 上锁
	ss.RLock()
	defer ss.RUnlock()
	// 查询
	return ss.s[key]
}
