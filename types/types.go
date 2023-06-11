package types

type Flight struct {
	Id             int    `json:"id"`
	Airlines       string `json:"airlines"`
	Flight_Number  string `json:"flightNumber"`
	Departure_Time string `json:"departureTime"`
	Arrival_Time   string `json:"arrivalTime"`
	Departure_Date string `json:"travelDate"`
	Departure_City string `json:"fromCity"`
	Arrival_City   string `json:"toCity"`
	Stops          string `json:"stops"`
}

type User struct {
	UserId string `json:"userId"`
	SubId  string `json:"subId"`
}

type LoginRequest struct {
	AccessToken string `json:"accessToken"`
}

type GoogleUserProfile struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	ProfilePic string `json:"picture"`
}

type Bag struct {
	Id               int    `json:"id"`
	NumberOfBags     int    `json:"numberOfBags"`
	Weight_Available string `json:"weightAvailable"`
	Price            int    `json:"price"`
	Flight           Flight `json:"flightInfo"`
}
