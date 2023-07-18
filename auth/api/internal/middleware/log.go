package middleware

import (
	"auth/api/internal"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/log"
)

const (
	// LogDataContextKey 是 gin.Context 的 data 上下文 key
	// 用于 log 保存提交的 body 的数据
	LogDataContextKey = "reqdata"
)

// Log 用于记录日志，应该作为全局第一个中间件
func Log(ctx *gin.Context) {
	now := time.Now()
	remoteAddr := ctx.Request.RemoteAddr
	if addr := ctx.GetHeader("X-Remote-Addr"); addr != "" {
		remoteAddr = addr
	}
	// 清理
	defer func() {
		// 花费时间
		cost := time.Since(now)
		// 日志
		var str strings.Builder
		fmt.Fprintf(&str, "%s %s %s cost %v", remoteAddr, ctx.Request.Method, ctx.Request.RequestURI, cost)
		// 日志数据
		data := ctx.Value(LogDataContextKey)
		if data != nil {
			str.WriteByte('\n')
			json.NewEncoder(&str).Encode(data)
		}
		re := recover()
		if re != nil {
			err := fmt.Errorf("%v", re)
			internal.Error500(ctx, err)
			str.WriteByte('\n')
			str.WriteString(err.Error())
		}
		log.Debug(str.String())
	}()
	// 执行
	ctx.Next()
}
