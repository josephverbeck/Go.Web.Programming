package main

import (
	"net/http"
)

// The Server Stuct type!
/*type Server struct {
    Addr           string
    Handler        Handler
    ReadTimeout    time.Duration
    WriteTimeout   time.Duration
	MaxHeaderBytes int
	TLSConfig
	TLSNextProto
	ConnState
	ErrorLog
}*/

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	// server.ListenAndServe() //http
	server.ListenAndServeTLS("cert.pem", "key.pem") //https
}
