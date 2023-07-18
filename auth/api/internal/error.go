package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/log"
)

// Error 用于返回错误的响应
type Error struct {
	// 简短信息
	Phrase string `json:"phrase,omitempty"`
	// 详细信息
	Detail string `json:"detail,omitempty"`
}

// Submit400 表示提交的数据错误
func Submit400(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusBadRequest, &Error{
		Phrase: "error submit data",
		Detail: err,
	})
	ctx.Abort()
}

// SubmitEmpty400 表示提交的数据错误
func SubmitEmpty400(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, &Error{
		Phrase: "empty submit data",
	})
	ctx.Abort()
}

// Data404 表示数据不存在
func Data404(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, &Error{
		Phrase: "data not found",
	})
	ctx.Abort()
}

// DB500 表示数据库操作错误
func DB500(ctx *gin.Context, err error) {
	log.Error(err)
	ctx.JSON(http.StatusInternalServerError, &Error{
		Phrase: "database error",
	})
	ctx.Abort()
}

// Error500 表示服务错误
func Error500(ctx *gin.Context, err error) {
	log.Error(err)
	ctx.JSON(http.StatusInternalServerError, &Error{
		Phrase: "server error",
	})
	ctx.Abort()
}

// Error502 表示远程调用错误
func Error502(ctx *gin.Context, err error) {
	log.Error(err)
	ctx.JSON(http.StatusBadGateway, &Error{
		Phrase: "api call error",
		Detail: err.Error(),
	})
	ctx.Abort()
}

// Error504 表示远程调用超时
func Error504(ctx *gin.Context, err error) {
	log.Error(err)
	ctx.JSON(http.StatusGatewayTimeout, &Error{
		Phrase: "api call timeout",
		Detail: err.Error(),
	})
	ctx.Abort()
}
