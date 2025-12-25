// Package config contains the configuration for the application
package config

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
func Load() Config {
	return Config{
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "note_user",
			Password: "note_password",
			Name:     "note_service",
		},
	}
}
