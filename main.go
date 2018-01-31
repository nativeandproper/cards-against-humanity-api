package main

import (
	"cards-against-humanity-api/server"
	"cards-against-humanity-api/sql"
	"cards-against-humanity-api/users"
	"github.com/rs/zerolog"
	"os"
)

const httpAddr = "0.0.0.0:8080"

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Connect to database
	sqlClient, err := sql.NewSQLClient(getEnvOrPanic("CAH_DATABASE_ADDRESS"))
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("cannot connect to SQL Client")
	}
	defer sqlClient.Close()

	userClient := users.NewUserClient(sqlClient, logger)

	// Start HTTP server
	srv := server.New(userClient, logger)
	srv.ListenAndServe(httpAddr)
}
