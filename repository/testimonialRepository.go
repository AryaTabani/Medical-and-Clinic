package repository

import (
	"context"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func GetAllTestimonials(ctx context.Context) ([]models.Testimonial, error) {
	query := `SELECT id, quote, rating, author_name, author_role, author_image_url FROM testimonials`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testimonials []models.Testimonial
	for rows.Next() {
		var t models.Testimonial
		if err := rows.Scan(&t.ID, &t.Quote, &t.Rating, &t.Author_name, &t.Author_role, &t.Author_image_url); err != nil {
			return nil, err
		}
		testimonials = append(testimonials, t)
	}
	return testimonials, nil
}
