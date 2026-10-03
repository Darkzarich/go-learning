package config

import (
	"errors"
	"os"
)

type Config struct {
	KafkaBrokerUrl string
	StorageUrl     string
	Env            string
}

func Load() (Config, error) {
	KafkaBrokerUrl := os.Getenv("KAFKA_BROKER")
	if KafkaBrokerUrl == "" {
		return Config{}, errors.New("KAFKA_BROKER is required")
	}

	StorageUrl := os.Getenv("STORAGE_URL")
	if StorageUrl == "" {
		return Config{}, errors.New("STORAGE_URL is required")
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	return Config{KafkaBrokerUrl: KafkaBrokerUrl, StorageUrl: StorageUrl, Env: env}, nil
}
