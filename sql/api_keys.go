package sql

import (
	"cards-against-humanity-api/models"
	"github.com/pkg/errors"
)

// InsertAPIKey inserts an API Key for a valid user
func (dc *DatabaseClient) InsertAPIKey(userID int, token string) (models.APIKey, error) {
	var UserAPIKey models.UserAPIKey
	var APIKey models.APIKey

	// Create transaction
	tx, err := dc.sqlClient.Begin()
	if err != nil {
		return APIKey, err
	}

	// Set API key values
	APIKey.APIKey = token

	// Insert API key and return
	err = APIKey.Insert(tx)
	if err != nil {
		tx.Rollback()
		return APIKey, errors.Wrap(err, "InsertAPIKey: error inserting API Key")
	}

	// Inser user API key values
	UserAPIKey.APIKeyID = APIKey.ID
	UserAPIKey.UserID = userID

	err = UserAPIKey.Insert(tx)
	if err != nil {
		tx.Rollback()
		return APIKey, errors.Wrap(err, "InsertAPIKey: error inserting User API Key")
	}

	tx.Commit()

	return APIKey, nil
}
