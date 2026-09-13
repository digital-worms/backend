package main

import (
	"net/http"
	"os"
	"time"

	"github.com/digital-worms/backend/internal/config"
	"github.com/digital-worms/backend/internal/database"
	"github.com/digital-worms/backend/internal/httpapi"
	"github.com/digital-worms/backend/internal/logger"
)

const (
	defaultHTTPAddr          = "localhost:8080"
	defaultReadHeaderTimeout = 5 * time.Second
)

func main() {
	appLogger := logger.New()
	cfg, err := config.Load()
	if err != nil {
		appLogger.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = defaultHTTPAddr
	}

	pool, err := database.ConnectPostgres(cfg.DatabaseURL)
	if err != nil {
		appLogger.Error("failed to connect to PostgreSQL", "error", err)
		os.Exit(1)
	}
	defer pool.Close()
	appLogger.Info("connected to PostgreSQL")

	router := httpapi.NewRouter()

	server := &http.Server{
		Addr:              httpAddr,
		Handler:           router,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
	}

	appLogger.Info("starting HTTP server", "address", httpAddr)
	if err := server.ListenAndServe(); err != nil {
		appLogger.Error("failed to start HTTP server", "error", err)
		os.Exit(1)
	}
}
