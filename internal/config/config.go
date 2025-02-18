package config

import (
	"os"
)

type Config struct {
	Port        string
	StorageType string
	PostgresDSN string
}

func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		StorageType: getEnv("STORAGE_TYPE", "memory"),
		PostgresDSN: getEnv("POSTGRES_DSN", "postgres://user:password@localhost:5432/shortener?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
