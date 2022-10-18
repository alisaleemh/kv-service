package service

import (
	"github.com/alisaleemh/kv-service/config"
	"github.com/alisaleemh/kv-service/storage"
)

type Service struct {
	Storage *storage.Storage
	EnvVars *config.EnvironmentVariables
}
