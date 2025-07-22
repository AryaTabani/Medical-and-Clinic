package models

type clinic struct {
	name        string `json:"name"`
	address     string `json:"address"`
	is_active   bool   `json:"is_active"`
	call_number string `json:"call_number"`
	email       string `json:"email"`
}
