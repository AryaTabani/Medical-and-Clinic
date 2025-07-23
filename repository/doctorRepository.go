package repository

import (
	"context"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func GetDoctors(ctx context.Context, specialty string) ([]models.Doctor, error) {
	query := `SELECT id, name, specialty, image_url, approval_rating, years_of_experience FROM doctors WHERE is_active = true`
	args := []interface{}{}

	if specialty != "" {
		query += " AND specialty = ?"
		args = append(args, specialty)
	}

	rows, err := db.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []models.Doctor
	for rows.Next() {
		var d models.Doctor
		if err := rows.Scan(&d.ID, &d.Name, &d.Specialty, &d.Image_url, &d.Approval_rating, &d.Years_of_experience); err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}
	return doctors, nil
}
