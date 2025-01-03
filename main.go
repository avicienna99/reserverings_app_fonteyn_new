package main

import (
	"github.com/avicienna99/reserverings_app_fonteyn_new/db"
	server "github.com/avicienna99/reserverings_app_fonteyn_new/webserver"
)

func main() {
	// Initialize database connection
	db.Init()
	defer db.Close()

	// Start the HTTP server
	server.Start()
}
