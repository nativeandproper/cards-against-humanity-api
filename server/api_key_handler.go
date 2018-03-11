package server

import (
	"cards-against-humanity-api/accounts"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (s *Server) postAPIKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Ensure userID param exists
	userIDStr := ps.ByName("userID")
	if userIDStr == "" {
		http.Error(w, "Forbidden: missing expected param", http.StatusForbidden)
	}

	// Parse userID to int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		s.logger.Error().Err(err).Msg("postAPIKey: Error parsing userID to int")
		http.Error(w, "Forbidden: malformed param", http.StatusForbidden)
	}

	// Create and store the API key
	APIKey, err := s.accounts.CreateAPIKey(userID)
	if err != nil {
		s.logger.Error().Err(err).Msg("postAPIKey: Error creating API key")
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

func (s *Server) deleteAPIKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Ensure userID param exists
	userIDStr := ps.ByName("userID")
	if userIDStr == "" {
		http.Error(w, "Forbidden: missing expected param", http.StatusForbidden)
	}

	// Parse userID to int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		s.logger.Error().Err(err).Msg("postAPIKey: Error parsing userID to int")
		http.Error(w, "Forbidden: malformed param", http.StatusForbidden)
	}

	// Ensure APIKey param exists
	APIKeyStr := ps.ByName("apiKey")
	if APIKeyStr == "" {
		http.Error(w, "Forbidden: missing expected param", http.StatusForbidden)
	}

	// Parse APIKey to int
	APIKey, err := strconv.Atoi(APIKeyStr)
	if err != nil {
		s.logger.Error().Err(err).Msg("postAPIKey: Error parsing userID to int")
		http.Error(w, "Forbidden: malformed param", http.StatusForbidden)
	}

	// Delete API key
	err = s.accounts.DeactivateAPIKey(userID, APIKey)
	if err != nil {
		s.logger.Error().Err(err).Msg("deleteAPIKey: Error invalidating API key")
		switch err {
		case accounts.ErrTokenNotFound:
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
