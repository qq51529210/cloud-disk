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
		p := new(UploadFile)
		p.hash = sha1.New()
		p.buff = make([]byte, defaultuploadFileBuffer)
		return p
	}
	downloadFilePool.New = func() interface{} {
		p := new(DownloadFile)
		p.buff = make([]byte, defaultdownloadFileBuffer)
		return p
	}
}

// Return file temp name.
func FileTempName(namespace, name string) string {
	var str strings.Builder
	str.WriteString(namespace)
	str.WriteByte('.')
	str.WriteString(name)
	str.WriteByte('.')
	str.WriteString(time.Now().Format("20060102150405.000"))
	return str.String()
}

// Return file name.
func FileName(namespace, name string) string {
	var str strings.Builder
	str.WriteString(namespace)
	str.WriteByte('.')
	str.WriteString(name)
	return str.String()
}

// To handle upload file.
type UploadFile struct {
	file *os.File
	hash hash.Hash
	buff []byte
}

// Read data from r.
func (f *UploadFile) Upload(r io.Reader, dir, namespace, name string, rate, dur int) (err error) {
	// Make sure there is a directory for new file.
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return
	}
	// Create file.
	filePath := filepath.Join(dir, FileTempName(namespace, name))
	f.file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
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
	f.hash.Reset()
	_, err = util.LimitRateCopy(f, r, f.buff, rate, dur)
	if err != nil {
		f.file.Close()
		return
	}
	err = f.file.Close()
	if err != nil {
		return
	}
	// New file name by file hex hash value.
	hashFilePath := filepath.Join(dir, hex.EncodeToString(f.hash.Sum(f.buff[:0])))
	// Rename file temp name if not exist.
	_, err = os.Stat(hashFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Rename(filePath, hashFilePath)
			if err != nil {
				return
			}
		}
	} else {
		// File extists, remove temp file.
		err = os.Remove(filePath)
		if err != nil {
			return
		}
	}
	// Create a symbolic link if not exist.
	linkName := filepath.Join(dir, FileName(namespace, name))
	_, err = os.Stat(linkName)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Symlink(hashFilePath, linkName)
		}
	}
	return
}

// For io.Copy
func (f *UploadFile) ReadFrom(r io.Reader) (n int64, err error) {
	var m int
	for {
		m, err = r.Read(f.buff)
		if err != nil {
			if err == io.EOF {
				f.Write(f.buff[:m])
				n += int64(m)
				err = nil
				return
			}
		}
		f.Write(f.buff[:m])
		n += int64(m)
	}
}

func (f *UploadFile) Write(b []byte) (int, error) {
	f.hash.Write(b)
	return f.file.Write(b)
}

// To handle download file.
type DownloadFile struct {
	file *os.File
	buff []byte
}

// Open file and return size or error.
func (f *DownloadFile) Open(dir, namespace, name string) (int64, error) {
	filePath := filepath.Join(dir, FileName(namespace, name))
	fi, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	f.file, err = os.Open(filePath)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func (f *DownloadFile) Download(w io.Writer, rate, dur int) (err error) {
	defer f.file.Close()
	_, err = util.LimitRateCopy(w, f.file, f.buff, rate, dur)
	return
}
