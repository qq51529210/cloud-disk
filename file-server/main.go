package main

import "os"

func main() {
	exit := make(chan os.Signal, 1)
	go func() {
		cfg := loadConfig()
		// HTTP service.
		{
			hs := new(HTTPServer)
			hs.Address = cfg.HTTP.Address
			hs.X509Cert = cfg.HTTP.X509Cert
			hs.X509Cert = cfg.HTTP.X509Cert
			hs.FileDir = cfg.HTTP.FileDir
			hs.RateDur = cfg.HTTP.RateDur
			go hs.Serve()
		}
		// GRPC service
		{

		}
	}()
	<-exit
}
