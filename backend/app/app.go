package app

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartRoute() error {
	r := gin.Default()

	cfg := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	r.Use(cors.New(cfg))

	MapUrls(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return r.Run(":" + port)
}
