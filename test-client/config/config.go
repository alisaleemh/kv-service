package config

import (
	"log"
	"os"
)

type EnvironmentVariables struct {
	PORT           string
	KV_SERVICE_URL string
}

const defaultPort string = "8081"

func GetEnvVars() *EnvironmentVariables {

	PORT, ok := os.LookupEnv("PORT")
	if !ok {
		PORT = defaultPort
		log.Printf("Port not specified, running on default port %v", PORT)

	}

	KV_SERVICE_URL, ok := os.LookupEnv("PORT")
	if !ok {
		KV_SERVICE_URL = "http://localhost:8080"
		log.Printf("Port not specified, running on default port %v", PORT)

	}

	return &EnvironmentVariables{
		PORT:           PORT,
		KV_SERVICE_URL: KV_SERVICE_URL,
	}
}
