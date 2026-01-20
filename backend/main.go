package main

import (
	"github.com/gin-gonic/gin"
	"ucc-arq-soft-I/database"
)

func main() {
	database.Connect()
	database.Migrate()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	r.Run(":8080")
}
