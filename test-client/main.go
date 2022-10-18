package main

import (
	"github.com/alisaleemh/test-client/api"
	"github.com/alisaleemh/test-client/config"
	"github.com/alisaleemh/test-client/service"
	"github.com/gorilla/mux"
)

func main() {

	// Get the env vars
	envVars := config.GetEnvVars()

	service := service.Service{
		EnvVars: envVars,
	}

	startHttpServer(&service, envVars)

}

func startHttpServer(service api.Service, envVars *config.EnvironmentVariables) {

	router := mux.NewRouter()
	server := api.NewServer(router, service, envVars)
	api.Initialize(server)

}
