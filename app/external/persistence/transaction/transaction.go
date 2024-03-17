// Package use sql transaction, but in this project, we use gorm gen,
// a Friendly & Safer GORM powered by Code Generation. It already
// implements the transaction
package transaction

import (
	"context"
	"kms/database/sqldb"

	"gorm.io/gorm"
)

// dbService is a concrete implementation for a PostgreSQL database
type dbService struct {
	DB     *gorm.DB
	isDone bool
}

// NewDB is an initializer for the DB struct
func NewDB(db *gorm.DB) sqldb.Datastorer {
	return &dbService{DB: db, isDone: false}
}

// BeginTx returns an acquired transaction from the db pool and
// adds app specific error handling
func (db *dbService) BeginTx() {
	db.DB = db.DB.Begin()
}

// RollbackTx is a wrapper for sql.Tx.Rollback in order to expose from
// the Datastore interface. Proper error handling is also considered.
func (db *dbService) RollbackTx() {
	if !db.isDone {
		db.DB.Rollback()
		db.isDone = true
	}
}

// CommitTx is a wrapper for sql.Tx.Commit in order to expose from
// the Datastore interface. Proper error handling is also considered.
func (db *dbService) CommitTx() (err error) {
	if !db.isDone {
		err = db.DB.Commit().Error
		db.isDone = true
	}
	return err
}

// GetTx to get the current transaction of this service
func (db *dbService) GetTx() interface{} {
	return db.DB
}

// RecoverTx to recover & roll back transaction of this service
func (db *dbService) RecoverTx() {
	if p := recover(); p != nil {
		db.RollbackTx()
		// propagate panic info for caller
		panic(p)
	}
}

// EndTx to end (commit or rollback) the current transaction of this service
func (db *dbService) EndTx(err error) error {
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
	ContextKeyRepoTX string = "context_key_repo_tx"
)

// AssignToContext assign transaction to context
func (db *dbService) AssignToContext(parentCtx context.Context) context.Context {
	return context.WithValue(parentCtx, ContextKeyRepoTX, db.GetTx())
}
