package config

import (
	"os"
	"time"
)

// Config holds runtime configuration.
type Config struct {
	Addr            string
	ShutdownTimeout time.Duration
	DatabaseURL     string
}

// Load reads configuration from the environment.
//
// TODO: parse and validate every value; fail fast on invalid input.
func Load() Config {
	return Config{
		Addr:            env("ADDR", ":8080"),
		ShutdownTimeout: 10 * time.Second,
		DatabaseURL:     env("DATABASE_URL", ""),
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
