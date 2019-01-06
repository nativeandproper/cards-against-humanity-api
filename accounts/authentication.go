package accounts

import (
	"golang.org/x/crypto/bcrypt"
	"unicode"

	"github.com/nativeandproper/cards-against-humanity-api/models"
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

// AuthenticateUser checks if valid user
func (a *AccountClient) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := a.databaseClient.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	// compare password
	validPassword := CheckPasswordHash(user.Password, []byte(password))
	if !validPassword {
		return user, ErrAuthenticationInvalid
	}

	return user, nil
}
