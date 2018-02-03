package server

import (
	"cards-against-humanity-api/accounts"
	"encoding/json"
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
