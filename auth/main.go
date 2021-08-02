package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/qq51529210/cloud-service/authentication/api"
	"github.com/qq51529210/cloud-service/authentication/db"
	"github.com/qq51529210/cloud-service/authentication/reg"
	"github.com/qq51529210/cloud-service/util"
)

type cfg struct {
	Address     string     `json:"address"`
	X509CertPEM []string   `json:"x509CertPEM"`
	X509KeyPEM  []string   `json:"x509KeyPEM"`
	PageDir     string     `json:"pageDir"`
	Cookie      *cookieCfg `json:"cookie"`
	Reg         *regCfg    `json:"reg"`
	DB          *dbCfg     `json:"db"`
}

func (c *cfg) Init() error {
	err := c.Reg.Init()
	if err != nil {
		return err
	}
	err = c.Cookie.Init()
	if err != nil {
		return err
	}
	err = c.DB.Init()
	if err != nil {
		return err
	}
	return nil
}

type regCfg struct {
	Account     string `json:"account"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	PhoneCode   string `json:"phoneVerificationCode"`
}

func (c *regCfg) Init() error {
	if c == nil {
		return nil
	}
	err := reg.Account.Compile(c.Account)
	if err != nil {
		return err
	}
	err = reg.Password.Compile(c.Password)
	if err != nil {
		return err
	}
	err = reg.PhoneNumber.Compile(c.PhoneNumber)
	if err != nil {
		return err
	}
	err = reg.PhoneVerificationCode.Compile(c.PhoneCode)
	if err != nil {
		return err
	}
	return nil
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
	api.CookieName = c.Name
	api.CookieDomain = c.Domain
	api.CookiePath = c.Path
	if c.MaxAge != "" {
		n, err := strconv.ParseInt(c.MaxAge, 10, 64)
		if err != nil {
			return err
		}
		api.CookieMaxAge = int(n)
	}
	return nil
}

type dbCfg struct {
	TokenUrl string `json:"tokenUrl"`
	PhoneUrl string `json:"phoneUrl"`
	MysqlUrl string `json:"mysqlUrl"`
}

func (c *dbCfg) Init() error {
	if c == nil {
		return nil
	}
	err := db.InitTokenRedis(c.TokenUrl)
	if err != nil {
		return err
	}
	err = db.InitPhoneNumberRedis(c.PhoneUrl)
	if err != nil {
		return err
	}
	err = db.InitMysql(c.MysqlUrl)
	if err != nil {
		return err
	}
	return nil
}

func initServer() *util.Server {
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
	err = cfg.Init()
	if err != nil {
		panic(err)
	}
	// Server.
	var ser util.Server
	ser.HTTP.Addr = cfg.Address
	ser.X509CertPEM = []byte(strings.Join(cfg.X509CertPEM, ""))
	ser.X509KeyPEM = []byte(strings.Join(cfg.X509KeyPEM, ""))
	ser.GRPG = api.InitGRPG()
	// HTTP router.
	pageDir := cfg.PageDir
	if pageDir == "" {
		pageDir = filepath.Join(filepath.Dir(os.Args[0]), "page")
	}
	err = api.InitRouter(&ser.Router, pageDir)
	if err != nil {
		panic(err)
	}
	return &ser
}

func main() {
	err := initServer().Serve()
	if err != nil {
		panic(err)
	}
}
