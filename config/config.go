package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl      string
	JWTSecret  string
	ServerPort string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	return &Config{
		DBUrl:      getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/hrportal?sslmode=disable"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecretkey"),
		ServerPort: getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
