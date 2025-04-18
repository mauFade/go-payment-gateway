package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func getEnv(key, defaultValue string) string {
	return ""
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file.")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)

}
