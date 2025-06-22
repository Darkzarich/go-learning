package config

import (
	"encoding/json"
	"log"
	"os"
	"server_load_balancer/types"
)

type Config struct {
	Servers []*types.Server `json:"servers"`
}

func ReadConfig() *Config {
	// read from config.json
	file, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var config Config

	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		log.Fatal(err)
	}

	return &config
}
