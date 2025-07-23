package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createtables()
}

func createtables() {
	createClinicsTable := `
	CREATE TABLE IF NOT EXISTS clinics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		address TEXT,
		phone_number TEXT,
		email TEXT,
		latitude REAL,
		longitude REAL,
		is_active BOOLEAN NOT NULL DEFAULT true
	);`

	_, err := DB.Exec(createClinicsTable)
	if err != nil {
		panic("Failed to create clinics table")
	}

	createDoctorsTable := `
	CREATE TABLE IF NOT EXISTS doctors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		specialty TEXT NOT NULL,
		image_url TEXT,
		approval_rating INTEGER,
		years_of_experience INTEGER,
		clinic_id INTEGER,
		is_active BOOLEAN NOT NULL DEFAULT true,
		FOREIGN KEY (clinic_id) REFERENCES clinics(id)
	);`
	_, err = DB.Exec(createDoctorsTable)
	if err != nil {
		panic("Failed to create doctors table")
	}
	createFacilitiesTable := `
	CREATE TABLE IF NOT EXISTS facilities (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		image_url TEXT NOT NULL
	);`

	_, err = DB.Exec(createFacilitiesTable)
	if err != nil {
		panic("Failed to create doctors table")
	}
	createTestimonialsTable := `
	CREATE TABLE IF NOT EXISTS testimonials (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		quote TEXT NOT NULL,
		rating INTEGER NOT NULL,
		author_name TEXT NOT NULL,
		author_role TEXT,
		author_image_url TEXT
	);`
	_, err = DB.Exec(createTestimonialsTable)
	if err != nil {
		panic("Failed to create testimonials table")
	}
}
