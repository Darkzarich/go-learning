package config

import (
	"errors"
	"os"
)

type Config struct {
	Port       string
	StorageURL string
}

func Load() (Config, error) {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8081"
	}

	storageURL := os.Getenv("STORAGE_URL")
	if storageURL == "" {
		return Config{}, errors.New("STORAGE_URL is required")
	}

	return Config{StorageURL: storageURL, Port: port}, nil
}
