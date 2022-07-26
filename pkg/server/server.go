package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router http.Server
}

func NewServer() *Server {
	return &Server{
		Router: http.Server{},
	}
}

func (s *Server) WithAddress(address string) *Server {
	s.Router.Addr = address
	return s
}

func (s *Server) WithErrorLogger(logger *log.Logger) *Server {
	s.Router.ErrorLog = logger
	return s
}

func (s *Server) WithRouter(router *mux.Router) *Server {
	s.Router.Handler = router
	return s
}

func (s *Server) Start() error {
	if len(s.Router.Addr) == 0 {
		return errors.New("Server missing address")
	}

	if s.Router.Handler == nil {
		return errors.New("Server missing handler")
	}

	return s.Router.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.Router.Close()
}

type Response map[string]interface{}
