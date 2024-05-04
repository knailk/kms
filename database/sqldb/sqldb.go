package sqldb

import (
	"context"
	"fmt"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/internal/shutdown"
	"kms/pkg/logger"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	// DBHostEnv is the database host environment variable name
	DBHostEnv string = "DB_HOST"
	// DBPortEnv is the database port environment variable name
	DBPortEnv string = "DB_PORT"
	// DBNameEnv is the database name environment variable name
	DBNameEnv string = "DB_NAME"
	// DBUserEnv is the database user environment variable name
	DBUserEnv string = "DB_USER"
	// DBPasswordEnv is the database user password environment variable name
	DBPasswordEnv string = "DB_PASSWORD"
	// DBSearchPathEnv is the database search path environment variable name
	DBSearchPathEnv string = "DB_SEARCH_PATH"
)

// PostgresqlDSN is a PostgreSQL datasource name
type PostgresqlDSN struct {
	Host       string
	Port       int
	DBName     string
	SearchPath string
	User       string
	Password   string
	Migration  bool
}

// ConnectionURI returns a formatted PostgreSQL datasource "Keyword/Value Connection String"
// The general form for a connection URI is:
// postgresql://[userspec@][hostspec][/dbname][?paramspec]
// where userspec is
//
//	user[:password]
//
// and hostspec is:
//
//	[host][:port][,...]
//
// and paramspec is:
//
//	name=value[&...]
//
// The URI scheme designator can be either postgresql:// or postgres://.
// Each of the remaining URI parts is optional.
// The following examples illustrate valid URI syntax:
//
//	postgresql://
//	postgresql://localhost
//	postgresql://localhost:5433
//	postgresql://localhost/mydb
//	postgresql://user@localhost
//	postgresql://user:secret@localhost
//	postgresql://other@localhost/otherdb?connect_timeout=10&application_name=myapp
//	postgresql://host1:123,host2:456/somedb?target_session_attrs=any&application_name=myapp
func (dsn PostgresqlDSN) ConnectionURI() string {

	const uriSchemeDesignator string = "postgresql"

	var h string
	h = dsn.Host
	if dsn.Port != 0 {
		h += ":" + strconv.Itoa(dsn.Port)
	}

	u := url.URL{
		Scheme: uriSchemeDesignator,
		User:   url.User(dsn.User),
		Host:   h,
		Path:   dsn.DBName,
	}

	if dsn.SearchPath != "" {
		q := u.Query()
		q.Set("options", fmt.Sprintf("-csearch_path=%s", dsn.SearchPath))
		u.RawQuery = q.Encode()
	}

	return u.String()
}

// KeywordValueConnectionString returns a formatted PostgreSQL datasource "Keyword/Value Connection String"
func (dsn PostgresqlDSN) KeywordValueConnectionString() string {

	var s string

	// if db connection does not have a password (should only be for local testing and preferably never),
	// the password parameter must be removed from the string, otherwise the connection will fail.
	switch dsn.Password {
	case "":
		s = fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=disable", dsn.Host, dsn.Port, dsn.DBName, dsn.User)
	default:
		s = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", dsn.Host, dsn.Port, dsn.DBName, dsn.User, dsn.Password)
	}

	// if search path needs to be explicitly set, will be added to the end of the datasource string
	switch dsn.SearchPath {
	case "":
		return s
	default:
		return s + " " + fmt.Sprintf("search_path=%s", dsn.SearchPath)
	}
}

// DBInit creates a new db pool and establishes a connection.
// In addition, returns a Close function to defer closing the pool.
func DBInit(ctx context.Context, dsn PostgresqlDSN, tasks *shutdown.Tasks) (*gorm.DB, error) {
	const op errs.Op = "sqldb/NewPostgreSQLPool"

	if tasks.HasStopSignal() {
		return nil, errs.E(op, shutdown.ErrAbortedAsGotStopSignal)
	}

	cnn := dsn.KeywordValueConnectionString()

	// Open the postgres database
	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		return nil, errs.E(op, err)
	}

	logger.InfoFF("sql database opened for %s on port %d", logrus.Fields{}, dsn.Host, dsn.Port)

	sqlDB, err := dbClient.DB()
	if err != nil {
		return nil, errs.E(op, err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, errs.E(op, err)
	}

	if dsn.Migration {
		autoMigrate(dbClient)
	}

	return dbClient, nil
}

func autoMigrate(db *gorm.DB) {
	// migration for miniapp
	// create enums
	execLines := []string{
		// `CREATE TYPE "UserRole" AS ENUM ('admin', 'student', 'teacher', 'chef', 'driver');`,
		// `CREATE TYPE "MessageType" AS ENUM ('text', 'image', 'video', 'audio', 'file', 'link', 'voice', 'sticker');`,
		// `CREATE TYPE "CheckInOutAction" AS ENUM ('check_in', 'check_out');`,
	}
	for _, line := range execLines {
		db.Debug().Exec(line)
	}
	// create tables
	db.AutoMigrate(
		// add more models here for auto migrate when update models
		// entity.User{},
		entity.UserRequested{},
		// entity.ChatMessage{},
		// entity.ChatParticipant{},
		entity.ChatSession{},
		// entity.CheckInOut{},
		// entity.Class{},
		// entity.Schedule{},
		// entity.UserClass{},
	)
}
