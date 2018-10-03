package accounts

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"unicode"
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

// ValidatePassword ensures a password meets minimum security requirements
func ValidatePassword(password string) bool {
	var containsNum, containsUpper, containsSpecial bool
	letters := 0

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			containsNum = true
		case unicode.IsUpper(c):
			containsUpper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			containsSpecial = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			return false
		}
	}
	validLen := letters >= 6 && letters <= 10
	return containsNum && containsUpper && containsSpecial && validLen
}

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

// AuthenticateUser checks if valid user
func (a *AccountClient) AuthenticateUser(email, password string) (int, error) {

	user, err := a.databaseClient.GetUserByEmail(email)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, ErrUserNotFound
	}

	// compare password
	validPassword := CheckPasswordHash(user.Password, []byte(password))
	if !validPassword {
		return user.ID, ErrAuthenticationInvalid
	}

	return user.ID, nil
}
