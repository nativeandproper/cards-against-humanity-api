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
	session, err := s.sessionStore.Get(r, sessionToken)
	if err != nil {
		s.logger.Error().Err(err).Msg("postLogoutHandler: Error retrieving authentication session")
		http.Error(w, "error authenticating", http.StatusInternalServerError)
		return
	}

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
			http.Error(w, fmt.Sprintf("error authenticating user: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	// set auth values on session
	session.Values["authenticated"] = true
	session.Values["userID"] = userID

	err = session.Save(r, w)
	if err != nil {
		s.logger.Error().Err(err).Msg("postLoginHandler: Error saving session")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Println(w, "Is Login session new? %b", session.IsNew)
	fmt.Println(w, "Cookie set to after login %v", session.Values)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// postLogoutHandler logs out a user from a session
func (s *Server) postLogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, err := s.sessionStore.Get(r, sessionToken)
	if err != nil {
		s.logger.Error().Err(err).Msg("postLogoutHandler: Error retrieving authentication session")
		http.Error(w, "error authenticating", http.StatusInternalServerError)
		return
	}

	fmt.Println(w, "Is Logout session new? %b", session.IsNew)

	// expire session
	session.Options.MaxAge = -1
	session.Values["authenticated"] = false

	// Save session
	err = session.Save(r, w)
	if err != nil {
		s.logger.Error().Err(err).Msg("postLogoutHandler: Error saving session")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Println(w, "Cookie set to %v on LOGOUT", session.Values)

	// Send back response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// UserAuthenticationRequired is authentication middleware for user requests
func (s *Server) UserAuthenticationRequired(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session, err := s.sessionStore.Get(r, sessionToken)
		if err != nil {
			s.logger.Error().Err(err).Msg("UserAuthenticationRequired: Error retrieving authentication session")
			http.Error(w, "error authenticating", http.StatusInternalServerError)
			return
		}

		fmt.Println(w, "Is middleware session new? %b", session.IsNew)

		fmt.Println("session values in middleware before:", session.Values)

		sessionUserID := session.Values["userID"]
		if sessionUserID == nil {
			http.Error(w, "Forbidden: Authentication Failed", http.StatusForbidden)
			return
		}

		paramUserID, err := strconv.Atoi(ps.ByName("userID"))
		if err != nil {
			http.Error(w, "Forbidden: Authentication Failed", http.StatusForbidden)
			return
		}

		// userID param does not match userID in session token
		if paramUserID != 0 && sessionUserID.(int) != paramUserID {
			http.Error(w, "Forbidden: Authentication Failed", http.StatusForbidden)
			return
		}

		// expiration time has expired or doesn't exist
		if isAuthenticated, ok := session.Values["authenticated"].(bool); !ok || !isAuthenticated {
			http.Error(w, "Forbidden: Authentication Failed", http.StatusForbidden)
			return
		}

		fmt.Println("session values in middleware after:", session.Values)

		// Save session
		err = session.Save(r, w)
		if err != nil {
			s.logger.Error().Err(err).Msg("postLogoutHandler: Error saving session in middleware")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		h(w, r, ps)
	}
}
