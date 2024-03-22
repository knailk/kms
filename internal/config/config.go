package config

import (
	"errors"
	"flag"
	"fmt"
	"kms/app/config"
	"kms/database/sqldb"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	// defaultConfigFile default config file path
	defaultConfigFile = "development/config.{{env}}.yaml"
)

var (
	ErrNoConfig = errors.New("need to specify either config file or ENV")
)

func Init() (*config.Config, error) {
	confFilepath := flag.String("c", "", "configuration file for app")
	flag.Parse()

	// If no config file passed in, then check the ENV var
	// and select the default config file for the env
	if confFilepath == nil || *confFilepath == "" {
		env := os.Getenv("ENV")
		if env == "" {
			return nil, ErrNoConfig
		}
		if env == config.EnvLocal {
			err := godotenv.Load()
			if err != nil {
				panic(fmt.Sprintf("failed to load .env by error: %v", err))
			}
		}
		filepath := strings.ReplaceAll(defaultConfigFile, "{{env}}", env)
		confFilepath = &filepath
	}

	conf, err := config.LoadConfig(*confFilepath)
	return conf, err
}

// NewPostgreSQLDSN initializes a sqldb.PostgreSQLDSN given a Flags struct
func NewPostgreSQLDSN(cfg *config.Config) sqldb.PostgresqlDSN {
	return sqldb.PostgresqlDSN{
		Host:       cfg.DB.Host,
		Port:       cfg.DB.Port,
		DBName:     cfg.DB.DBName,
		SearchPath: cfg.DB.SearchPath,
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Migration:  cfg.DB.Migration,
	}
}
