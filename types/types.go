package types

type AllFlightsList struct {
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

type AllBagsList struct {
	Id               int    `json:"id"`
	Flight_Number    string `json:"flightNumber"`
	Weight_Available string `json:"weightAvailable"`
	Departure_Time   string `json:"departureTime"`
	Arrival_Time     string `json:"arrivalTime"`
	Departure_Date   string `json:"travelDate"`
	Departure_City   string `json:"fromCity"`
	Arrival_City     string `json:"toCity"`
	Price            string `json:"price"`
}
