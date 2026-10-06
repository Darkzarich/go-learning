package config

import (
	"errors"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	Env         string
}

func Load() (Config, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	return Config{Port: port, DatabaseURL: dsn, Env: env}, nil
}
