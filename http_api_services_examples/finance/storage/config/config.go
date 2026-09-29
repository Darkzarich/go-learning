package config

import (
	"errors"
	"os"
)

type Config struct {
	Addr        string
	DatabaseURL string
	Env         string
}

func Load() (Config, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = "8081"
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	return Config{Addr: addr, DatabaseURL: dsn, Env: env}, nil
}
