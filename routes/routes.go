package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"example.com/bag-share/types"
	"github.com/gin-gonic/gin"
)

// Router functions
func Health_Check(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Healthy")
}

func All_Flights(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	fmt.Println(context.Request.URL.Query())
	url := os.Getenv("BACKEND_URL") + "/flights?"
	count := 0
	for key, value := range queryParams {
		url += key + "=" + value[0]
		if count != len(queryParams)-1 {
			url += "&"
		}
		count++
	}
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject []types.AllFlightsList
	json.Unmarshal(responseData, &responseObject)
	context.IndentedJSON(http.StatusOK, responseObject)
}

func All_Bags(context *gin.Context) {
	url := os.Getenv("BACKEND_URL") + "/bags"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject []types.AllBagsList
	json.Unmarshal(responseData, &responseObject)
	context.IndentedJSON(http.StatusOK, responseObject)
}
