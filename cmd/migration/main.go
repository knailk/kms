package main

import (
	"errors"
	"flag"
	"fmt"
	"kms/internal/config"
	"kms/pkg/logger"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	dir = "cmd/migration/migrations"
)

var (
	up   bool
	down bool
)

func main() {
	flag.BoolVar(&up, "up", false, "involves creating new tables, columns, or other database structures")
	flag.BoolVar(&down, "down", false, "involves dropping tables, columns, or other structures")
	flag.Parse()

	cfg, err := config.Init()
	if err != nil {
		logger.Error(err.Error())
	}

	// Open the postgres database
	db, err := gorm.Open(postgres.Open(config.NewPostgreSQLDSN(cfg).KeywordValueConnectionString()), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		logger.Fatal("sqldb.DBInit error: ", err)
	}
	// defer db.

	if up {
		err = migrate(db, "up")
		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	if down {
		err = migrate(db, "down")
		if err != nil {
			logger.Fatal(err.Error())
		}
	}
}

func migrate(db *gorm.DB, action string) (err error) {
	logger.Info("running migration ", action)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	files, err := filepath.Glob(filepath.Join(cwd, dir, fmt.Sprintf("*.%s.sql", action)))
	if err != nil {
		return errors.New("error when get files name")
	}


	for i, file := range files {
		logger.Info("Executing migration")
		data, err := os.ReadFile(file)
		if err != nil {
			return errors.New("error when read file")
		}

		tx := db.Exec(string(data))
		if tx.Error != nil {
			fmt.Println("error when exec: ", tx.Error)
			// return fmt.Errorf("error when exec query in file:%v", file)
		}

		fmt.Println("exec file number: ", i)
	}

	logger.Info("migration success with action ", action)

	return
}
