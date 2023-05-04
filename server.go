package main

import (
	"github.com/gin-gonic/gin"
)

// Main function
func main() {
	router := gin.Default()

	// Gin Routes
	router.GET("/health", health_check)

	// Gin Server
	router.Run("0.0.0.0:5000")
}
