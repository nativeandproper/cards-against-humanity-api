package accounts

import (
	"cards-against-humanity-api/models"
	"github.com/pkg/errors"
)

// DeactivateAPIKey soft deletes an API key for a user
func (a *AccountClient) DeactivateAPIKey(userID int, keyID int) error {
	// Soft delete API key
	err := a.databaseClient.DeleteAPIKey(userID, keyID)
	if err != nil {
		if err.Error() == "Not found" {
			return ErrTokenNotFound
		}
		return errors.Wrap(err, "DeactivateAPIKey: Error deactivating API key")
	}

	return nil
}

// CreateAPIKey creates an API key for a user
func (a *AccountClient) CreateAPIKey(userID int) (models.APIKey, error) {
	// Create API key
	token := generateBase64Token()

	// Insert API key
	APIKey, err := a.databaseClient.InsertAPIKey(userID, token)
	if err != nil {
		return APIKey, errors.Wrap(err, "CreateAPIKey: Error creating API key")
	}

	return APIKey, nil
}

// ListAPIKeys lists all API keys associated with a user
func (a *AccountClient) ListAPIKeys(userID int) (models.APIKeySlice, error) {
	// Retrieve API keys
	APIKeysList, err := a.databaseClient.GetAPIKeys(userID)
	if err != nil {
		return nil, errors.Wrap(err, "ListAPIKeys: Error getting list of API keys")
	}

	return APIKeysList, nil
}
