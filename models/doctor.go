package models

type Doctor struct {
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	Specialty           string `json:"specialty"`
	Image_url           string `json:"image_url"`
	Approval_rating     int    `json:"approval_rating"`
	Years_of_experience int    `json:"years_of_experience"`
}
