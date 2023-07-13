package db

import (
	"time"

	"github.com/qq51529210/uuid"
)

// Session 表示用户的会话
type Session struct {
	// 会话 ID
	ID string
	// 用户
	User *User
	// 创建时间
	Time int64
}

// NewSession 创建会话
func NewSession(user *User) (*Session, error) {
	s := &Session{
		ID:   uuid.LowerV1WithoutHyphen(),
		User: user,
		Time: time.Now().Unix(),
	}
	// 保存
	//
	return s, nil
}

// GetSession 获取会话
func GetSession(id string) (*Session, error) {
	return nil, nil
}
