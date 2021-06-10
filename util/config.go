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
