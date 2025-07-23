package repository

import (
	"context"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func GetAllClinics(ctx context.Context) ([]models.Clinic, error) {
	query := `SELECT id, name, phone_number, email, address, latitude, longitude FROM clinics WHERE is_active = true`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clinics []models.Clinic
	for rows.Next() {
		var c models.Clinic
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone_number, &c.Email, &c.Address, &c.Latitude, &c.Longitude); err != nil {
			return nil, err
		}
		clinics = append(clinics, c)
	}

	return clinics, nil
}