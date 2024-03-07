package transaction

import (
	"context"

	"gorm.io/gorm"
)

// DB is a concrete implementation for a PostgreSQL database
type DB struct {
	DB     *gorm.DB
	isDone bool
}

// NewDB is an initializer for the DB struct
func NewDB(db *gorm.DB) *DB {
	return &DB{DB: db, isDone: false}
}

// BeginTx returns an acquired transaction from the db pool and
// adds app specific error handling
func (db *DB) BeginTx() {
	db.DB = db.DB.Begin()
}

// RollbackTx is a wrapper for sql.Tx.Rollback in order to expose from
// the Datastore interface. Proper error handling is also considered.
func (db *DB) RollbackTx() {
	if !db.isDone {
		db.DB.Rollback()
		db.isDone = true
	}
}

// CommitTx is a wrapper for sql.Tx.Commit in order to expose from
// the Datastore interface. Proper error handling is also considered.
func (db *DB) CommitTx() (err error) {
	if !db.isDone {
		err = db.DB.Commit().Error
		db.isDone = true
	}
	return err
}

// GetTx to get the current transaction of this service
func (db *DB) GetTx() interface{} {
	return db.DB
}

// RecoverTx to recover & roll back transaction of this service
func (db *DB) RecoverTx() {
	if p := recover(); p != nil {
		db.RollbackTx()
		// propagate panic info for caller
		panic(p)
	}
}

// EndTx to end (commit or rollback) the current transaction of this service
func (db *DB) EndTx(err error) error {
	if err != nil {
		db.RollbackTx()
	} else {
		err = db.CommitTx()
		if err != nil {
			db.RollbackTx()
		}
	}
	return err
}

// ContextKeyRepoTX define key send through context
const (
	ContextKeyRepoTX = "context_key_repo_tx"
)

// AssignToContext assign transaction to context
func (db *DB) AssignToContext(parentCtx context.Context) context.Context {
	return context.WithValue(parentCtx, ContextKeyRepoTX, db.GetTx())
}
