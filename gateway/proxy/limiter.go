package proxy

import (
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	_limiter limiter
)

func init() {
	_limiter.d = make(map[string]*time.Time)
}

// limiter 用于对单个 ip 进行限流
type limiter struct {
	sync.Mutex
	d map[string]*time.Time
}

func limite(w http.ResponseWriter, r *http.Request, d time.Duration) bool {
	// 拿到 ip
	i := strings.LastIndex(r.RemoteAddr, ":")
	if i < 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		return false
	}
	ip := r.RemoteAddr[:i]
	now := time.Now()
	// 检查
	_limiter.Lock()
	defer _limiter.Unlock()
	//
	t := _limiter.d[ip]
	if t == nil {
		t = &now
		_limiter.d[ip] = t
		//
		return true
	}
	// 距离上一次的请求，时间很短
	if now.Sub(*t) <= d {
		w.WriteHeader(http.StatusTooManyRequests)
		return false
	}
	_limiter.d[ip] = &now
	//
	return true
}
