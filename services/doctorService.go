package services

import (
	"context"
	"fmt"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
)

func GetDoctors(ctx context.Context, specialty string) ([]models.Doctor, error) {
	doctors, err := repository.GetDoctors(ctx, specialty)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve doctors: %w", err)
	}
	return doctors, nil
}
