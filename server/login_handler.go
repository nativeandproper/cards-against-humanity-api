package server

import (
	"cards-against-humanity-api/accounts"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// postLoginHandler authenticates a user and sends back a JWT token
func (s *Server) postLoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		http.Error(w, fmt.Sprintf("error decoding JSON: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if auth.Email == "" || auth.Password == "" {
		http.Error(w, "email and password cannot be blank", http.StatusBadRequest)
		return
	}

	// authenticate User
	user, err := s.accounts.AuthenticateUser(auth.Email, auth.Password)
	if err != nil {
		s.logger.Error().Err(err).Msg("postLoginHandler: Error authenticating user")
		switch err {
		case accounts.ErrAuthenticationInvalid:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case accounts.ErrUserNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, fmt.Sprintf("error authenticating user: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	signedToken, err := s.auth.Issue(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("error authenticating user: %s", err.Error()), http.StatusInternalServerError)
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", signedToken))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// postLogoutHandler logs out a user
func (s *Server) postLogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO: sign token with last logged in timestamp

	// Send back response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// UserAuthenticationRequired is authentication middleware for user requests
func (s *Server) UserAuthenticationRequired(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")

		isValid, err := s.auth.Validate(token)
		if err != nil {
			http.Error(w, "error authenticating", http.StatusInternalServerError)
			return
		}
		if !isValid {
			http.Error(w, "Forbidden: Authentication Failed", http.StatusForbidden)
			return
		}

		h(w, r, ps)
	}
}
