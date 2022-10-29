package database

import (
	"errors"
	"os"

	"github.com/Goscord/DocGen/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() error {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECT_LINK")), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		return errors.New("failed to connect database")
	}

	db.AutoMigrate(&models.Version{})

	dbConn, err := db.DB()
	if err != nil {
		return errors.New("failed to get database")
	}

	dbConn.Close()

	return nil
}

func GetDB() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECT_LINK")), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}
