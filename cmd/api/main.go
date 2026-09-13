package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/digital-worms/backend/internal/config"
	"github.com/digital-worms/backend/internal/database"
	"github.com/digital-worms/backend/internal/httpapi"
)

const (
	defaultHTTPAddr          = "localhost:8080"
	defaultReadHeaderTimeout = 5 * time.Second
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка конфига: %v", err)
	}

	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = defaultHTTPAddr
	}

	pool, err := database.ConnectPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к бд: %v", err)
	}
	defer pool.Close()
	log.Println("Подключение к бд прошло успешно")

	router := httpapi.NewRouter()

	server := &http.Server{
		Addr:              httpAddr,
		Handler:           router,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
	}

	log.Printf("Сервер запускается на %s", httpAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
