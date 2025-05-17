package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWT_SECRET string

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("(WARN): .env file not found")
	}
	JWT_SECRET = os.Getenv("JWT_SECRET")
}

func GetDotEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
