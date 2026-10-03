package config

import (
	"errors"
	"os"
)

type Config struct {
	KafkaBrokerURL string
	StorageURL     string
	Env            string
}

func Load() (Config, error) {
	kafkaBrokerURL := os.Getenv("KAFKA_BROKER")
	if kafkaBrokerURL == "" {
		return Config{}, errors.New("KAFKA_BROKER is required")
	}

	storageURL := os.Getenv("STORAGE_URL")
	if storageURL == "" {
		return Config{}, errors.New("STORAGE_URL is required")
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	return Config{KafkaBrokerURL: kafkaBrokerURL, StorageURL: storageURL, Env: env}, nil
}
