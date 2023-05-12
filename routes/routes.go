package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AllBookingsList struct {
	Id               int    `json:"id"`
	Flight_Number    string `json:"flightNumber"`
	Weight_Available string `json:"weightAvailable"`
	Departure_Time   string `json:"departureTime"`
	Arrival_Time     string `json:"arrivalTime"`
	Departure_Date   string `json:"departureDate"`
	Departure_City   string `json:"departureCity"`
	Arrival_City     string `json:"arrivalCity"`
	Price            string `json:"price"`
}

// Router functions
func Health_Check(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Healthy")
}

func All_Bookings(context *gin.Context) {
	url := os.Getenv("BACKEND_URL") + "/bookings"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject []AllBookingsList
	json.Unmarshal(responseData, &responseObject)
	context.IndentedJSON(http.StatusOK, responseObject)
}
