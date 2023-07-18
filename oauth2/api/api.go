package api

import (
	"io"
	"io/fs"
	"net/http"
	"oauth2/api/apps"
	"oauth2/api/internal/middleware"
	"oauth2/api/login"
	"oauth2/api/oauth2"
	"oauth2/api/test"
	"oauth2/api/users"
	"oauth2/cfg"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	// gin
	g = gin.New()
)

// Serve 开始服务
func Serve(statics fs.FS) error {
	gin.SetMode(gin.ReleaseMode)
	// 测试服务
	if cfg.Cfg.Test != "" {
		// 启动服务
		go testServer()
	}
	// 静态文件
	initStatic(statics)
	// 路由
	initRouter()
	// 监听
	return http.ListenAndServe(cfg.Cfg.Addr, g)
}

// 初始化路由
func initRouter() {
	// 全局
	g.Use(middleware.Log)
	//
	apps.Init(g)
	users.Init(g)
	login.Init(g)
	oauth2.Init(g)
}

// 初始化静态文件
func initStatic(statics fs.FS) (err error) {
	index := "/index.html"
	return fs.WalkDir(statics, ".", func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		g.StaticFileFS(path, path, http.FS(statics))
		if strings.HasSuffix(path, index) {
			initStaticIndex(statics, path, index)
		}
		return nil
	})
}

// 以免 gin 内部对 index.html 一直重定向
func initStaticIndex(statics fs.FS, path, index string) {
	g.GET(path[:len(path)-len(index)], func(ctx *gin.Context) {
		f, err := statics.Open(path)
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}
		//
		ctx.Writer.Header().Set("Content-Type", gin.MIMEHTML)
		//
		io.Copy(ctx.Writer, f)
	})
}

// 模拟第三方服务
func testServer() {
	//
	g := gin.New()
	// 路由
	test.Init(g)
	// 监听
	http.ListenAndServe(cfg.Cfg.Test, g)
}
