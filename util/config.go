package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// 使用命令行第一个参数作为路径(本地磁盘/http资源)，解析JSON 到参数a。
// 如果没有参数，则尝试加载当前目录下的名为"{app-name}.json"的文件。
func ReadJSONConf(a interface{}) error {
	var path string
	if len(os.Args) < 2 {
		dir, file := filepath.Split(os.Args[0])
		ext := filepath.Ext(file)
		if ext != "" {
			file = file[:len(file)-len(ext)]
		}
		path = filepath.Join(dir, file+".json")
	} else {
		path = os.Args[1]
	}
	var data []byte
	var err error
	if strings.HasPrefix(path, "http") || strings.HasPrefix(path, "https") {
		res, err := http.Get(path)
		if err != nil {
			return err
		}
		data, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
	} else {
		data, err = ioutil.ReadFile(path)
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(data, a)
}

type Conf struct {
	Listen      string        `json:"listen"`
	X509CertPEM []string      `json:"x509CertPEM"`
	X509KeyPEM  []string      `json:"x509KeyPEM"`
	RootDir     string        `json:"rootDir"`
	Cookie      *CookieConf   `json:"cookie"`
	UUID        *UUIDConf     `json:"uuid"`
	Database    *DatabaseConf `json:"database"`
	Cache       *CacheConf    `json:"cache"`
	Custom      interface{}   `json:"custom"`
}

type CookieConf struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
	MaxAge int64  `json:"maxAge"`
}

type UUIDConf struct {
	GroupID   byte   `json:"groupID"`
	MechineID byte   `json:"mechineID"`
	Node      string `json:"node"`
}

type CacheConf struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
}

type DatabaseConf struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
}
