package accounts

import (
	"github.com/pkg/errors"
	"strings"
)

// CreateUser inserts a new user into the database
func (a *AccountClient) CreateUser(user *User) error {

	// Check if user with that email already exists
	userExists, err := a.databaseClient.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return errors.Wrap(err, "CreateUser: Error checking if user exists")
	}

	// Do not create allow user to be created if email is already taken
	if userExists {
		return ErrEmailMustBeUnique
	}

	// Hash password
	hash, err := HashPassword(user.Password)
	if err != nil {
		return errors.Wrap(err, "CreateUser: Error could not hash password")
	}

	// Capitalize first letter of names
	firstName := strings.Title(user.FirstName)
	lastName := strings.Title(user.LastName)

	// Insert new user
	err = a.databaseClient.InsertUser(user.Email, firstName, lastName, hash)
	if err != nil {
		return errors.Wrap(err, "CreateUser: Error inserting user")
	}

	return nil
}
