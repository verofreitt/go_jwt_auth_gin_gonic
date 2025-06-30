package main

import (
	"os"

	"go_jwt_auth_gin_gonic/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Register authentication and user routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// Example public endpoints
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "API is running"})
	})

	router.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{"version": "1.0.0", "author": "Your Name"})
	})

	router.Run(":" + port)
}
