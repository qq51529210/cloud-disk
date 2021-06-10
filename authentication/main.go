package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/qq51529210/cloud-service/authentication/api"
	"github.com/qq51529210/cloud-service/authentication/db"
	"github.com/qq51529210/cloud-service/util"
)

var (
	httpSer                 http.Server
	x509CertPEM, x509KeyPEM []byte
)

func init() {
	// Load configure.
	data, err := util.ReadConfig()
	if err != nil {
		panic(err)
	}
	var cfg cfg
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	// Init.
	httpSer.Addr = cfg.Address
	x509CertPEM = []byte(cfg.X509CertPEM)
	x509KeyPEM = []byte(cfg.X509KeyPEM)
	err = cfg.Reg.Init()
	if err != nil {
		panic(err)
	}
	err = cfg.Cookie.Init()
	if err != nil {
		panic(err)
	}
	err = cfg.Redis.Init()
	if err != nil {
		panic(err)
	}
	err = db.InitMysql(cfg.Mysql)
	if err != nil {
		panic(err)
	}
}

type cfg struct {
	Address     string     `json:"address"`
	X509CertPEM string     `json:"x509CertPEM"`
	X509KeyPEM  string     `json:"x509KeyPEM"`
	Reg         *regCfg    `json:"reg"`
	Cookie      *cookieCfg `json:"cookie"`
	Redis       *redisCfg  `json:"redis"`
	Mysql       string     `json:"mysql"`
}

type regCfg struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	PhoneSMS    string `json:"phoneSMS"`
}

func (c *regCfg) Init() error {
	if c == nil {
		return nil
	}
	if c.Username != "" {
		api.UsernameRegexp = c.Username
	}
	if c.Password != "" {
		api.PasswordRegexp = c.Password
	}
	if c.PhoneNumber != "" {
		api.PhoneNumberRegexp = c.PhoneNumber
	}
	if c.PhoneSMS != "" {
		api.PhoneSMSRegexp = c.PhoneSMS
	}
	return api.InitRegExp()
}

type cookieCfg struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
	MaxAge string `json:"maxAge"`
}

func (c *cookieCfg) Init() error {
	if c == nil {
		return nil
	}
	if c.Name != "" {
		api.CookieName = c.Name
	}
	if c.Domain != "" {
		api.CookieDomain = c.Domain
	}
	if c.Path != "" {
		api.CookiePath = c.Path
	}
	if c.MaxAge != "" {
		n, err := strconv.ParseInt(c.MaxAge, 10, 64)
		if err != nil {
			return err
		}
		api.CookieMaxAge = int(n)
	}
	return nil
}

type redisCfg struct {
	TokenUrl string `json:"tokenUrl"`
	PhoneUrl string `json:"phoneUrl"`
}

func (c *redisCfg) Init() error {
	if c == nil {
		return nil
	}
	return db.InitRedis(c.TokenUrl, c.PhoneUrl)
}

func main() {
	listener, err := util.NewListener(httpSer.Addr, x509CertPEM, x509KeyPEM)
	if err != nil {
		panic(err)
	}
	err = httpSer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
