package main

import (
	"github.com/alisaleemh/kv-service/api"
	"github.com/alisaleemh/kv-service/config"
	"github.com/alisaleemh/kv-service/service"
	"github.com/alisaleemh/kv-service/storage"
	"github.com/gorilla/mux"
)

func main() {

	// Get the env vars
	envVars := config.GetEnvVars()

	service := service.Service{
		EnvVars: envVars,
		Storage: storage.NewStorage(envVars),
	}

	startHttpServer(&service, envVars)

}

func startHttpServer(service api.Service, envVars *config.EnvironmentVariables) {

	router := mux.NewRouter()
	server := api.NewServer(router, service, envVars)
	api.Initialize(server)

}
