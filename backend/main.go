package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true, "msg": "pong"})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	_ = r.Run(":" + port)
}