package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var dailyMaxLimit = 200

func (s *Server) RateLimit(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// API key header
		headerKey := r.Header.Get("Authorization")

		if len(headerKey) != 64 {
			http.Error(w, "error forbidden: invalid API key", http.StatusForbidden)
			return
		}

		apiKey, err := s.accounts.GetAPIKey(headerKey)
		if err != nil {
			http.Error(w, "error API key not found", http.StatusNotFound)
			return
		}

		fmt.Println("api key:", apiKey)

		rateExceeded := true
		// TODO: check cache if rate limit is exceeded

		if rateExceeded {
			w.Header().Set("X-Rate-Limit-Retry", "12345")
			http.Error(w, fmt.Sprintf("daily max rate of %d requests exceeded", dailyMaxLimit), http.StatusTooManyRequests)
			return
		}

		// rate limit

		// add req limit headers to response
		w.Header().Set("X-Rate-Limit-Duration", "12345")
		w.Header().Set("X-Rate-Limit", "12345")
	}
}
