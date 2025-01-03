package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

type House struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

type Reservation struct {
	ID        int
	HouseID   int
	Name      string
	Email     string
	StartDate string
	EndDate   string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection successful")
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}
}

func GetHouses() ([]House, error) {
	rows, err := DB.Query("SELECT id, name, description, price FROM houses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var houses []House
	for rows.Next() {
		var house House
		if err := rows.Scan(&house.ID, &house.Name, &house.Description, &house.Price); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		houses = append(houses, house)
	}
	return houses, nil
}

func AddReservation(res Reservation) error {
	_, err := DB.Exec(
		"INSERT INTO reservations (house_id, name, email, start_date, end_date) VALUES (?, ?, ?, ?, ?)",
		res.HouseID, res.Name, res.Email, res.StartDate, res.EndDate,
	)
	return err
}
