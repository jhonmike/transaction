package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	port   string = "9090"
	dbHost string = "mockhost"
	dbPort string = "5433"
	dbUser string = "maka"
	dbPass string = "123456"
	dbBase string = "bla"
)

func TestConfigMustReadFromEnvDefaults(t *testing.T) {
	cfg := MustReadFromEnv()

	assert.Equal(t, "8080", cfg.Port)
	assert.Equal(t, "localhost", cfg.DbHost)
	assert.Equal(t, "5432", cfg.DbPort)
	assert.Equal(t, "postgres", cfg.DbUser)
	assert.Equal(t, "postgres", cfg.DbPass)
	assert.Equal(t, "transaction", cfg.DbBase)
}

func TestConfigMustReadFromEnv(t *testing.T) {
	os.Setenv("PORT", port)
	os.Setenv("DB_HOST", dbHost)
	os.Setenv("DB_PORT", dbPort)
	os.Setenv("DB_USER", dbUser)
	os.Setenv("DB_PASS", dbPass)
	os.Setenv("DB_BASE", dbBase)
	cfg := MustReadFromEnv()

	assert.Equal(t, port, cfg.Port)
}
