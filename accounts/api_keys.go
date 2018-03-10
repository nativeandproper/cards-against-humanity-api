package accounts

import (
	"cards-against-humanity-api/models"
	"github.com/pkg/errors"
)

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
