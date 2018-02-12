package accounts

import (
	"golang.org/x/crypto/bcrypt"
)

const hashCost = 14

// HashPassword hashes a password with salt
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return bytes, err
}

// CheckPasswordHash compares a hashed password with a plain-text password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
