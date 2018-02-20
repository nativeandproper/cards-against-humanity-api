package accounts

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const hashCost = 14

// HashPassword hashes a password with salt
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return bytes, err
}

// CheckPasswordHash compares a hashed password with a plain-text password
func CheckPasswordHash(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

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

// AuthenticateUser checks if user is valid
func (a *AccountClient) AuthenticateUser(email, password string) error {

	// Get user by email
	user, err := a.databaseClient.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "Not Found" {
			return ErrUserNotFound
		}
		return err
	}

	// Compare password
	validPassword := CheckPasswordHash(user.Password, []byte(password))

	if !validPassword {
		return ErrAuthenticationInvalid
	}

	return nil
}
