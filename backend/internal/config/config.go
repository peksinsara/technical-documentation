package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	JWTSecret  string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{}

	// Set default database configuration
	config.DB.Host = getEnvOrDefault("DB_HOST", "localhost")
	config.DB.Port = getEnvOrDefault("DB_PORT", "3306")
	config.DB.User = getEnvOrDefault("DB_USER", "root")
	config.DB.Password = getEnvOrDefault("DB_PASSWORD", "maindatabase")
	config.DB.Name = getEnvOrDefault("DB_NAME", "techdocs")
	config.JWTSecret = getEnvOrDefault("JWT_SECRET", "your-secret-key")
	config.ServerPort = getEnvOrDefault("SERVER_PORT", "8081")

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func NewConfig() *Config {
	cfg := &Config{}

	// Set default database configuration
	cfg.DB.Host = "localhost"
	cfg.DB.Port = "3306"
	cfg.DB.User = "root"
	cfg.DB.Password = "maindatabase"
	cfg.DB.Name = "techdocs"

	return cfg
}
