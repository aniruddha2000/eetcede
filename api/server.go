package api

import "github.com/aniruddha2000/goEtcd/api/controllers"

var server controllers.Server

// Initialize and run the server
func Run() {
	server.Initialize()
	server.Run(":8080")
}
