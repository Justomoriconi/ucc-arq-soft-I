package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Controller "ucc-arq-soft-I/controllers/user"
)

func MapUrls(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api.POST("/signin", Controller.CreateUser)
}
