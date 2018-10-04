package server

import (
	"cards-against-humanity-api/accounts"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
	"time"
)

const (
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
	sessionToken = "cah-session-token"
)

// Server struct
type Server struct {
	accounts     *accounts.AccountClient
	logger       zerolog.Logger
	sessionStore *sessions.CookieStore
}

// New creates a new instance of Server
func New(accountClient *accounts.AccountClient, sessionStore *sessions.CookieStore, logger zerolog.Logger) *Server {
	return &Server{
		accounts:     accountClient,
		sessionStore: sessionStore,
		logger:       logger,
	}
}

// ListenAndServe creates a new http server instance
func (s *Server) ListenAndServe(httpAddr string) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Addr:         httpAddr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      c.Handler(context.ClearHandler(s.newRouter())),
	}

	s.logger.Info().Msgf("Listening on port %s", strings.Split(httpAddr, ":")[1])

	err := srv.ListenAndServe()
	if err != nil {
		s.logger.Error().Err(err)
	}
}

// newRouter returns an http router with routes
func (s *Server) newRouter() *httprouter.Router {
	router := httprouter.New()

	// Routes
	router.PUT("/v1/signup", s.putSignupHandler)
	router.GET("/status", statusHandler)
	router.POST("/v1/signup", s.postSignupHandler)
	router.POST("/v1/login", s.postLoginHandler)
	router.POST("/v1/logout", s.UserAuthenticationRequired(s.postLogoutHandler))

	router.GET("/v1/user/:userID/apikey", s.UserAuthenticationRequired(s.getAPIKeys))
	router.POST("/v1/user/:userID/apikey", s.UserAuthenticationRequired(s.postAPIKey))
	router.DELETE("/v1/user/:userID/apikey/:apiKey", s.UserAuthenticationRequired(s.deleteAPIKey))

	return router
}

// statusHandler handles requests to the /status endpoint
func statusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
