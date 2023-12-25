package db

import (
	"ESPeather/internal/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "db/data.db"

func openDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func prepareInsertStatement(db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.Prepare("INSERT INTO indoor_readings (temperature, humidity) VALUES (?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
	}
	return stmt, err
}

func formatReadingValues(reading models.Reading) (string, string) {
	temperature := fmt.Sprintf("%.2f", reading.Temperature)
	humidity := fmt.Sprintf("%.2f", reading.Humidity)
	return temperature, humidity
}

func InsertDB(reading models.Reading) {
	db, err := openDatabase()
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := prepareInsertStatement(db)
	if err != nil {
		return
	}
	defer stmt.Close()

	temperature, humidity := formatReadingValues(reading)

	_, err = stmt.Exec(temperature, humidity)
	if err != nil {
		log.Println("Error inserting data into database:", err)
		return
	}

	fmt.Println("Data inserted into database successfully.")
}