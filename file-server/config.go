package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	HTTP     *httpConfig    `json:"http"`
	GRPC     *grpcConfig    `json:"grpc"`
	Manager  *managerConfig `json:"manager"`
	X509Cert string         `json:"x509Cert"`
	X509Key  string         `json:"x509Key"`
}

type httpConfig struct {
	// listen address.
	Address string `json:"address"`
	// X509 public key certificate data.
	X509Cert string `json:"x509Cert"`
	// X509 private key certificate data.
	X509Key string `json:"x509Key"`
	// Directory where the upload file save.
	FileDir string `json:"fileDir"`
	// Timer duration of rate.
	RateDur int `json:"rateDur"`
}

type grpcConfig struct {
	// listen address.
	Address string `json:"address"`
}

type managerConfig struct {
	// listen address.
	Address string `json:"address"`
}

func loadConfig() *config {
	var data []byte
	var err error
	if len(os.Args) < 2 {
		// No args, "appName.json" will be used as the configuration file.
		dir, file := filepath.Split(os.Args[0])
		data, err = ioutil.ReadFile(filepath.Join(dir, file+".json"))
	} else {
		if strings.HasPrefix(os.Args[1], "http") || strings.HasPrefix(os.Args[1], "https") {
			// It's a http url.
			var res *http.Response
			res, err = http.Get(os.Args[1])
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			data, err = ioutil.ReadAll(res.Body)
		} else {
			// It's a local file path.
			data, err = ioutil.ReadFile(os.Args[1])
		}
	}
	if err != nil {
		panic(err)
	}
	cfg := new(config)
	err = json.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
