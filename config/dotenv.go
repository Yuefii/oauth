package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("(WARN): .env file not found")
	}
}

func GetDotEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
