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

func prepareInsertStatement(db *sql.DB, topic string) (*sql.Stmt, error) {
	query := fmt.Sprintf("INSERT INTO %s (temperature, humidity) VALUES (?, ?)", topic)
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error preparing statement:", err)
	}
	return stmt, err
}

func prepareReadStatement(db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.Prepare("SELECT * FROM indoor_readings ORDER BY id DESC LIMIT 10")
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

func InsertDB(topic string, reading models.Reading) {
	db, err := openDatabase()
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := prepareInsertStatement(db, topic)
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

func ReadDB() []models.ReadingFull {

	readings := make([]models.ReadingFull, 0)

	db, err := openDatabase()
	if err != nil {
		return readings
	}
	defer db.Close()

	stmt, err := prepareReadStatement(db)
	if err != nil {
		return readings
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error executing query:", err)
		return readings
	}

	for rows.Next() {
		var reading models.ReadingFull
		if err := rows.Scan(&reading.ID, &reading.Temperature, &reading.Humidity, &reading.Timestamp); err != nil {
			log.Println("Error scanning row:", err)
			return readings
		}
		readings = append(readings, reading)
	}

	fmt.Println(readings)

	return readings
}
