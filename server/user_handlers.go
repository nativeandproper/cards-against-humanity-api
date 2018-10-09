package server

import (
	"cards-against-humanity-api/accounts"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

//func (s *Server) postLoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
func (s *Server) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userIDStr := ps.ByName("userID")

	// parse userID to int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		s.logger.Error().Err(err).Msg("getUser: Error parsing userID to int")
		http.Error(w, "Forbidden: malformed param", http.StatusForbidden)
		return
	}

	// get user
	user, err := s.accounts.GetUser(userID)
	if err != nil {
		switch err {
		case accounts.ErrUserNotFound:
			http.Error(w, accounts.ErrUserNotFound.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	// remove password
	user.Password = []byte("")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}
