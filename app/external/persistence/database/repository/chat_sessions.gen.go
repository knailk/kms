// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repository

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"kms/app/domain/entity"
)

func newChatSession(db *gorm.DB, opts ...gen.DOOption) chatSession {
	_chatSession := chatSession{}

	_chatSession.chatSessionDo.UseDB(db, opts...)
	_chatSession.chatSessionDo.UseModel(&entity.ChatSession{})

	tableName := _chatSession.chatSessionDo.TableName()
	_chatSession.ALL = field.NewAsterisk(tableName)
	_chatSession.ID = field.NewField(tableName, "id")
	_chatSession.Name = field.NewString(tableName, "name")
	_chatSession.LatestMessageID = field.NewField(tableName, "latest_message_id")
	_chatSession.ClassID = field.NewField(tableName, "class_id")
	_chatSession.CreatedAt = field.NewTime(tableName, "created_at")
	_chatSession.UpdatedAt = field.NewTime(tableName, "updated_at")
	_chatSession.IsDeleted = field.NewUint(tableName, "is_deleted")
	_chatSession.LatestMessage = chatSessionHasOneLatestMessage{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("LatestMessage", "entity.ChatMessage"),
	}

	_chatSession.ChatParticipants = chatSessionHasManyChatParticipants{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ChatParticipants", "entity.ChatParticipant"),
		User: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("ChatParticipants.User", "entity.User"),
		},
	}

	_chatSession.ChatMessages = chatSessionHasManyChatMessages{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ChatMessages", "entity.ChatMessage"),
	}

	_chatSession.fillFieldMap()

	return _chatSession
}

type chatSession struct {
	chatSessionDo

	ALL             field.Asterisk
	ID              field.Field
	Name            field.String
	LatestMessageID field.Field
	ClassID         field.Field
	CreatedAt       field.Time
	UpdatedAt       field.Time
	IsDeleted       field.Uint
	LatestMessage   chatSessionHasOneLatestMessage

	ChatParticipants chatSessionHasManyChatParticipants

	ChatMessages chatSessionHasManyChatMessages

	fieldMap map[string]field.Expr
}

func (c chatSession) Table(newTableName string) *chatSession {
	c.chatSessionDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chatSession) As(alias string) *chatSession {
	c.chatSessionDo.DO = *(c.chatSessionDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chatSession) updateTableName(table string) *chatSession {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.Name = field.NewString(table, "name")
	c.LatestMessageID = field.NewField(table, "latest_message_id")
	c.ClassID = field.NewField(table, "class_id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.IsDeleted = field.NewUint(table, "is_deleted")

	c.fillFieldMap()

	return c
}

func (c *chatSession) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chatSession) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
	c.fieldMap["latest_message_id"] = c.LatestMessageID
	c.fieldMap["class_id"] = c.ClassID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["is_deleted"] = c.IsDeleted

}

func (c chatSession) clone(db *gorm.DB) chatSession {
	c.chatSessionDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chatSession) replaceDB(db *gorm.DB) chatSession {
	c.chatSessionDo.ReplaceDB(db)
	return c
}

type chatSessionHasOneLatestMessage struct {
	db *gorm.DB

	field.RelationField
}

func (a chatSessionHasOneLatestMessage) Where(conds ...field.Expr) *chatSessionHasOneLatestMessage {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a chatSessionHasOneLatestMessage) WithContext(ctx context.Context) *chatSessionHasOneLatestMessage {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a chatSessionHasOneLatestMessage) Session(session *gorm.Session) *chatSessionHasOneLatestMessage {
	a.db = a.db.Session(session)
	return &a
}

func (a chatSessionHasOneLatestMessage) Model(m *entity.ChatSession) *chatSessionHasOneLatestMessageTx {
	return &chatSessionHasOneLatestMessageTx{a.db.Model(m).Association(a.Name())}
}

type chatSessionHasOneLatestMessageTx struct{ tx *gorm.Association }

func (a chatSessionHasOneLatestMessageTx) Find() (result *entity.ChatMessage, err error) {
	return result, a.tx.Find(&result)
}

func (a chatSessionHasOneLatestMessageTx) Append(values ...*entity.ChatMessage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a chatSessionHasOneLatestMessageTx) Replace(values ...*entity.ChatMessage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a chatSessionHasOneLatestMessageTx) Delete(values ...*entity.ChatMessage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a chatSessionHasOneLatestMessageTx) Clear() error {
	return a.tx.Clear()
}

func (a chatSessionHasOneLatestMessageTx) Count() int64 {
	return a.tx.Count()
}

type chatSessionHasManyChatParticipants struct {
	db *gorm.DB

	field.RelationField

	User struct {
		field.RelationField
	}
}

func (a chatSessionHasManyChatParticipants) Where(conds ...field.Expr) *chatSessionHasManyChatParticipants {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a chatSessionHasManyChatParticipants) WithContext(ctx context.Context) *chatSessionHasManyChatParticipants {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a chatSessionHasManyChatParticipants) Session(session *gorm.Session) *chatSessionHasManyChatParticipants {
	a.db = a.db.Session(session)
	return &a
}

func (a chatSessionHasManyChatParticipants) Model(m *entity.ChatSession) *chatSessionHasManyChatParticipantsTx {
	return &chatSessionHasManyChatParticipantsTx{a.db.Model(m).Association(a.Name())}
}

type chatSessionHasManyChatParticipantsTx struct{ tx *gorm.Association }

func (a chatSessionHasManyChatParticipantsTx) Find() (result []*entity.ChatParticipant, err error) {
	return result, a.tx.Find(&result)
}

func (a chatSessionHasManyChatParticipantsTx) Append(values ...*entity.ChatParticipant) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a chatSessionHasManyChatParticipantsTx) Replace(values ...*entity.ChatParticipant) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a chatSessionHasManyChatParticipantsTx) Delete(values ...*entity.ChatParticipant) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a chatSessionHasManyChatParticipantsTx) Clear() error {
	return a.tx.Clear()
}

func (a chatSessionHasManyChatParticipantsTx) Count() int64 {
	return a.tx.Count()
}

type chatSessionHasManyChatMessages struct {
	db *gorm.DB

	field.RelationField
}

func (a chatSessionHasManyChatMessages) Where(conds ...field.Expr) *chatSessionHasManyChatMessages {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a chatSessionHasManyChatMessages) WithContext(ctx context.Context) *chatSessionHasManyChatMessages {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a chatSessionHasManyChatMessages) Session(session *gorm.Session) *chatSessionHasManyChatMessages {
	a.db = a.db.Session(session)
	return &a
}

func (a chatSessionHasManyChatMessages) Model(m *entity.ChatSession) *chatSessionHasManyChatMessagesTx {
	return &chatSessionHasManyChatMessagesTx{a.db.Model(m).Association(a.Name())}
}

type chatSessionHasManyChatMessagesTx struct{ tx *gorm.Association }

func (a chatSessionHasManyChatMessagesTx) Find() (result []*entity.ChatMessage, err error) {
	return result, a.tx.Find(&result)
}

func (a chatSessionHasManyChatMessagesTx) Append(values ...*entity.ChatMessage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a chatSessionHasManyChatMessagesTx) Replace(values ...*entity.ChatMessage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a chatSessionHasManyChatMessagesTx) Delete(values ...*entity.ChatMessage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a chatSessionHasManyChatMessagesTx) Clear() error {
	return a.tx.Clear()
}

func (a chatSessionHasManyChatMessagesTx) Count() int64 {
	return a.tx.Count()
}

type chatSessionDo struct{ gen.DO }

type IChatSessionDo interface {
	gen.SubQuery
	Debug() IChatSessionDo
	WithContext(ctx context.Context) IChatSessionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IChatSessionDo
	WriteDB() IChatSessionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IChatSessionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IChatSessionDo
	Not(conds ...gen.Condition) IChatSessionDo
	Or(conds ...gen.Condition) IChatSessionDo
	Select(conds ...field.Expr) IChatSessionDo
	Where(conds ...gen.Condition) IChatSessionDo
	Order(conds ...field.Expr) IChatSessionDo
	Distinct(cols ...field.Expr) IChatSessionDo
	Omit(cols ...field.Expr) IChatSessionDo
	Join(table schema.Tabler, on ...field.Expr) IChatSessionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IChatSessionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IChatSessionDo
	Group(cols ...field.Expr) IChatSessionDo
	Having(conds ...gen.Condition) IChatSessionDo
	Limit(limit int) IChatSessionDo
	Offset(offset int) IChatSessionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IChatSessionDo
	Unscoped() IChatSessionDo
	Create(values ...*entity.ChatSession) error
	CreateInBatches(values []*entity.ChatSession, batchSize int) error
	Save(values ...*entity.ChatSession) error
	First() (*entity.ChatSession, error)
	Take() (*entity.ChatSession, error)
	Last() (*entity.ChatSession, error)
	Find() ([]*entity.ChatSession, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.ChatSession, err error)
	FindInBatches(result *[]*entity.ChatSession, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.ChatSession) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IChatSessionDo
	Assign(attrs ...field.AssignExpr) IChatSessionDo
	Joins(fields ...field.RelationField) IChatSessionDo
	Preload(fields ...field.RelationField) IChatSessionDo
	FirstOrInit() (*entity.ChatSession, error)
	FirstOrCreate() (*entity.ChatSession, error)
	FindByPage(offset int, limit int) (result []*entity.ChatSession, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IChatSessionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c chatSessionDo) Debug() IChatSessionDo {
	return c.withDO(c.DO.Debug())
}

func (c chatSessionDo) WithContext(ctx context.Context) IChatSessionDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chatSessionDo) ReadDB() IChatSessionDo {
	return c.Clauses(dbresolver.Read)
}

func (c chatSessionDo) WriteDB() IChatSessionDo {
	return c.Clauses(dbresolver.Write)
}

func (c chatSessionDo) Session(config *gorm.Session) IChatSessionDo {
	return c.withDO(c.DO.Session(config))
}

func (c chatSessionDo) Clauses(conds ...clause.Expression) IChatSessionDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chatSessionDo) Returning(value interface{}, columns ...string) IChatSessionDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chatSessionDo) Not(conds ...gen.Condition) IChatSessionDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chatSessionDo) Or(conds ...gen.Condition) IChatSessionDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chatSessionDo) Select(conds ...field.Expr) IChatSessionDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chatSessionDo) Where(conds ...gen.Condition) IChatSessionDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chatSessionDo) Order(conds ...field.Expr) IChatSessionDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chatSessionDo) Distinct(cols ...field.Expr) IChatSessionDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chatSessionDo) Omit(cols ...field.Expr) IChatSessionDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chatSessionDo) Join(table schema.Tabler, on ...field.Expr) IChatSessionDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chatSessionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IChatSessionDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chatSessionDo) RightJoin(table schema.Tabler, on ...field.Expr) IChatSessionDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chatSessionDo) Group(cols ...field.Expr) IChatSessionDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chatSessionDo) Having(conds ...gen.Condition) IChatSessionDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chatSessionDo) Limit(limit int) IChatSessionDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chatSessionDo) Offset(offset int) IChatSessionDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chatSessionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IChatSessionDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chatSessionDo) Unscoped() IChatSessionDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chatSessionDo) Create(values ...*entity.ChatSession) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chatSessionDo) CreateInBatches(values []*entity.ChatSession, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chatSessionDo) Save(values ...*entity.ChatSession) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chatSessionDo) First() (*entity.ChatSession, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatSession), nil
	}
}

func (c chatSessionDo) Take() (*entity.ChatSession, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatSession), nil
	}
}

func (c chatSessionDo) Last() (*entity.ChatSession, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatSession), nil
	}
}

func (c chatSessionDo) Find() ([]*entity.ChatSession, error) {
	result, err := c.DO.Find()
	return result.([]*entity.ChatSession), err
}

func (c chatSessionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.ChatSession, err error) {
	buf := make([]*entity.ChatSession, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chatSessionDo) FindInBatches(result *[]*entity.ChatSession, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chatSessionDo) Attrs(attrs ...field.AssignExpr) IChatSessionDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chatSessionDo) Assign(attrs ...field.AssignExpr) IChatSessionDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chatSessionDo) Joins(fields ...field.RelationField) IChatSessionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chatSessionDo) Preload(fields ...field.RelationField) IChatSessionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chatSessionDo) FirstOrInit() (*entity.ChatSession, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatSession), nil
	}
}

func (c chatSessionDo) FirstOrCreate() (*entity.ChatSession, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatSession), nil
	}
}

func (c chatSessionDo) FindByPage(offset int, limit int) (result []*entity.ChatSession, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c chatSessionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chatSessionDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chatSessionDo) Delete(models ...*entity.ChatSession) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chatSessionDo) withDO(do gen.Dao) *chatSessionDo {
	c.DO = *do.(*gen.DO)
	return c
}
