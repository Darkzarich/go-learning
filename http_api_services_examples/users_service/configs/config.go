package config

import "os"

type Config struct {
	DBPath string // SQLite file
	Port   string
}

func Load() *Config {
	cfg := &Config{
		DBPath: "data.db",
		Port:   "3000",
	}

	if v := os.Getenv("DB_PATH"); v != "" {
		cfg.DBPath = v
	}
	if v := os.Getenv("PORT"); v != "" {
		cfg.Port = v
	}

	return cfg
}
