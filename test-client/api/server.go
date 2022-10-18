package api

import (
	"net/http"

	"github.com/alisaleemh/test-client/config"
	"github.com/gorilla/mux"
)

type Service interface {
	TestOverwrite() error
	TestDelete() error
}

type Server struct {
	Service Service
	Router  *mux.Router
	EnvVars *config.EnvironmentVariables
}

func Initialize(server *Server) {
	http.ListenAndServe(":"+server.EnvVars.PORT, server.Router)
}

func NewServer(router *mux.Router, service Service, envVars *config.EnvironmentVariables) *Server {

	s := &Server{
		Service: service,
		Router:  router,
		EnvVars: envVars,
	}

	s.routes()

	return s
}
