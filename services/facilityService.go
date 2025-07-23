package services

import (
	"context"
	"fmt"

	"example.com/m/v2/repository"
	"example.com/m/v2/models"
)


func GetAllFacilities(ctx context.Context) ([]models.Facility, error) {
	facilities, err := repository.GetAllFacilities(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve facilities: %w", err)
	}
	return facilities, nil
}