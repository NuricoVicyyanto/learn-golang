package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds runtime configuration for both binaries.
type Config struct {
	HTTPAddr        string
	DatabaseURL     string
	WorkerCount     int
	ShutdownTimeout time.Duration
}

// Load reads and validates configuration.
//
// TODO: return an error, reject invalid values, and redact secrets in logs.
func Load() Config {
	return Config{
		HTTPAddr:        env("HTTP_ADDR", ":8080"),
		DatabaseURL:     env("DATABASE_URL", ""),
		WorkerCount:     envInt("WORKER_COUNT", 4),
		ShutdownTimeout: 20 * time.Second,
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func envInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	n, err := strconv.Atoi(value)
	if err != nil || n <= 0 {
		return fallback
	}
	return n
}
