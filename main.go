package main

import (
	"cards-against-humanity-api/accounts"
	"cards-against-humanity-api/server"
	"cards-against-humanity-api/sql"
	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

const httpAddr = "0.0.0.0:8080"

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	accountVerificationURL := getEnvOrPanic("CAH_VERIFICATION_URL")

	// SendGrid
	sendGridAPIToken := getEnvOrPanic("CAH_SENDGRID_API_TOKEN")
	mailClient := sendgrid.NewSendClient(sendGridAPIToken)

	// Connect to database
	sqlClient, err := sql.NewSQLClient(getEnvOrPanic("CAH_DATABASE_ADDRESS"))
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("cannot connect to SQL Client")
	}
	defer sqlClient.Close()

	databaseClient := sql.NewDatabaseClient(sqlClient, logger)

	accountClient := accounts.NewAccountClient(databaseClient, logger, mailClient, accountVerificationURL)

	// Start HTTP server
	srv := server.New(accountClient, logger)
	srv.ListenAndServe(httpAddr)
}
