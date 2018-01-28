package server

import (
	"github.com/gorilla/mux"
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
type Server struct{}

// New creates a new instance of Server
func New() *Server {
	return &Server{}
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

// newRouter returns a mux with routes
func (s *Server) newRouter() *mux.Router {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/status", statusHandler)

	return router
}

// statusHandler handles requests to the status endpoint
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
