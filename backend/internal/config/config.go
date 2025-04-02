package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{
		DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:     getEnvOrDefault("DB_PORT", "3306"),
		DBUser:     getEnvOrDefault("DB_USER", "root"),
		DBPassword: getEnvOrDefault("DB_PASSWORD", "maindatabase"),
		DBName:     getEnvOrDefault("DB_NAME", "techdocs"),
		JWTSecret:  getEnvOrDefault("JWT_SECRET", "your-secret-key"),
		ServerPort: getEnvOrDefault("SERVER_PORT", "8080"),
	}

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
} 