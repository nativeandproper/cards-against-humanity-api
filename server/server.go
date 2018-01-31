package server

import (
	users "cards-against-humanity-api/users"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
	"time"
)

const (
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
)

// Server struct
type Server struct {
	User   *users.UserClient
	Logger zerolog.Logger
}

// New creates a new instance of Server w
func New(userClient *users.UserClient, logger zerolog.Logger) *Server {
	return &Server{
		User:   userClient,
		Logger: logger,
	}
}

// ListenAndServe creates a new http server instance
func (s *Server) ListenAndServe(httpAddr string) {
	srv := &http.Server{
		Addr:         httpAddr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      s.newRouter(),
	}

	s.Logger.Info().Msgf("Listening on port %s", strings.Split(httpAddr, ":")[1])

	err := srv.ListenAndServe()
	if err != nil {
		s.Logger.Error().Err(err)
	}
}

// newRouter returns an http router with routes
func (s *Server) newRouter() *httprouter.Router {
	router := httprouter.New()

	// Routes
	router.GET("/status", statusHandler)
	router.POST("/v1/signup", s.postSignupHandler)

	return router
}

// statusHandler handles requests to the /status endpoint
func statusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
