package server

import (
	"cards-against-humanity-api/accounts"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func setSessionExpirationTime() time.Time {
	return time.Now().UTC().Add(30 * time.Minute)
}

// postLoginHandler authenticates a user and creates a session
func (s *Server) postLoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, _ := s.sessionStore.Get(r, sessionToken)

	auth := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		s.logger.Error().Err(err).Msg("postLoginHandler: Error decoding json")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Authenticate User
	userID, err := s.accounts.AuthenticateUser(auth.Email, auth.Password)
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
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	// Set user as authenticated
	// Associate auth token to userID
	session.Values["authenticated"] = true
	session.Values["userID"] = userID

	err = session.Save(r, w)
	if err != nil {
		s.logger.Error().Err(err).Msg("postLoginHandler: Error saving session")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Println("session token:", w)
	fmt.Println("session values in Middleware:", session.Values)

	// Send session token
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// postLogoutHandler logs out a user from a session
func (s *Server) postLogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, _ := s.sessionStore.Get(r, sessionToken)

	// Invalidate session
	session.Values["autenticated"] = false

	// Save session
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	// Send back response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// UserAuthenticationRequired is authentication middleware for user requests
func (s *Server) UserAuthenticationRequired(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session, _ := s.sessionStore.Get(r, sessionToken)

		// TODO: handle cases that these values are nil
		authUserID := session.Values["userID"].(int)
		paramUserIDStr := ps.ByName("userID")

		// Convert userID param to int
		paramUserID, err := strconv.Atoi(paramUserIDStr)
		if err != nil {
			http.Error(w, "Forbidden: Authentication Failed", http.StatusForbidden)
			return
		}

		// If the userID in the param does not match the userID in session token
		if paramUserID != 0 && authUserID != paramUserID {
			http.Error(w, "Forbidden: Authentication Failed", http.StatusForbidden)
			return
		}

		// If expiration time has expired or expiration doesn't exist
		if isAuthenticated, ok := session.Values["authenticated"].(bool); !ok || !isAuthenticated {
			http.Error(w, "Forbidden: Authentication Failed", http.StatusForbidden)
			return
		}

		// Save session
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		h(w, r, ps)
	}
}
