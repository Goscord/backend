package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn string = fmt.Sprintf("postgres://%s:%s@host.docker.internal:5432/%s", os.Getenv("POSTGRES_INIT_USER"), os.Getenv("POSTGRES_INIT_PASSWORD"), os.Getenv("POSTGRES_INIT_DBNAME"))

func Init() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

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
	return gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}
