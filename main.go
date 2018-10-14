package main

import (
	"cards-against-humanity-api/accounts"
	"cards-against-humanity-api/auth"
	"cards-against-humanity-api/ratelimiter"
	"cards-against-humanity-api/server"
	"cards-against-humanity-api/sql"
	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go"
	"os"
	"time"
)

const httpAddr = "0.0.0.0:9000"

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// url for email verification
	emailVerificationURL := getEnvOrPanic("CAH_VERIFICATION_URL")

	// JWT auth secret
	jwtAuthSecret := []byte(getEnvOrPanic("CAH_AUTH_SECRET"))

	// sendGrid client
	sendGridAPIToken := getEnvOrPanic("CAH_SENDGRID_API_TOKEN")
	mailClient := sendgrid.NewSendClient(sendGridAPIToken)

	// connect to database
	sqlClient, err := sql.NewSQLClient(getEnvOrPanic("CAH_DATABASE_ADDRESS"))
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("cannot connect to SQL Client")
	}
	defer sqlClient.Close()

	redisAddr := getEnvOrPanic("CAH_REDIS_ADDR")
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	rateLimiterClient := ratelimiter.New(redisClient, &ratelimiter.LimiterOptions{
		Namespace:     "rate-limit-",
		SlidingWindow: 5 * time.Minute,
		Limit:         int64(5),
	})

	databaseClient := sql.NewDatabaseClient(sqlClient, logger)
	accountClient := accounts.NewAccountClient(databaseClient, logger, mailClient, emailVerificationURL)
	authClient := auth.New(jwtAuthSecret)

	srv := server.New(accountClient, authClient, rateLimiterClient, logger)
	srv.ListenAndServe(httpAddr)
}
