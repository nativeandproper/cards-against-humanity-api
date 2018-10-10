package server

import (
	"cards-against-humanity-api/accounts"
	"cards-against-humanity-api/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"time"
)

type contextKey string

const ctxUser contextKey = "user"

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
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %v", signedToken))
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("access-control-expose-headers", "Authorization")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// postLogoutHandler logs out a user
func (s *Server) postLogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// get user from context
	userCTX := r.Context().Value("user")
	if userCTX == nil {
		http.Error(w, "error no user on context", http.StatusInternalServerError)
		return
	}

	// cast to user
	user, ok := userCTX.(*models.User)
	if !ok {
		http.Error(w, "error parsing user context", http.StatusInternalServerError)
		return
	}

	// update logout date
	user.LoggedOutAt.Time = time.Now().UTC()
	user.LoggedOutAt.Valid = true

	err := s.accounts.UpdateUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating user: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// UserAuthenticationRequired is authentication middleware for user requests
func (s *Server) UserAuthenticationRequired(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "error authentication failed", http.StatusForbidden)
			return
		}

		// trim token header
		token = strings.TrimPrefix(token, "Bearer ")

		// validate token
		isValid, claims, err := s.auth.Validate(token)
		if err != nil {
			s.logger.Error().Err(err).Msg("error authenticating token")
			http.Error(w, "error authenticating", http.StatusInternalServerError)
			return
		}
		if !isValid {
			http.Error(w, "error invalid authentication", http.StatusForbidden)
			return
		}

		userIDClaim, ok := claims["userID"]
		if !ok {
			http.Error(w, "error authentication failed: could not parseUserID", http.StatusForbidden)
			return
		}

		// parse userID to int
		userID, ok := userIDClaim.(int)
		if !ok {
			http.Error(w, "error authentication failed: could not cast userID to int", http.StatusInternalServerError)
			return
		}

		user, err := s.accounts.GetUser(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set user on context
		ctx := context.WithValue(r.Context(), ctxUser, user)
		h(w, r.WithContext(ctx), ps)
	}
}
