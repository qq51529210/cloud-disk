package util

import (
	"crypto/tls"
	"net"

	"golang.org/x/net/http2"
)

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
