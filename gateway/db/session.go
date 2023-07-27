package db

import (
	"gateway/cfg"
	"sync"

	"github.com/qq51529210/uuid"
)

var (
	_sessions sessions
)

func init() {
	_sessions.d = make(map[string]*Session)
}

type sessions struct {
	sync.RWMutex
	d map[string]*Session
}

// Session 表示
type Session struct {
	Token   string
	Expired int64
}

// NewSession 创建会话
func NewSession() (*Session, error) {
	// 上锁
	_sessions.Lock()
	defer _sessions.Unlock()
	// 添加
	s := new(Session)
	s.Token = uuid.LowerV1WithoutHyphen()
	s.Expired = cfg.Cfg.Session.Expires
	return s, nil
}

// GetSession 返回会话
func GetSession(token string) (*Session, error) {
	// 上锁
	_sessions.RLock()
	defer _sessions.RUnlock()
	// 查询
	return _sessions.d[token], nil
}

// DelSession 删除会话
func DelSession(token string) error {
	// 上锁
	_sessions.Lock()
	defer _sessions.Unlock()
	// 查询
	delete(_sessions.d, token)
	return nil
}
