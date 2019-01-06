package sql

import (
	"database/sql"
	"time"

	"github.com/nativeandproper/cards-against-humanity-api/models"
	"github.com/pkg/errors"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

// DeleteAPIKey expires an API Key associated with a user
func (dc *DatabaseClient) DeleteAPIKey(userID int, userAPIKeyID int) (bool, error) {
	// verify API token is associated with user
	userAPIKey, err := models.FindUserAPIKey(dc.sqlClient, userAPIKeyID)
	if err != nil {
		return false, errors.Wrap(err, "DeleteAPIKey: Error retrieving user API key")
	}
	if userAPIKey == nil {
		return false, nil
	}

	apiKey, err := models.FindAPIKey(dc.sqlClient, userAPIKeyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.Wrap(err, "DeleteAPIKey: Error retrieving API key")
	}

	// already deactivated
	if apiKey.DeletedAt.Valid {
		return true, nil
	}

	// set DeletedAt timestamp
	apiKey.DeletedAt.Time = time.Now().UTC()
	apiKey.DeletedAt.Valid = true

	// update expiration
	err = apiKey.Update(dc.sqlClient)
	if err != nil {
		return false, errors.Wrap(err, "DeleteAPIKey: Error deleting API Key")
	}

	return true, nil
}

// InsertAPIKey inserts an API Key for a valid user
func (dc *DatabaseClient) InsertAPIKey(userID int, token string) (*models.APIKey, error) {
	var UserAPIKey models.UserAPIKey
	var apiKey models.APIKey

	tx, err := dc.sqlClient.Begin()
	if err != nil {
		return nil, err
	}

	// set API token
	apiKey.APIKey = token

	// insert API key
	err = apiKey.Insert(tx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			dc.logger.Error().Err(err).Msg("InsertAPIKey: Error rolling back transaction")
		}
		return nil, errors.Wrap(err, "InsertAPIKey: Error inserting API Key")
	}

	// insert user API key values
	UserAPIKey.APIKeyID = apiKey.ID
	UserAPIKey.UserID = userID

	err = UserAPIKey.Insert(tx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			dc.logger.Error().Err(err).Msg("InsertAPIKey: Error rolling back transaction")
		}
		return nil, errors.Wrap(err, "InsertAPIKey: Error inserting User API Key")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "InsertAPIKey: Error committing transaction")
	}

	return &apiKey, nil
}

// GetAPIKeys gets list of all API Keys associated with user
func (dc *DatabaseClient) GetAPIKeys(userID int) (models.APIKeySlice, error) {
	return models.APIKeys(dc.sqlClient,
		Select("api_keys.id, api_keys.api_key, api_keys.created_at, api_keys.deleted_at"),
		InnerJoin("user_api_keys uak on uak.api_key_id = api_keys.id"),
		Where("uak.user_id= ?", userID),
	).All()
}
