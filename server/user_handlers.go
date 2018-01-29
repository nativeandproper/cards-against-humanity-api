package server

import (
	users "cards-against-humanity-api/users"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// signupHandler handles requests to the /signup endpoint
func (s *Server) postSignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := &users.User{}

	// decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
