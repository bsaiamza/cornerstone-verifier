package server

import (
	"errors"
	"log"
	"net/http"
)

// Server serves HTTP requests for our services.
type Server struct {
	Router *http.Server
}

// NewServer creates a new HTTP server and set up routing.
func NewServer() *Server {
	server := &Server{
		Router: &http.Server{},
	}

	return server
}

// WithAddress returns the server address.
func (s *Server) WithAddress(addr string) *Server {
	s.Router.Addr = addr
	return s
}

// WithErrLogger returns server logging.
func (s *Server) WithErrLogger(l *log.Logger) *Server {
	s.Router.ErrorLog = l
	return s
}

// WithRouter returns server routing.
func (s *Server) WithRouter(router *http.ServeMux) *Server {
	s.Router.Handler = router
	return s
}

// Start starts up the server.
func (s *Server) Start() error {
	if len(s.Router.Addr) == 0 {
		return errors.New("Server missing address")
	}

	if s.Router.Handler == nil {
		return errors.New("Server missing handler")
	}

	return s.Router.ListenAndServe()
}

// Stop shuts down the server.
func (s *Server) Stop() error {
	return s.Router.Close()
}

// Res is a map type for http response data.
type Res map[string]interface{}
