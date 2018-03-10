package server

import (
	"cards-against-humanity-api/accounts"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *Server) postAPIKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check login credentials for user matches the userID on the request
	userID := 4

	// Create and store the API key
	APIKey, err := s.accounts.CreateAPIKey(userID)
	if err != nil {
		s.logger.Error().Err(err).Msg("postAPIToken: Error creating API token")
		switch err {
		case accounts.ErrUserNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	// Return API Key
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(APIKey)
}
