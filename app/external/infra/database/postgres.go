package database

import (
	"fmt"
	"kms/internal/database/sqldb"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	CONNECT_STRING = "host=%s user=%s password=%s dbname=%s port=%v"
)

func Open(dbs *sqldb.PostgreSQLDSN) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getConnectString(dbs)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("postgres open: %w", err)
	}

	return db, nil
}

func getConnectString(dbs *sqldb.PostgreSQLDSN) string {
	return fmt.Sprintf(CONNECT_STRING, dbs.Host, dbs.User, dbs.Password, dbs.DBName, dbs.Port)
}
