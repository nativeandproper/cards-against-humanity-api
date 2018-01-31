package server

import (
	users "cards-against-humanity-api/users"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

// postSignupHandler handles requests to the /signup endpoint
func (s *Server) postSignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := &users.User{}

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Hash user password
	hash, err := HashPassword(user.Password)
	if err != nil {
		s.Logger.Error().Err(err).Msg("Could not hash password")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Insert new user
	err = s.User.CreateUser(user, hash)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
