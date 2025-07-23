package services

import (
	"context"
	"fmt"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
)

func GetAllClinics(ctx context.Context) ([]models.Clinic, error) {
	clinics, err := repository.GetAllClinics(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve clinics: %w", err)
	}
	return clinics, nil
}