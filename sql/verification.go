package sql

import (
	"cards-against-humanity-api/models"
	"database/sql"
	"github.com/pkg/errors"
	. "github.com/volatiletech/sqlboiler/queries/qm"
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
		return errors.Wrap(err, "InsertUserVerification: Error inserting row for email verification")
	}

	return nil
}

// GetUserVerificationByToken retrieves user verification by token
func (dc *DatabaseClient) GetUserVerificationByToken(token string) (*models.UserVerificationToken, error) {

	// Get user verification by token
	userVerification, err := models.UserVerificationTokens(dc.sqlClient, Select("id", "expires_at", "verified_at"), Where("token=?", token)).One()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "GetUserVerificationByToken: Error retrieving row for user verification")
	}

	return userVerification, nil
}

// UpdateUserVerification sets a user as verified
func (dc *DatabaseClient) UpdateUserVerification(ID int) error {

	// Get user verification by ID
	userVerification, err := models.FindUserVerificationToken(dc.sqlClient, ID)
	if err != nil {
		return errors.Wrap(err, "UpdateUserVerification: Error retrieving user")
	}

	// Set verifiedAt column to current time
	userVerification.VerifiedAt.Time = time.Now().UTC()
	userVerification.VerifiedAt.Valid = true

	// Update user verification
	err = userVerification.Update(dc.sqlClient)
	if err != nil {
		return errors.Wrap(err, "UpdateUserVerification: Error updating user as verified")
	}

	return nil
}
