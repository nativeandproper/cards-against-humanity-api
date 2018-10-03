package server

import (
	"cards-against-humanity-api/accounts"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// postSignupHandler handles requests for user signup
func (s *Server) postSignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newUser := &accounts.User{}

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(newUser); err != nil {
		http.Error(w, fmt.Sprintf("error: could not decode JSON payload %s", err.Error()), http.StatusBadRequest)
		return
	}

	if newUser.FirstName == "" || newUser.LastName == "" {
		http.Error(w, fmt.Sprintf("error: name fields cannot be blank"), http.StatusBadRequest)
		return
	}

	min, max := 3, 10
	if len(newUser.FirstName) < min || len(newUser.FirstName) > max || len(newUser.LastName) < min || len(newUser.LastName) > max {
		http.Error(w, fmt.Sprintf("error: name fields must be in range of %d to %d characters", min, max), http.StatusBadRequest)
		return
	}

	if newUser.Email == "" {
		http.Error(w, fmt.Sprintf("error: email cannot be empty"), http.StatusBadRequest)
		return
	}

	if !accounts.ValidatePassword(newUser.Password) {
		http.Error(w, fmt.Sprintf("error: password must be between 6-10 characters with one uppercase letter, one number and one special character"), http.StatusBadRequest)
		return
	}

	// Create new user
	err := s.accounts.CreateUser(newUser)
	if err != nil {
		switch err {
		case accounts.ErrEmailMustBeUnique:
			http.Error(w, err.Error(), http.StatusConflict)
			return
		default:
			http.Error(w, fmt.Sprintf("error creating account: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	// Create and send user verification link to user email
	user, err := s.accounts.CreateEmailVerification(newUser.Email)
	if err != nil {
		s.logger.Error().Err(err).Msg("postSignupHandler: Error user verification failed")

		switch err {
		case accounts.ErrEmailVerificationNotDeliverable:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, fmt.Sprintf("error creating email verification: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// putSignupHandler verifies email address associated with account via token
func (s *Server) putSignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payload := &struct {
		Token string `json:"verification_token"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("error decoding JSON payload %s", err.Error()), http.StatusBadRequest)
		return
	}

	if payload.Token == "" {
		http.Error(w, fmt.Sprintf("email verification token can't be empty"), http.StatusBadRequest)
		return
	}

	err := s.accounts.UpdateUserVerifyEmail(payload.Token)
	if err != nil {
		s.logger.Error().Err(err).Msg("putSignupHandler: Error user verification failed")
		switch err {
		case accounts.ErrEmailVerificationTokenExpired:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case accounts.ErrTokenNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
