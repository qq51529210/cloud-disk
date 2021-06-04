package main

import (
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/qq51529210/cloud-service/util"
	"github.com/qq51529210/log"
)

const (
	defaultuploadFileBuffer         = 1024 * 64
	defaultuploadFileTickDuration   = 30
	defaultdownloadFileBuffer       = 1024 * 64
	defaultdownloadFileTickDuration = 30
)

var (
	uploadFilePool   sync.Pool
	downloadFilePool sync.Pool
)

func init() {
	uploadFilePool.New = func() interface{} {
		p := new(uploadFile)
		p.Hash = sha1.New()
		p.sum = make([]byte, sha1.Size)
		p.RateLimiter.Buff = make([]byte, defaultuploadFileBuffer)
		return p
	}
	downloadFilePool.New = func() interface{} {
		p := new(downloadFile)
		p.RateLimiter.Buff = make([]byte, defaultdownloadFileBuffer)
		return p
	}
}

// To handle upload file.
type uploadFile struct {
	util.RateLimiter
	*os.File
	hash.Hash
	sum       []byte
	dir       string
	namespace string
	name      string
}

// Read data from r and save to dir/name,
// rate is bytes per second,
// dur is timer duration, millisecond, default is defaultuploadFileTickDuration.
// dur need to be test.
func (f *uploadFile) ReadFrom(r io.Reader) (n int64, err error) {
	// Make sure there is a directory for new file.
	err = os.MkdirAll(f.dir, os.ModePerm)
	if err != nil {
		return
	}
	// Create file.
	filePath := filepath.Join(f.dir, f.FileTempName())
	f.File, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			// If there is any error, remove file.
			re := os.Remove(filePath)
			if re != nil {
				log.Error(re)
			}
		}
	}()
	// Save data to file.
	f.Hash.Reset()
	n, err = f.RateLimiter.Copy(f, r)
	if err != nil {
		f.File.Close()
		return
	}
	err = f.File.Close()
	if err != nil {
		return
	}
	f.Hash.Sum(f.sum[:0])
	// New file name by file hex hash value.
	newFilePath := filepath.Join(f.dir, hex.EncodeToString(f.sum))
	_, err = os.Stat(newFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Rename(filePath, newFilePath)
			if err != nil {
				return
			}
		}
	}
	// Create a symbolic link.
	_, err = os.Stat(f.name)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Symlink(newFilePath, f.name)
		}
	}
	return
}

// Return file name, namespace.name.time.
func (f *uploadFile) FileTempName() string {
	var str strings.Builder
	str.WriteString(f.namespace)
	str.WriteByte('.')
	str.WriteString(f.name)
	str.WriteByte('.')
	str.WriteString(time.Now().Format("20060102150405.000"))
	return str.String()
}

func (f *uploadFile) Write(b []byte) (int, error) {
	f.Hash.Write(b)
	return f.File.Write(b)
}

// To handle download file.
type downloadFile struct {
	util.RateLimiter
}

func (f *downloadFile) WriteTo(w io.Writer, dir, namespace, name string, dur int) (err error) {
	return nil
}
