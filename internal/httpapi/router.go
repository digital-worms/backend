package httpapi

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(database DatabasePinger) *gin.Engine {
	router := gin.Default()
	router.GET("/health", healthHandler)
	router.GET("/ready", readyHandler(database))

	return router
}
