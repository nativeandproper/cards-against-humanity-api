package main

import (
	"cards-against-humanity-api/server"
)

const httpAddr = "0.0.0.0:8080"

func main() {
	// Start HTTP server
	srv := server.New()
	srv.ListenAndServe(httpAddr)
}
