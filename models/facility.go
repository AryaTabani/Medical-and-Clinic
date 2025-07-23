package models

type Facility struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"image_url"`
}