package config

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config 表示程序启动的一些配置参数，支持 json/yaml/xml 三种格式的数据
type Config struct {
	// 日志
	Log logConfig `json:"log" yaml:"log" xml:"Log"`
	// HTTP
	HTTP httpConfig `json:"http" yaml:"http" xml:"Http"`
	// DB
	DB dbConfig `json:"db" yaml:"db" xml:"DB"`
	//
	Data any
}

// readFromFile 从本地文件中加载配置
func (c *Config) readFromFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	switch strings.ToLower(filepath.Ext(path)) {
	case ".json":
		err = json.Unmarshal(data, c)
	case ".yml", ".yaml":
		err = yaml.Unmarshal(data, c)
	case "xml":
		err = xml.Unmarshal(data, c)
	default:
		panic(fmt.Errorf("unsupport file ext %s", filepath.Ext(path)))
	}
	if err != nil {
		panic(err)
	}
}

// readFromHTTP 从 web 服务读取配置数据
func (c *Config) readFromHTTP(url string) {
	res, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
}

// Load 加载并返回配置。
// 如果没有输入参数，则加载程序目录下的 appname.json 文件。
// file://filepath.json，本地 json 文件。
// file://filepath.yml 本地 yml 文件。
// file://filepath.yaml 本地 yaml 文件。
// http://path，从 http 读取。
// https://path，从 https 读取。
func Load() *Config {
	var cfg *Config
	// 读取
	if len(os.Args) > 1 {
		_url, err := url.Parse(os.Args[1])
		if err != nil {
			panic(err)
		}
		switch _url.Scheme {
		case "http", "https":
			cfg.readFromHTTP(os.Args[1])
		default:
			cfg.readFromFile(_url.Path)
		}
	} else {
		// 程序目录下的appname.json
		cfg.readFromFile(os.Args[0] + ".json")
	}
	// 读取环境变量
	cfg.Log.ReadEnv()
	cfg.HTTP.ReadEnv()
	cfg.DB.ReadEnv()
	// 检查字段
	cfg.Log.Check()
	cfg.HTTP.Check()
	cfg.DB.Check()
	//
	return cfg
}
