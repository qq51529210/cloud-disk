package service

import (
	"net/http"
	"strings"

	"github.com/qq51529210/cloud-service/util"
	router "github.com/qq51529210/http-router"
	"google.golang.org/grpc"
)

type Server struct {
	httpSer http.Server
	grpcSer grpc.Server
}

func (s *Server) Serve(address, pageDir string, x509CertPEM, x509KeyPEM []byte) error {
	// Listener.
	listener, err := util.NewListener(address, x509CertPEM, x509KeyPEM)
	if err != nil {
		return err
	}
	// Router.
	var httpRouter router.Router
	// Pages.
	err = httpRouter.AddStatic(http.MethodGet, "/", pageDir, true)
	if err != nil {
		return err
	}
	// Routes.
	httpRouter.SetBefore(s.tokenInterceptor)
	httpRouter.AddGet("/api/:user", s.ApiGetUser)
	// Serve.
	if len(x509CertPEM) > 0 && len(x509KeyPEM) > 0 {
		s.httpSer.Handler = http.HandlerFunc(func(rs http.ResponseWriter, rq *http.Request) {
			if rq.ProtoMajor == 2 && strings.Contains(rq.Header.Get("Content-Type"), "application/grpc") {
				// Handle http2 grpc
				s.grpcSer.ServeHTTP(rs, rq)
				return
			}
			// Handle http
			httpRouter.ServeHTTP(rs, rq)
		})
	} else {
		s.httpSer.Handler = &httpRouter
	}
	return s.httpSer.Serve(listener)
}

func (s *Server) Close() error {
	s.grpcSer.Stop()
	return s.httpSer.Close()
}

func (s *Server) tokenInterceptor(c *router.Context) bool {
	return true
}
