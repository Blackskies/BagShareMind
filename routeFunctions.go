package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router functions
func health_check(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Healthy")
}
