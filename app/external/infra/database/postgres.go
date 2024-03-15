package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(cnn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cnn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("postgres open: %w", err)
	}

	return db, nil
}
