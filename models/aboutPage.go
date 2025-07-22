package models

type AboutPageData struct {
	Clinics_count int    `json:"clinics_count"`
	Doctors_count int    `json:"doctors_count"`
	Hours_open    string `json:"hours_open"`
	Info1      string `json:"headline"`
	Info2          string `json:"body"`
}
