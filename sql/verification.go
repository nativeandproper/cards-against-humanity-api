package sql

import (
	"database/sql"
	"github.com/pkg/errors"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"time"

	"github.com/nativeandproper/cards-against-humanity-api/models"
)

// InsertEmailVerification associates a verification token with a user
func (dc *DatabaseClient) InsertEmailVerification(userID int, token string, expiration time.Time) error {

	var emailVerification models.UserVerificationToken
	emailVerification.UserID = userID
	emailVerification.Token = token
	emailVerification.ExpiresAt = expiration

	// insert email verification
	err := emailVerification.Insert(dc.sqlClient)
	if err != nil {
		return errors.Wrap(err, "InsertEmailVerification: Error inserting row")
	}

	return nil
}

// GetEmailVerificationToken retrieves email verification by token
func (dc *DatabaseClient) GetEmailVerificationToken(token string) (*models.UserVerificationToken, error) {

	// get email verification by token
	emailVerification, err := models.UserVerificationTokens(dc.sqlClient, Select("id", "expires_at", "verified_at"), Where("token=?", token)).One()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "GetEmailVerificationToken: Error retrieving email verification")
	}

	return emailVerification, nil
}

// UpdateEmailVerification sets user as verified
func (dc *DatabaseClient) UpdateEmailVerification(ID int) error {

	// get email verification by ID
	emailVerification, err := models.FindUserVerificationToken(dc.sqlClient, ID)
	if err != nil {
		return errors.Wrap(err, "UpdateEmailVerification: Error retrieving email verification")
	}

	// set verifiedAt to current time
	emailVerification.VerifiedAt.Time = time.Now().UTC()
	emailVerification.VerifiedAt.Valid = true

	err = emailVerification.Update(dc.sqlClient)
	if err != nil {
		return errors.Wrap(err, "UpdateEmailVerification: Error updating email as verified")
	}

	return nil
}
