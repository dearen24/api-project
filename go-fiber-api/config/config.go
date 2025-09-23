package config

import (
	"os"
	"path/filepath"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	Port       string
	JWTSecret  string
}

func LoadConfig() *Config {
	envPath := filepath.Join("..", ".env")

	// Load env vars from file
	if err := godotenv.Load(envPath); err != nil {
		log.Println("⚠️ Could not load .env file, relying on system environment")
	}
	return &Config{
		DBUser:     getEnv("DB_USER", "appuser"),
		DBPassword: getEnv("DB_PASSWORD", "apppassword"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "myapp"),
		Port:       getEnv("APP_PORT", "3000"),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
