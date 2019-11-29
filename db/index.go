package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// Required by gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func buildConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", host, port, user, dbName)
}

// Connect : connect to db
func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	connectionString := buildConnectionString()
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	db.LogMode(false)
	fmt.Println("Connected to database.")

	return db, nil
}
