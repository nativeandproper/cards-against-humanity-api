package sql

import (
	"cards-against-humanity-api/models"
	"time"
)

// InsertUserVerification associates a verification token with a user
func (dc *DatabaseClient) InsertUserVerification(userID int, token string, expiration time.Time) error {
	var userVerification models.UserVerificationToken
	userVerification.UserID = userID
	userVerification.Token = token
	userVerification.ExpiresAt = expiration

	// Insert user verification
	err := userVerification.Insert(dc.sqlClient)
	if err != nil {
		dc.logger.Error().Err(err).Msg("CreateUserVerification: error inserting user verification")
		return err
	}

	return nil
}
