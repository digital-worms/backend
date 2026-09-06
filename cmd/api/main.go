package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	defaultHTTPAddr          = "localhost:8080"
	defaultReadHeaderTimeout = 5 * time.Second
)

func main() {
	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = defaultHTTPAddr
	}

	router := registerRoutes()

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

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"status": "ok",
		},
	)
}

func registerRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/health", healthHandler)

	return router
}
