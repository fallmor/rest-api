package database

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

// DatabaseSetup - set up a new database
func DatabaseSetup() (*gorm.DB, error) {
	fmt.Println("Setting up database")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	fmt.Println(connectionString)
	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		fmt.Println("can't connect to the DB")
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		fmt.Println("Can't ping the DB")
		return nil, err
	}
	return db, nil
}
