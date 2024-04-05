// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repository

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q               = new(Query)
	ChatMessage     *chatMessage
	ChatParticipant *chatParticipant
	ChatSession     *chatSession
	CheckInOut      *checkInOut
	Class           *class
	Schedule        *schedule
	User            *user
	UserClass       *userClass
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	ChatMessage = &Q.ChatMessage
	ChatParticipant = &Q.ChatParticipant
	ChatSession = &Q.ChatSession
	CheckInOut = &Q.CheckInOut
	Class = &Q.Class
	Schedule = &Q.Schedule
	User = &Q.User
	UserClass = &Q.UserClass
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:              db,
		ChatMessage:     newChatMessage(db, opts...),
		ChatParticipant: newChatParticipant(db, opts...),
		ChatSession:     newChatSession(db, opts...),
		CheckInOut:      newCheckInOut(db, opts...),
		Class:           newClass(db, opts...),
		Schedule:        newSchedule(db, opts...),
		User:            newUser(db, opts...),
		UserClass:       newUserClass(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	ChatMessage     chatMessage
	ChatParticipant chatParticipant
	ChatSession     chatSession
	CheckInOut      checkInOut
	Class           class
	Schedule        schedule
	User            user
	UserClass       userClass
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		ChatMessage:     q.ChatMessage.clone(db),
		ChatParticipant: q.ChatParticipant.clone(db),
		ChatSession:     q.ChatSession.clone(db),
		CheckInOut:      q.CheckInOut.clone(db),
		Class:           q.Class.clone(db),
		Schedule:        q.Schedule.clone(db),
		User:            q.User.clone(db),
		UserClass:       q.UserClass.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		ChatMessage:     q.ChatMessage.replaceDB(db),
		ChatParticipant: q.ChatParticipant.replaceDB(db),
		ChatSession:     q.ChatSession.replaceDB(db),
		CheckInOut:      q.CheckInOut.replaceDB(db),
		Class:           q.Class.replaceDB(db),
		Schedule:        q.Schedule.replaceDB(db),
		User:            q.User.replaceDB(db),
		UserClass:       q.UserClass.replaceDB(db),
	}
}

type queryCtx struct {
	ChatMessage     IChatMessageDo
	ChatParticipant IChatParticipantDo
	ChatSession     IChatSessionDo
	CheckInOut      ICheckInOutDo
	Class           IClassDo
	Schedule        IScheduleDo
	User            IUserDo
	UserClass       IUserClassDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		ChatMessage:     q.ChatMessage.WithContext(ctx),
		ChatParticipant: q.ChatParticipant.WithContext(ctx),
		ChatSession:     q.ChatSession.WithContext(ctx),
		CheckInOut:      q.CheckInOut.WithContext(ctx),
		Class:           q.Class.WithContext(ctx),
		Schedule:        q.Schedule.WithContext(ctx),
		User:            q.User.WithContext(ctx),
		UserClass:       q.UserClass.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
