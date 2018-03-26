package sql

import (
	"cards-against-humanity-api/models"
	"database/sql"
	"github.com/pkg/errors"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

// DeleteAPIKey soft deletes an API Key associated with a user
func (dc *DatabaseClient) DeleteAPIKey(userID int, keyID int) error {

	// Verify API token is associated with user
	exists, err := models.UserAPIKeys(dc.sqlClient, Where("user_id=?", userID), Where("api_key_id=?", keyID)).Exists()
	if err != nil {
		return errors.Wrap(err, "DeleteAPIKey: error retrieving user API key")
	}

	// If user does not have an API token with keyID associated with account
	if !exists {
		return errors.New("Not found")
	}

	// Get API key
	APIKey, err := models.FindAPIKey(dc.sqlClient, keyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Not Found")
		}
		return errors.Wrap(err, "DeleteAPIKey: error retrieving API key")
	}

	// Set DeletedAt timestamp
	APIKey.DeletedAt.Time = time.Now().UTC()
	APIKey.DeletedAt.Valid = true

	// Insert API key and return
	err = APIKey.Update(dc.sqlClient)
	if err != nil {
		return errors.Wrap(err, "DeleteAPIKey: error deleting API Key")
	}

	return nil
}

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

// GetAPIKeys gets list of all API Keys associated with user
func (dc *DatabaseClient) GetAPIKeys(userID int) (models.APIKeySlice, error) {
	// Get API keys where the user_id matches userID
	APIKeys, err := models.APIKeys(dc.sqlClient,
		Select("*"),
		InnerJoin("user_api_keys uak on uak.api_key_id = api_keys.id"),
		Where("uak.user_id= ?", userID),
	).All()
	if err != nil {
		return nil, errors.Wrap(err, "GetAPIKeys: error inserting API keys")
	}

	return APIKeys, nil
}
