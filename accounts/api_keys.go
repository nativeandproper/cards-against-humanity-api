package accounts

import (
	"cards-against-humanity-api/models"
	"fmt"
	"github.com/pkg/errors"
)

const maxAPIKeyLimit = 3

// DeactivateAPIKey soft deletes an API key for a user
func (a *AccountClient) DeactivateAPIKey(userID int, keyID int) error {
	keyExists, err := a.databaseClient.DeleteAPIKey(userID, keyID)
	if err != nil {
		return errors.Wrap(err, "DeactivateAPIKey: Error deactivating API key")
	}
	if !keyExists {
		return ErrTokenNotFound
	}

	return nil
}

// CreateAPIKey generates an API key for a user
func (a *AccountClient) CreateAPIKey(userID int) (*models.APIKey, error) {
	apiKeysList, err := a.databaseClient.GetAPIKeys(userID)
	if err != nil {
		return nil, errors.Wrap(err, "CreateAPIKeys: Error getting list of API keys")
	}

	var activeKeysCount int
	for _, key := range apiKeysList {
		// api key is active (hasn't been deleted)
		if !key.DeletedAt.Valid {
			activeKeysCount++
		}
	}

	if activeKeysCount >= maxAPIKeyLimit {
		return nil, fmt.Errorf("max active API key limit of %d reached", maxAPIKeyLimit)
	}

	// generate new api key
	token := generateBase64Token()

	// insert API key
	apiKey, err := a.databaseClient.InsertAPIKey(userID, token)
	if err != nil {
		return nil, errors.Wrap(err, "CreateAPIKey: Error creating API key")
	}

	return apiKey, nil
}

// ListAPIKeys lists all API keys associated with a user
func (a *AccountClient) ListAPIKeys(userID int) (models.APIKeySlice, error) {
	apiKeys, err := a.databaseClient.GetAPIKeys(userID)
	if err != nil {
		return nil, errors.Wrap(err, "ListAPIKeys: Error getting list of user API keys")
	}

	return apiKeys, nil
}
