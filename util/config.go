package util

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// If app has not args, read from file "app_dir/app.json",
// otherwise read from os.Args[1], which can be a http url, a local file.
func ReadConfig() ([]byte, error) {
	var uri string
	if len(os.Args) < 2 {
		// No arg.
		dir, file := filepath.Split(os.Args[0])
		ext := filepath.Ext(file)
		if ext != "" {
			file = file[:len(file)-len(ext)]
		}
		uri = filepath.Join(dir, file+".json")
	} else {
		// Use first arg.
		uri = os.Args[1]
	}
	// uri is a http url.
	if strings.HasPrefix(uri, "http") || strings.HasPrefix(uri, "https") {
		res, err := http.Get(uri)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		return ioutil.ReadAll(res.Body)
	}
	// uri is a local file path.
	return ioutil.ReadFile(uri)
}

type Config struct {
	Address  string   `json:"address"`  // Server listen address.
	X509Cert []string `json:"x509Cert"` // Pem format.
	X509Key  []string `json:"x509Key"`  // Pem format.
	RootDir  string   `json:"rootDir"`
}

func (c *Config) Check() {
	if c.Address == "" {
		c.Address = ":0"
	}
	if c.RootDir == "" {
		c.RootDir = filepath.Dir(os.Args[0])
	}
}

type CookieConfig struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
	MaxAge int64  `json:"maxAge"`
}
