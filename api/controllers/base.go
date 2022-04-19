package controllers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/aniruddha2000/goEtcd/api/models"
)

// Server dependencies
type Server struct {
	Router *http.ServeMux
	Cache  models.Storage
}

// Initialize the routes
func (s *Server) Initialize(storageType string) {
	s.Router = http.NewServeMux()

	switch storageType {
	case "in-memory":
		s.Cache = models.NewCache()
	case "disk":
		s.Cache = models.NewDisk()
	default:
		log.Fatal("Use flags `in-memory` or `disk`")
	}

	log.Printf("Starting server with %v storage", storageType)

	s.initializeRoutes()
}

// Run the server on desired port and logs the status
func (s *Server) Run(addr string) {
	cert, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
	if err != nil {
		log.Fatalf("Couldn't load the certificate: %v", cert)
	}

	server := &http.Server{
		Addr:    ":" + addr,
		Handler: s.Router,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	fmt.Println("Listenning to port", addr)
	log.Fatal(server.ListenAndServeTLS("", ""))
}
