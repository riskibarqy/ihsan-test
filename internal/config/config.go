package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	AppMode string
	AppPort string
	DBHost  string
	DBPort  string
	DBName  string
	DBUser  string
	DBPass  string
	DBSSL   string
}

// LoadConfig loads environment variables into Config struct
func LoadConfig() *Config {
	cfg := &Config{
		AppMode: getEnv("APP_MODE", "development"),
		AppPort: getEnv("APP_PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBName:  getEnv("DB_NAME", ""),
		DBUser:  getEnv("DB_USERNAME", ""),
		DBPass:  getEnv("DB_PASSWORD", ""),
		DBSSL:   getEnv("DB_SSLMODE", "disable"),
	}

	// validate required environment
	if cfg.DBName == "" || cfg.DBUser == "" || cfg.DBPass == "" {
		log.Fatal("Database configuration is missing required values")
	}

	return cfg
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetDatabaseDSN returns the Postgre connection string
func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName, c.DBSSL,
	)
}
