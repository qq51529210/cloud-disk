package cfg

import (
	"net"
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
	// 管理地址
	AdminAddr string `json:"adminAddr" yaml:"adminAddr" validate:"required,ipAddr"`
	// 代理地址
	ProxyAddr string `json:"proxyAddr" yaml:"proxyAddr" validate:"required,ipAddr"`
	// 日志配置
	Log util.LogCfg `json:"log" yaml:"log"`
	// 数据库配置
	DB struct {
		URL string `json:"url" yaml:"url" validate:"required"`
	} `json:"db" yaml:"db"`
	// 缓存配置
	Redis struct {
		// 客户端名称
		Name string `json:"name" yaml:"name"`
		// 服务地址
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required"`
		// 数据库
		DB int `json:"db" yaml:"db" validate:"min=0,max=32"`
		// 用户名 6.0 以上
		Username string `json:"username" yaml:"username"`
		// 密码
		Password string `json:"password" yaml:"password"`
		// 集群
		Master string `json:"master" yaml:"master"`
		// 哨兵用户名
		SentinelUsername string `json:"sentinelUsername" yaml:"sentinelUsername"`
		// 哨兵密码
		SentinelPassword string `json:"sentinelPassword" yaml:"sentinelPassword"`
		// 执行一次命令的超时时间
		CmdTimeout int64 `json:"cmdTimeout" yaml:"cmdTimeout" validate:"required,min=1"`
	} `json:"redis" yaml:"redis"`
	// 会话配置
	Session struct {
		// 过期时间
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
	err = val.RegisterValidation("ipAddr", validateIPAddr)
	if err != nil {
		return err
	}
	err = val.Struct(&Cfg)
	if err != nil {
		return err
	}
	//
	return nil
}

func validateIPAddr(fl validator.FieldLevel) bool {
	a, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	_, err := net.ResolveTCPAddr("tcp", a)
	return err == nil
}
