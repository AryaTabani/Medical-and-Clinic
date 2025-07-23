package models

type Testimonial struct {
	ID               int64  `json:"id"`
	Quote            string `json:"quote"`
	Rating           int    `json:"rating"`
	Author_name      string `json:"author_name"`
	Author_role      string `json:"author_role"`
	Author_image_url string `json:"author_image_url"`
}
