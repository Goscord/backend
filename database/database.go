package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn string = fmt.Sprintf("host=127.0.0.1 user=%s password=%s dbname=%s port=%s", os.Getenv("POSTGRES_INIT_USER"), os.Getenv("POSTGRES_INIT_PASSWORD"), os.Getenv("POSTGRES_INIT_DBNAME"), os.Getenv("PORT"))

func Init() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return errors.New("failed to connect database")
	}

	// Migrate models

	dbConn, err := db.DB()
	if err != nil {
		return errors.New("failed to get database")
	}

	dbConn.Close()

	return nil
}

func GetDB() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
