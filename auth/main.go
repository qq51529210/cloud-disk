package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/api"
	"github.com/qq51529210/micro-services/auth/db"
	"github.com/qq51529210/micro-services/auth/reg"
	"github.com/qq51529210/micro-services/util"
	"github.com/qq51529210/uuid"
)

type uuidConfig struct {
	SnowflakeGroupID   byte   `json:"snowflakeGroupID,omitempty"`
	SnowflakeMechineID byte   `json:"snowflakeMechineID,omitempty"`
	V1Node             string `json:"v1Node,omitempty"`
}

type config struct {
	util.Config
	Api  api.Config `json:"api,omitempty"`
	Reg  reg.Config `json:"reg,omitempty"`
	DB   db.Config  `json:"db,omitempty"`
	UUID uuidConfig `json:"uuid,omitempty"`
}

func loadConfig() *config {
	//
	data, err := util.ReadConfig()
	if err != nil {
		panic(err)
	}
	//
	cfg := new(config)
	err = json.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	//
	cfg.Config.Check()
	// Reg
	reg.Init(&cfg.Reg)
	// DB
	db.Init(&cfg.DB)
	// UUID Snowflake
	uuid.SetSnowflakeGroupID(cfg.UUID.SnowflakeGroupID)
	uuid.SetSnowflakeMechineID(cfg.UUID.SnowflakeMechineID)
	// UUID v1 node
	b := []byte(cfg.UUID.V1Node)
	if len(b) != 6 {
		panic(fmt.Errorf("uuid v1 node must be 6 character"))
	}
	var v1node [6]byte
	copy(v1node[:], b)
	uuid.SetV1Node(v1node)
	//
	return cfg
}

func initServer() *util.Server {
	// Config
	cfg := loadConfig()
	// Server
	ser := new(util.Server)
	ser.HTTP.Addr = cfg.Config.Address
	ser.X509CertPEM = []byte(strings.Join(cfg.Config.X509Cert, ""))
	ser.X509KeyPEM = []byte(strings.Join(cfg.Config.X509Key, ""))
	// Router.
	pageDir := cfg.Config.RootDir
	if pageDir == "" {
		pageDir = filepath.Join(filepath.Dir(os.Args[0]), "page")
	}
	api.Init(&ser.Router, &cfg.Api)
	//
	return ser
}

func main() {
	log.SetFormatStackHeader(log.FormatFileNameStackHeader)
	err := initServer().Serve()
	if err != nil {
		panic(err)
	}
}
