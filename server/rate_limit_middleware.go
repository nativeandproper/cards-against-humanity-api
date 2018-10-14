package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

// TODO: move this to the ratelimiter service
var dailyMaxLimit = 200

func (s *Server) RateLimit(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// API key header
		headerKey := r.Header.Get("Authorization")

		apiKey, err := s.accounts.GetAPIKey(headerKey)
		if err != nil {
			http.Error(w, "error API key not found", http.StatusNotFound)
			return
		}

		// TODO: check cache if daily rate limit is exceeded in cache
		// http.Error(w, fmt.Sprintf("daily max rate of %d requests exceeded", dailyMaxLimit), http.StatusTooManyRequests)

		remaining, err := s.rateLimiter.Enforce(apiKey.APIKey)
		if err != nil {
			// log error, but allow request to complete
			s.logger.Error().Err(err).Msg(fmt.Sprintf("error enforcing rate limit for api key: [%s]", apiKey.APIKey))
		}

		w.Header().Set("X-Rate-Limit-Duration", fmt.Sprintf("%d", int64(s.rateLimiter.SlidingWindow.Seconds())))
		w.Header().Set("X-Rate-Limit-Remaining", fmt.Sprintf("%d", remaining))

		if err != nil || remaining <= 0 {
			retryAt := time.Now().UTC().Add(s.rateLimiter.SlidingWindow).Unix()
			w.Header().Set("X-Rate-Limit-Retry", fmt.Sprintf("%d", retryAt))
			http.Error(w, "requests exceeded for current interval", http.StatusTooManyRequests)
			return
		}
	}
}
