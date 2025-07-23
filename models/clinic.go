package models

type Clinic struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Phone_number string  `json:"phone_number"`
	Email        string  `json:"email"`
	Address      string  `json:"address"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}
