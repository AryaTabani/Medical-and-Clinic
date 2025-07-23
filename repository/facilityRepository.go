package repository

import (
	"context"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func GetAllFacilities(ctx context.Context) ([]models.Facility, error) {
	query := `SELECT id, name, description, image_url FROM facilities`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var facilities []models.Facility
	for rows.Next() {
		var f models.Facility
		if err := rows.Scan(&f.ID, &f.Name, &f.Description, &f.Image_url); err != nil {
			return nil, err
		}
		facilities = append(facilities, f)
	}

	return facilities, nil
}
