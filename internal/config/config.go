package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
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
	// Load .env file (ignore errors if it doesn't exist)
	_ = godotenv.Load()

	cfg := &Config{
		AppMode: getEnv("APP_MODE", "development"),
		AppPort: getEnv("APP_PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "postgres"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBName:  getEnv("DB_NAME", "ihsan_test_db"),
		DBUser:  getEnv("DB_USERNAME", "user"),
		DBPass:  getEnv("DB_PASSWORD", "password"),
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

// GetDatabaseDSN get database connection string
func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName, c.DBSSL,
	)
}
