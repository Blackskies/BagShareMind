package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"example.com/bag-share/types"
	"github.com/gin-gonic/gin"
)

// Router functions
func Health_Check(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Healthy")
}

func Get_Google_User_Profile(access_token string) types.GoogleUserProfile {
	emptyGoogleUserProfile := types.GoogleUserProfile{
		Id:         "",
		Email:      "",
		Name:       "",
		GivenName:  "",
		FamilyName: "",
		ProfilePic: "",
	}
	client := &http.Client{}
	url := os.Getenv("GOOGLE_USER_PROFILE_API_URL") + "?access_token=" + access_token
	// Create an HTTP request with custom headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return emptyGoogleUserProfile
	}
	req.Header.Add("Authorization", "Bearer "+access_token)
	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return emptyGoogleUserProfile
	}
	// Read the response body
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return emptyGoogleUserProfile
	}
	// Print the response body
	var responseObject types.GoogleUserProfile
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}

func Is_Existing_User(subId string) string {
	url := os.Getenv("BACKEND_API_URL") + "/users?subId=" + subId
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject []types.User
	json.Unmarshal(responseData, &responseObject)
	return responseObject[0].UserId
}

func User_Signup(subId string) string {
	url := os.Getenv("BACKEND_API_URL") + "/users"
	userId := ""
	for i := 0; i < 20; i++ {
		userId += (string)(rand.Intn(10) + 48) //48 is the ascii value for 0
	}
	user := types.User{
		UserId: userId,
		SubId:  subId,
	}
	marshalled, err := json.Marshal(user)
	bodyReader := bytes.NewReader(marshalled)
	if err != nil {
		log.Fatalf("impossible to marshall user: %s", err)
	}
	client := http.Client{Timeout: 10 * time.Second}
	// Create an HTTP request with custom headers
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		log.Fatalf("user signup build request failed: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// Send the HTTP request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("send request failed: %s", err)
	}
	responseData, err := io.ReadAll(response.Body)
	_ = responseData
	if err != nil {
		log.Fatalf("response read failed: %s", err)
	}
	return userId
}

func Login(context *gin.Context) {
	var loginRequest types.LoginRequest
	if err := context.BindJSON(&loginRequest); err != nil {
		context.IndentedJSON(http.StatusOK, gin.H{"errorCode": 400, "errorMessage": "please try again later"})
	}
	accessToken := loginRequest.AccessToken
	userProfile := Get_Google_User_Profile(accessToken)
	userId := Is_Existing_User(userProfile.Id)
	if userId == "" {
		userId = User_Signup(userProfile.Id)
	}
	userProfile.Id = userId
	context.JSON(http.StatusAccepted, &userProfile)
}

func All_Flights(context *gin.Context) {
	queryParams := context.Request.URL.Query()
	url := os.Getenv("BACKEND_API_URL") + "/flights?"
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
	var responseObject []types.Flight
	json.Unmarshal(responseData, &responseObject)
	context.IndentedJSON(http.StatusOK, responseObject)
}

func Flight_Bags(context *gin.Context) {
	url := os.Getenv("BACKEND_API_URL") + "/bags?"
	flightNumber := context.Request.URL.Query().Get("flightNumber")
	travelDate := context.Request.URL.Query().Get("travelDate")
	if flightNumber != "" {
		url += "flightInfo.flightNumber=" + flightNumber + "&"
	}
	if travelDate != "" {
		url += "flightInfo.travelDate" + travelDate
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
	var responseObject []types.Bag
	json.Unmarshal(responseData, &responseObject)
	context.IndentedJSON(http.StatusOK, responseObject)
}

func All_Bags(context *gin.Context) {
	url := os.Getenv("BACKEND_API_URL") + "/bags"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject []types.Bag
	json.Unmarshal(responseData, &responseObject)
	context.IndentedJSON(http.StatusOK, responseObject)
}
