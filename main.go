package main

import (
	"cards-against-humanity-api/accounts"
	"cards-against-humanity-api/auth"
	"cards-against-humanity-api/server"
	"cards-against-humanity-api/sql"
	"github.com/gorilla/sessions"
	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

const httpAddr = "0.0.0.0:9000"
const sessionExpiration = 86400 * 1

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// url for email verification
	emailVerificationURL := getEnvOrPanic("CAH_VERIFICATION_URL")

	// JWT auth secret
	jwtAuthSecret := []byte(getEnvOrPanic("CAH_AUTH_SECRET"))

	// sendGrid client
	sendGridAPIToken := getEnvOrPanic("CAH_SENDGRID_API_TOKEN")
	mailClient := sendgrid.NewSendClient(sendGridAPIToken)

	// create store and set session options
	sessionStore := sessions.NewCookieStore([]byte(getEnvOrPanic("CAH_AUTH_SECRET")))
	sessionStore.Options = &sessions.Options{
		MaxAge:   sessionExpiration,
		HttpOnly: true,
	}

	// connect to database
	sqlClient, err := sql.NewSQLClient(getEnvOrPanic("CAH_DATABASE_ADDRESS"))
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("cannot connect to SQL Client")
	}
	defer sqlClient.Close()

	databaseClient := sql.NewDatabaseClient(sqlClient, logger)
	accountClient := accounts.NewAccountClient(databaseClient, logger, mailClient, emailVerificationURL)
	authClient := auth.New(jwtAuthSecret)

	srv := server.New(accountClient, authClient, sessionStore, logger)
	srv.ListenAndServe(httpAddr)
}
