package main

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/qq51529210/cloud-service/util"
	router "github.com/qq51529210/http-router"
	"github.com/qq51529210/log"
)

type HTTPServer struct {
	server http.Server
	router router.Router
	// Listen address.
	Address string
	// X509 public key certificate data.
	X509Cert string
	// X509 private key certificate data.
	X509Key string
	// File directory.
	FileDir string
	// Rate limit timer duration.
	RateDur int
}

func (s *HTTPServer) Serve() error {
	// Router.
	s.router.SetNotfound(router.Notfound)
	s.router.AddPost("/:dir", s.parseFileToken, s.PostFile)
	s.router.AddGet("/:dir/:file", s.parseFileToken, s.GetFile)
	s.router.AddGet("/hashes/:hash", s.parseFileToken, s.GetFile)
	s.server.Handler = &s.router
	// Serve.
	listener, err := util.NewListener(s.Address, s.X509Cert, s.X509Key)
	if err != nil {
		return err
	}
	s.server.Handler = &s.router
	return s.server.Serve(listener)
}

// Parse query param token, assign to c.Data.
func (s *HTTPServer) parseFileToken(c *router.Context) bool {
	err := c.Req.ParseForm()
	if err != nil {
		log.Error(err)
		return false
	}
	token := c.Req.Form.Get("token")
	if token == "" {
		c.WriteJSON(http.StatusUnauthorized, map[string]string{
			"error": "Token is required.",
		})
		return false
	}
	c.Data = token
	return true
}

// File upload.
func (s *HTTPServer) PostFile(c *router.Context) bool {
	// Parse "multipart/data".
	mediaType, params, err := mime.ParseMediaType(c.Req.Header.Get("Content-Type"))
	if err != nil {
		log.Error(err)
		return false
	}
	if mediaType != "multipart/form-data" {
		// todo: configure response body describe.
		c.WriteJSON(http.StatusBadRequest, map[string]string{
			"error": "Content-Type must be multipart/form-data.",
		})
		return false
	}
	// Get upload information.
	info, err := ApiGetUploadInfo(c.Data.(string))
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]string{
			"error": "Query token failed.",
		})
		return false
	}
	if info == nil {
		c.WriteJSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid token.",
		})
		return false
	}
	// Upload.
	file := uploadFilePool.Get().(*UploadFile)
	defer uploadFilePool.Put(file)
	reader := multipart.NewReader(c.Req.Body, params["boundary"])
	for {
		p, err := reader.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Error(err)
			return false
		}
		name := p.FileName()
		if name == "" {
			name = p.FormName()
		}
		err = file.Upload(p, s.FileDir, c.SHA1(c.Param[0]), c.SHA1(name), info.Rate, s.RateDur)
		if err != nil {
			log.Error(err)
			return false
		}
	}
	return true
}

// File download.
func (s *HTTPServer) GetFile(c *router.Context) bool {
	// Get download information.
	info, err := ApiGetDownloadInfo(c.Data.(string))
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusInternalServerError, map[string]string{
			"error": "Query token failed.",
		})
		return false
	}
	if info == nil {
		c.WriteJSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid token.",
		})
		return false
	}
	// Download.
	file := downloadFilePool.Get().(*DownloadFile)
	defer downloadFilePool.Put(file)
	// File.
	size, err := file.Open(s.FileDir, c.SHA1(c.Param[0]), c.SHA1(c.Param[1]))
	if err != nil {
		c.WriteJSON(http.StatusNotFound, map[string]string{
			"error": fmt.Sprintf("file %s not found.", c.Param[1]),
		})
		return false
	}
	// Response header
	c.Res.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	c.Res.Header().Set("Content-Type", mime.TypeByExtension(c.Param[1]))
	// Response body.
	err = file.Download(c.Res, info.Rate, s.RateDur)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

// Get file hash, response 200 if exists, 404 if not.
func (s *HTTPServer) GetFileHash(c *router.Context) bool {
	_, err := os.Stat(c.Param[0])
	if err != nil {
		c.WriteJSON(http.StatusNotFound, map[string]string{
			"error": fmt.Sprintf("hash %s not found", c.Param[0]),
		})
		return false
	}
	return true
}
