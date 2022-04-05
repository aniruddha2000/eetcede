package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aniruddha2000/goEtcd/api/models"
	"github.com/gorilla/mux"
)

// Server dependencies
type Server struct {
	Router *mux.Router
	Cache  models.Storage
}

// Initialize the routes
func (s *Server) Initialize(storageType string) {
	s.Router = mux.NewRouter()
	s.Cache = models.NewRecord()
	s.initializeRoutes()
}

// Run the server on desired port and logs the status
func (server *Server) Run(addr string) {
	fmt.Println("Listenning to port", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
