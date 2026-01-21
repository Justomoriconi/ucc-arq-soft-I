package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapUrls(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
