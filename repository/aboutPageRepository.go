package repository

import (
	"context"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func GetAboutPageData(ctx context.Context) (*models.AboutPageData, error) {
	var clinicsCount int
	var doctorsCount int

	clinicQuery := `SELECT COUNT(*) FROM clinics WHERE is_active = true`
	if err := db.DB.QueryRowContext(ctx, clinicQuery).Scan(&clinicsCount); err != nil {
		return nil, err
	}

	doctorQuery := `SELECT COUNT(*) FROM doctors WHERE is_active = true`
	if err := db.DB.QueryRowContext(ctx, doctorQuery).Scan(&doctorsCount); err != nil {
		return nil, err
	}

	data := &models.AboutPageData{
		Clinics_count: clinicsCount,
		Doctors_count: doctorsCount,
		Hours_open:    "NUM",
		Info1:         "Text1",
		Info2:         "Text2",
	}

	return data, nil
}
