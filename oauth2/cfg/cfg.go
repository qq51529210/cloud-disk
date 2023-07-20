package cfg

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/qq51529210/util"
)

var (
	// Cfg 实例
	Cfg Config
)

// Config 表示程序启动配置
type Config struct {
	// 服务名称，日志使用
	Name string `json:"name" yaml:"name" validate:"required,max=32"`
	// 监听地址
	Addr string `json:"addr" yaml:"addr" validate:"required"`
	// 测试地址
	Test string `json:"test" yaml:"test"`
	// 日志配置
	Log util.LogCfg `json:"log" yaml:"log"`
	// 数据库配置
	DB struct {
		URL string `json:"url" yaml:"url" validate:"required"`
	} `json:"db" yaml:"db"`
	// 缓存配置
	Redis struct {
		Name             string   `json:"name" yaml:"name"`
		Addrs            []string `json:"addrs" yaml:"addrs" validate:"required"`
		DB               int      `json:"db" yaml:"db" validate:"min=0,max=32"`
		Username         string   `json:"username" yaml:"username"`
		Password         string   `json:"password" yaml:"password"`
		Master           string   `json:"master" yaml:"master"`
		SentinelUsername string   `json:"sentinelUsername" yaml:"sentinelUsername"`
		SentinelPassword string   `json:"sentinelPassword" yaml:"sentinelPassword"`
		CmdTimeout       int64    `json:"cmdTimeout" yaml:"cmdTimeout" validate:"required,min=1"`
	} `json:"redis" yaml:"redis"`
	// 会话配置
	Session struct {
		Expires int64 `json:"expires" yaml:"expires" validate:"required,min=60"`
	} `json:"session" yaml:"session"`
}

// Load 加载配置
func Load() error {
	// 路径，默认程序同目录下的 cfg.yaml 文件
	path := "cfg.yaml"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	// 加载
	err := util.ReadCfg(path, &Cfg)
	if err != nil {
		return err
	}
	// 检查字段
	val := validator.New()
	err = val.Struct(&Cfg)
	if err != nil {
		return err
	}
	//
	return nil
}
