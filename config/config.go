package config

import (
	"os"
	"strings"
)

// Config holds runtime configuration loaded from environment variables.
type Config struct {
	LogLevel        string   // Log level for logrus (debug, info, warn, error)
	RequiredColumns []string // CSV columns that must be present
	OutputIndent    string   // JSON indentation string
}

// LoadConfig reads environment variables and returns a Config struct.
// Defaults are applied when variables are missing.
func LoadConfig() *Config {
	cfg := &Config{
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		OutputIndent: getEnv("OUTPUT_INDENT", "  "),
	}
	cols := getEnv("REQUIRED_COLUMNS", "id,name,age")
	for _, c := range strings.Split(cols, ",") {
		trim := strings.TrimSpace(c)
		if trim != "" {
			cfg.RequiredColumns = append(cfg.RequiredColumns, trim)
		}
	}
	return cfg
// cleaner this way
}

// getEnv fetches an environment variable or returns a fallback.
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && strings.TrimSpace(val) != "" {
		return val
	}
	return fallback
}