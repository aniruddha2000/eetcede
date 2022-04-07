package api

import (
	"flag"

	"github.com/aniruddha2000/goEtcd/api/controllers"
)

var server controllers.Server

// Initialize and run the server
func Run() {
	var storageType string

	flag.StringVar(&storageType, "storage-type", "in-memory",
		"Define the storage type that will be used in the server. By defaut the value is in-memory.")
	flag.Parse()

	server.Initialize(storageType)
	server.Run("8080")
}
