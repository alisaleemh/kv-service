package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvVars(t *testing.T) {

	defer cleanup()
	envVars := GetEnvVars()

	assert.Equal(t, envVars.PORT, defaultPort, "Port equals the default port")

	os.Setenv("PORT", "1")

	envVars = GetEnvVars()

	assert.Equal(t, envVars.PORT, "1", "Port equals the set port")

}

func cleanup() {
	os.Unsetenv("PORT")
}
