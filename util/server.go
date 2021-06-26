package util

import (
	"crypto/tls"
	"net"
	"net/http"
	"strings"

	router "github.com/qq51529210/http-router"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

// If x509CertPEM and x509KeyPEM both are not empty, return tls.Lisenter.
// other wise, return net.TcpListener.
func NewListener(address string, x509CertPEM, x509KeyPEM []byte) (net.Listener, error) {
	if len(x509CertPEM) > 0 && len(x509KeyPEM) > 0 {
		cert, err := tls.X509KeyPair(x509CertPEM, x509KeyPEM)
		if err != nil {
			return nil, err
		}
		return tls.Listen("tcp", address, &tls.Config{
			Certificates: []tls.Certificate{cert},
			NextProtos:   []string{http2.NextProtoTLS},
		})
	} else {
		return net.Listen("tcp", address)
	}
}

// A server support http2 and grpc.
type Server struct {
	HTTP        http.Server
	GRPG        *grpc.Server
	Router      router.Router
	X509CertPEM []byte
	X509KeyPEM  []byte
}

func (s *Server) Serve() error {
	// Should init grpc server first
	if s.GRPG == nil {
		return Error("must init grpc server first")
	}
	// Load x509.
	cert, err := tls.X509KeyPair(s.X509CertPEM, s.X509KeyPEM)
	if err != nil {
		return err
	}
	// Create listener.
	listener, err := tls.Listen("tcp", s.HTTP.Addr, &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{http2.NextProtoTLS},
	})
	if err != nil {
		return err
	}
	// ServeHTTP
	s.HTTP.Handler = http.HandlerFunc(func(rs http.ResponseWriter, rq *http.Request) {
		if rq.ProtoMajor == 2 && strings.Contains(rq.Header.Get("Content-Type"), "application/grpc") {
			// Handle http2 grpc
			s.GRPG.ServeHTTP(rs, rq)
			return
		}
		// Handle http
		s.Router.ServeHTTP(rs, rq)
	})
	return s.HTTP.Serve(listener)
}
