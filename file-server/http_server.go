package main

import (
	"io"
	"mime"
	"mime/multipart"
	"net/http"

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
	s.router.AddGet("/:dir/:file", s.parseFileToken, s.GetFile)
	s.router.AddPost("/:dir", s.parseFileToken, s.PostFile)
	s.server.Handler = &s.router
	// Serve.
	listener, err := util.NewListener(s.Address, s.X509Cert, s.X509Key)
	if err != nil {
		return err
	}
	s.server.Handler = &s.router
	return s.server.Serve(listener)
}

// Handle file upload request.
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
	var info uploadInfo
	err = apiGetUploadInfo(c.Data.(string), &info)
	if err != nil {
		log.Error(err)
		c.WriteJSON(http.StatusBadRequest, map[string]string{
			"error": "Content-Type must be multipart/form-data.",
		})
		return false
	}
	// Init uploadFile.
	file := uploadFilePool.Get().(*UploadFile)
	defer uploadFilePool.Put(file)
	file.dir = s.FileDir
	file.namespace = c.SHA1(c.Param[0])
	file.rate = info.Rate
	file.dur = s.RateDur
	// Save file.
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
		file.name = c.SHA1(p.FileName())
		_, err = file.ReadFrom(p)
		if err != nil {
			log.Error(err)
			return false
		}
	}
	return true
}

// Handle file download request.
func (s *HTTPServer) GetFile(c *router.Context) bool {

	return true
}

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
