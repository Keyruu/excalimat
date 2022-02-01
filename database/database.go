package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/keyruu/excalimat-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func Connect() {
	var err error
	p := config.DB_PORT
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("db port isn't a number ??")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST, port, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connection to database established")
}
