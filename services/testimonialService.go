package services

import (
	"context"
	"fmt"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
)

func GetAllTestimonials(ctx context.Context) ([]models.Testimonial, error) {
	testimonials, err := repository.GetAllTestimonials(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve testimonials: %w", err)
	}
	return testimonials, nil
}