package main

import (
	"cards-against-humanity-api/server"
	"cards-against-humanity-api/sql"
	"cards-against-humanity-api/users"
)

const httpAddr = "0.0.0.0:8080"

func main() {

	// Connect to database
	sqlClient, err := sql.NewSQLClient(getEnvOrPanic("CAH_DATABASE_ADDRESS"))
	if err != nil {
		panic("cannot connect to SQL Client")
	}
	defer sqlClient.Close()

	userClient := users.NewUserClient(sqlClient)

	// Start HTTP server
	srv := server.New(userClient)
	srv.ListenAndServe(httpAddr)
}
