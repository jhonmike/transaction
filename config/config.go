package config

import (
	"github.com/caarlos0/env"
)

// Config Default system configs struct
type Config struct {
	Port   string `env:"PORT" envDefault:"8080"`
	DbHost string `env:"DB_HOST" envDefault:"localhost"`
	DbPort string `env:"DB_PORT" envDefault:"5432"`
	DbUser string `env:"DB_USER" envDefault:"postgres"`
	DbPass string `env:"DB_PASS" envDefault:"postgres"`
	DbBase string `env:"DB_BASE" envDefault:"transaction"`
}

// MustReadFromEnv Used to read env variables and return in a Config Struct
func MustReadFromEnv() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
