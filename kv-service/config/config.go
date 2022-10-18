package config

import (
	"log"
	"os"
)

type EnvironmentVariables struct {
	PORT string
}

const defaultPort string = "8080"

func GetEnvVars() *EnvironmentVariables {

	PORT, ok := os.LookupEnv("PORT")
	if !ok {
		PORT = defaultPort
		log.Printf("Port not specified, running on default port %v", PORT)

	}

	return &EnvironmentVariables{
		PORT: PORT,
	}
}
