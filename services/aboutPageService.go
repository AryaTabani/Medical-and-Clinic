package services

import (
	"context"
	"fmt"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
)

func GetAboutPageData(ctx context.Context) (*models.AboutPageData, error) {
	data, err := repository.GetAboutPageData(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get about page data: %w", err)
	}
	return data, nil
}
