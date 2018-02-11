package server

import (
	"cards-against-humanity-api/accounts"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// postSignupHandler handles requests to the /signup endpoint
func (s *Server) postSignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newUser := &accounts.User{}

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(newUser); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Create new user
	err := s.accounts.CreateUser(newUser)
	if err != nil {
		s.logger.Error().Err(err).Msg("postSignupHandler: Error create user failed")
		switch err {
		case accounts.ErrEmailMustBeUnique:
			http.Error(w, err.Error(), http.StatusConflict)
			return
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	// Create and send account verification link to user email
	user, err := s.accounts.CreateUserVerification(newUser.Email)
	if err != nil {
		s.logger.Error().Err(err).Msg("postSignupHandler: Error user verification failed")

		switch err {
		case accounts.ErrEmailVerificationNotDeliverable:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	// Send back user information
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// putSignupHandler handles put requests to the /signup endpoint
func (s *Server) putSignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payload := &struct {
		Token string `json:"token"`
	}{}

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Verify token
	err := s.accounts.UpdateUserVerification(payload.Token)
	if err != nil {
		s.logger.Error().Err(err).Msg("putSignupHandler: Error user verification failed")
		switch err {
		case accounts.ErrUserVerificationTokenHasExpired:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case accounts.ErrUserNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ok"))
}
