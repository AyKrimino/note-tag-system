// Package config contains the configuration for the application
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DatabaseConfig contains the configuration for the database
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Config contains the configuration for the application
type Config struct {
	Database DatabaseConfig
}

// Load returns the configuration
func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbCfg := DatabaseConfig{
		Host:     getenv("DB_HOST", "localhost"),
		Port:     getenv("DB_PORT", "5432"),
		User:     getenv("DB_USER", "_note_user"),
		Password: getenv("DB_PASSWORD", "_note_password"),
		Name:     getenv("DB_NAME", "_note_service"),
	}

	return Config{
		Database: dbCfg,
	}, nil
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return fallback
}
