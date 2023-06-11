package main

import (
	"os"

	"example.com/bag-share/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Main function
func main() {

	os.Setenv("BACKEND_API_URL", "http://localhost:3000")
	os.Setenv("GOOGLE_USER_PROFILE_API_URL", "https://www.googleapis.com/oauth2/v1/userinfo")

	router := gin.Default()

	// enable all origins
	router.Use(cors.Default())

	// Gin Routes
	router.GET("/health", routes.Health_Check)
	router.POST("/login", routes.Login)
	router.GET("/flights", routes.All_Flights)
	router.GET("/bags/all", routes.All_Bags)
	router.GET("/bag", routes.Flight_Bags)

	// Gin Server
	router.Run("0.0.0.0:5000")
}
