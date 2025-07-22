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
		is_active BOOLEAN NOT NULL DEFAULT true,
		call_number TEXT,
		email TEXT
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
		clinic_id INTEGER,
		is_active BOOLEAN NOT NULL DEFAULT true,
		FOREIGN KEY (clinic_id) REFERENCES clinics(id)
	);`

	_, err = DB.Exec(createDoctorsTable)
	if err != nil {
		panic("Failed to create doctors table")
	}
}
