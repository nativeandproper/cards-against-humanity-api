package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/nativeandproper/cards-against-humanity-api/accounts"
	"github.com/nativeandproper/cards-against-humanity-api/auth"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

const (
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
)

// Server struct
type Server struct {
	accounts *accounts.AccountClient
	auth     *auth.AuthClient
	logger   zerolog.Logger
}

// New creates a new instance of Server
func New(accountClient *accounts.AccountClient, authClient *auth.AuthClient, logger zerolog.Logger) *Server {
	return &Server{
		accounts: accountClient,
		auth:     authClient,
		logger:   logger,
	}
}

// ListenAndServe creates a new http server instance
func (s *Server) ListenAndServe(httpAddr string) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
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
	router.GET("/status", statusHandler)

	router.PUT("/v1/signup", s.putSignupHandler)
	router.POST("/v1/signup", s.postSignupHandler)
	router.POST("/v1/login", s.postLoginHandler)
	router.POST("/v1/logout", s.UserAuthenticationRequired(s.postLogoutHandler))
	router.GET("/v1/auth", s.getAuthStatus)

	router.GET("/v1/user/:userID", s.UserAuthenticationRequired(s.getUser))
	router.GET("/v1/user/:userID/apikey", s.UserAuthenticationRequired(s.getAPIKeys))
	router.POST("/v1/user/:userID/apikey", s.UserAuthenticationRequired(s.postAPIKey))
	router.DELETE("/v1/user/:userID/apikey/:apiKeyID", s.UserAuthenticationRequired(s.deleteAPIKey))

	return router
}

// statusHandler handles requests to the /status endpoint
func statusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
