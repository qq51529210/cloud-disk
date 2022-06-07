package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/qq51529210/log"
)

// logConfig 表示日志的配置
type logConfig struct {
	// 生成者
	AppID string `json:"appID" yaml:"appID" xml:"AppID"`
	// 级别
	Level []string `json:"levels" yaml:"levels" xml:"Levels"`
	// 头格式化
	Format string `json:"format" yaml:"format" xml:"Format"`
	// 保存的目录，默认是当前目录下的 log 文件夹
	Dir string `json:"dir" yaml:"dir" xml:"Dir"`
	// 每一份日志文件的最大字节，使用 1.5/K/M/G/T 这样的字符表示。
	MaxFileSize string `json:"maxFileSize" yaml:"maxFileSize" xml:"MaxFileSize"`
	// 保存的最大天数，最小是1天。
	MaxKeepDay float64 `json:"maxKeepDay" yaml:"maxKeepDay" xml:"MaxKeepDay"`
	// 同步到磁盘的时间间隔，单位，毫秒。最小是10毫秒。
	SyncInterval int `json:"syncInterval" yaml:"syncInterval" xml:"SyncInterval"`
}

// Check 检测字段
func (c *logConfig) Check() {
	// LogDir
	if c.Dir == "" {
		c.Dir = filepath.Join(filepath.Dir(os.Args[0]), "log")
	}
	// 日志级别
	for _, s := range c.Level {
		if strings.ToUpper(s) == "DEBUG" {
			log.SetLevel(log.DebugLevel)
			continue
		}
		if strings.ToUpper(s) == "INFO" {
			log.SetLevel(log.InfoLevel)
			continue
		}
		if strings.ToUpper(s) == "WARN" {
			log.SetLevel(log.WarnLevel)
			continue
		}
		if strings.ToUpper(s) == "ERROR" {
			log.SetLevel(log.ErrorLevel)
			continue
		}
	} // 日志头
	log.SetHeaderFormater(log.NewHeaderFormater(log.HeaderFormaterType(c.Format), c.AppID))
	log.NewFile(&log.FileConfig{
		RootDir:      c.Dir,
		MaxFileSize:  c.MaxFileSize,
		MaxKeepDay:   c.MaxKeepDay,
		SyncInterval: c.SyncInterval,
	})
}

// ReadEnv 使用环境变量替换掉配置的字段
func (c *logConfig) ReadEnv() {
	// AppID
	str := os.Getenv("LOG_APP_ID")
	if str != "" {
		c.AppID = str
	}
	// Level
	str = os.Getenv("LOG_LEVEL")
	if str != "" {
		c.Level = strings.Split(str, ",")
	}
	// Format
	str = os.Getenv("LOG_FORMAT")
	if str != "" {
		c.Format = str
	}
	// Dir
	str = os.Getenv("LOG_DIR")
	if str != "" {
		c.Dir = str
	}
}
