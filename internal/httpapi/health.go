package httpapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DatabasePinger interface {
	Ping(context.Context) error
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"status": "ok",
		},
	)
}

func readyHandler(database DatabasePinger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()

		if err := database.Ping(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable,
				gin.H{
					"status": "unavailable",
				},
			)
			return
		}

		c.JSON(http.StatusOK,
			gin.H{
				"status": "ok",
			},
		)
	}
}
