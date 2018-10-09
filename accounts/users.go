package accounts

import (
	"cards-against-humanity-api/models"
	"github.com/pkg/errors"
	"strings"
)

// CreateUser inserts a new user into the database
func (a *AccountClient) CreateUser(user *User) error {

	userExists, err := a.databaseClient.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return errors.Wrap(err, "CreateUser: Error checking if user exists")
	}

	// err email already taken
	if userExists {
		return ErrEmailMustBeUnique
	}

	// hash password
	hash, err := HashPassword(user.Password)
	if err != nil {
		return errors.Wrap(err, "CreateUser: Error could not hash password")
	}

	err = a.databaseClient.InsertUser(user.Email, strings.Title(user.FirstName), strings.Title(user.LastName), hash)
	if err != nil {
		return errors.Wrap(err, "CreateUser: Error inserting user")
	}

	return nil
}

// GetUser fetches a user by ID
func (a *AccountClient) GetUser(userID int) (*models.User, error) {
	user, err := a.databaseClient.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
