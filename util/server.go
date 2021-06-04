package util

import (
	"crypto/tls"
	"net"
)

// If x509Cert and x509Key both are not empty, return tls.Lisenter.
// other wise, return net.TcpListener.
func NewListener(address, x509Cert, x509Key string) (net.Listener, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	// tls
	if x509Cert != "" && x509Key != "" {
		certificate, err := tls.X509KeyPair([]byte(x509Cert), []byte(x509Key))
		if err != nil {
			return nil, err
		}
		listener = tls.NewListener(listener, &tls.Config{
			Certificates: []tls.Certificate{certificate},
		})
	}
	return listener, nil
}
