package sqldb

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Datastorer is an interface for working with the Database
type Datastorer interface {
	// BeginTx starts a pgx.Tx using the input context
	BeginTx()
	// RollbackTx rolls back the input pgx.Tx
	RollbackTx()
	// CommitTx commits the Tx
	CommitTx() error
	// GetTx to get the current transaction of this service
	GetTx() interface{}
	// RecoverTx to recover & roll back transaction of this service
	RecoverTx()
	// EndTx to end (commit or rollback) the current transaction of this service
	EndTx(err error) error
	// AssignToContext assign transaction to context
	AssignToContext(parentCtx context.Context) context.Context
}

// DBTX interface mirrors the interface generated by https://github.com/kyleconroy/sqlc
// to allow passing a Pool or a Tx
type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

// NewNullString returns a null if s is empty, otherwise it returns
// the string which was input
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// NewNullTime returns a null if t is the zero value for time.Time,
// otherwise it returns the time which was input
func NewNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

// NewNullInt64 returns a null if i == 0, otherwise it returns
// the int64 which was input.
func NewNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

// NewNullInt32 returns a null if i == 0, otherwise it returns
// the int32 which was input.
func NewNullInt32(i int32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

// NewNullUUID returns a null if i == uuid.Nil, otherwise it returns
// the int32 which was input.
func NewNullUUID(i uuid.UUID) uuid.NullUUID {
	if i == uuid.Nil {
		return uuid.NullUUID{}
	}
	return uuid.NullUUID{
		UUID:  i,
		Valid: true,
	}
}