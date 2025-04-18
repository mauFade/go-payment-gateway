package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mauFade/go-payment-gateway/internal/repository"
	"github.com/mauFade/go-payment-gateway/internal/service"
	"github.com/mauFade/go-payment-gateway/internal/web/server"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	} else {
		return defaultValue
	}
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

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("error connecting to db.")
	}

	defer db.Close()

	accRepository := repository.NewAccountRepository(db)

	accService := service.NewAccountService(accRepository)

	port := getEnv("HTTP_PORT", "8080")

	srv := server.NewServer(accService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("error starting http server.")
	}
}
