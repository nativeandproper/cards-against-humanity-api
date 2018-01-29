package server

import (
	users "cards-against-humanity-api/users"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

const (
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
)

// Server struct
type Server struct {
	User *users.UserClient
}

// New creates a new instance of Server w
func New(userClient *users.UserClient) *Server {
	return &Server{
		userClient,
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

	log.Println("Listening ...")

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
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
