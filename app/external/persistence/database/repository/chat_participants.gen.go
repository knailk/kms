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

func newChatParticipant(db *gorm.DB, opts ...gen.DOOption) chatParticipant {
	_chatParticipant := chatParticipant{}

	_chatParticipant.chatParticipantDo.UseDB(db, opts...)
	_chatParticipant.chatParticipantDo.UseModel(&entity.ChatParticipant{})

	tableName := _chatParticipant.chatParticipantDo.TableName()
	_chatParticipant.ALL = field.NewAsterisk(tableName)
	_chatParticipant.ID = field.NewField(tableName, "id")
	_chatParticipant.ChatSessionID = field.NewField(tableName, "chat_session_id")
	_chatParticipant.Username = field.NewString(tableName, "username")
	_chatParticipant.IsOwner = field.NewBool(tableName, "is_owner")
	_chatParticipant.CreatedAt = field.NewTime(tableName, "created_at")
	_chatParticipant.IsDeleted = field.NewUint(tableName, "is_deleted")
	_chatParticipant.User = chatParticipantHasOneUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "entity.User"),
	}

	_chatParticipant.fillFieldMap()

	return _chatParticipant
}

type chatParticipant struct {
	chatParticipantDo

	ALL           field.Asterisk
	ID            field.Field
	ChatSessionID field.Field
	Username      field.String
	IsOwner       field.Bool
	CreatedAt     field.Time
	IsDeleted     field.Uint
	User          chatParticipantHasOneUser

	fieldMap map[string]field.Expr
}

func (c chatParticipant) Table(newTableName string) *chatParticipant {
	c.chatParticipantDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chatParticipant) As(alias string) *chatParticipant {
	c.chatParticipantDo.DO = *(c.chatParticipantDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chatParticipant) updateTableName(table string) *chatParticipant {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.ChatSessionID = field.NewField(table, "chat_session_id")
	c.Username = field.NewString(table, "username")
	c.IsOwner = field.NewBool(table, "is_owner")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.IsDeleted = field.NewUint(table, "is_deleted")

	c.fillFieldMap()

	return c
}

func (c *chatParticipant) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chatParticipant) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["chat_session_id"] = c.ChatSessionID
	c.fieldMap["username"] = c.Username
	c.fieldMap["is_owner"] = c.IsOwner
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["is_deleted"] = c.IsDeleted

}

func (c chatParticipant) clone(db *gorm.DB) chatParticipant {
	c.chatParticipantDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chatParticipant) replaceDB(db *gorm.DB) chatParticipant {
	c.chatParticipantDo.ReplaceDB(db)
	return c
}

type chatParticipantHasOneUser struct {
	db *gorm.DB

	field.RelationField
}

func (a chatParticipantHasOneUser) Where(conds ...field.Expr) *chatParticipantHasOneUser {
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

func (a chatParticipantHasOneUser) WithContext(ctx context.Context) *chatParticipantHasOneUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a chatParticipantHasOneUser) Session(session *gorm.Session) *chatParticipantHasOneUser {
	a.db = a.db.Session(session)
	return &a
}

func (a chatParticipantHasOneUser) Model(m *entity.ChatParticipant) *chatParticipantHasOneUserTx {
	return &chatParticipantHasOneUserTx{a.db.Model(m).Association(a.Name())}
}

type chatParticipantHasOneUserTx struct{ tx *gorm.Association }

func (a chatParticipantHasOneUserTx) Find() (result *entity.User, err error) {
	return result, a.tx.Find(&result)
}

func (a chatParticipantHasOneUserTx) Append(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a chatParticipantHasOneUserTx) Replace(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a chatParticipantHasOneUserTx) Delete(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a chatParticipantHasOneUserTx) Clear() error {
	return a.tx.Clear()
}

func (a chatParticipantHasOneUserTx) Count() int64 {
	return a.tx.Count()
}

type chatParticipantDo struct{ gen.DO }

type IChatParticipantDo interface {
	gen.SubQuery
	Debug() IChatParticipantDo
	WithContext(ctx context.Context) IChatParticipantDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IChatParticipantDo
	WriteDB() IChatParticipantDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IChatParticipantDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IChatParticipantDo
	Not(conds ...gen.Condition) IChatParticipantDo
	Or(conds ...gen.Condition) IChatParticipantDo
	Select(conds ...field.Expr) IChatParticipantDo
	Where(conds ...gen.Condition) IChatParticipantDo
	Order(conds ...field.Expr) IChatParticipantDo
	Distinct(cols ...field.Expr) IChatParticipantDo
	Omit(cols ...field.Expr) IChatParticipantDo
	Join(table schema.Tabler, on ...field.Expr) IChatParticipantDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IChatParticipantDo
	RightJoin(table schema.Tabler, on ...field.Expr) IChatParticipantDo
	Group(cols ...field.Expr) IChatParticipantDo
	Having(conds ...gen.Condition) IChatParticipantDo
	Limit(limit int) IChatParticipantDo
	Offset(offset int) IChatParticipantDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IChatParticipantDo
	Unscoped() IChatParticipantDo
	Create(values ...*entity.ChatParticipant) error
	CreateInBatches(values []*entity.ChatParticipant, batchSize int) error
	Save(values ...*entity.ChatParticipant) error
	First() (*entity.ChatParticipant, error)
	Take() (*entity.ChatParticipant, error)
	Last() (*entity.ChatParticipant, error)
	Find() ([]*entity.ChatParticipant, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.ChatParticipant, err error)
	FindInBatches(result *[]*entity.ChatParticipant, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.ChatParticipant) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IChatParticipantDo
	Assign(attrs ...field.AssignExpr) IChatParticipantDo
	Joins(fields ...field.RelationField) IChatParticipantDo
	Preload(fields ...field.RelationField) IChatParticipantDo
	FirstOrInit() (*entity.ChatParticipant, error)
	FirstOrCreate() (*entity.ChatParticipant, error)
	FindByPage(offset int, limit int) (result []*entity.ChatParticipant, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IChatParticipantDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c chatParticipantDo) Debug() IChatParticipantDo {
	return c.withDO(c.DO.Debug())
}

func (c chatParticipantDo) WithContext(ctx context.Context) IChatParticipantDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chatParticipantDo) ReadDB() IChatParticipantDo {
	return c.Clauses(dbresolver.Read)
}

func (c chatParticipantDo) WriteDB() IChatParticipantDo {
	return c.Clauses(dbresolver.Write)
}

func (c chatParticipantDo) Session(config *gorm.Session) IChatParticipantDo {
	return c.withDO(c.DO.Session(config))
}

func (c chatParticipantDo) Clauses(conds ...clause.Expression) IChatParticipantDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chatParticipantDo) Returning(value interface{}, columns ...string) IChatParticipantDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chatParticipantDo) Not(conds ...gen.Condition) IChatParticipantDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chatParticipantDo) Or(conds ...gen.Condition) IChatParticipantDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chatParticipantDo) Select(conds ...field.Expr) IChatParticipantDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chatParticipantDo) Where(conds ...gen.Condition) IChatParticipantDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chatParticipantDo) Order(conds ...field.Expr) IChatParticipantDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chatParticipantDo) Distinct(cols ...field.Expr) IChatParticipantDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chatParticipantDo) Omit(cols ...field.Expr) IChatParticipantDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chatParticipantDo) Join(table schema.Tabler, on ...field.Expr) IChatParticipantDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chatParticipantDo) LeftJoin(table schema.Tabler, on ...field.Expr) IChatParticipantDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chatParticipantDo) RightJoin(table schema.Tabler, on ...field.Expr) IChatParticipantDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chatParticipantDo) Group(cols ...field.Expr) IChatParticipantDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chatParticipantDo) Having(conds ...gen.Condition) IChatParticipantDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chatParticipantDo) Limit(limit int) IChatParticipantDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chatParticipantDo) Offset(offset int) IChatParticipantDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chatParticipantDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IChatParticipantDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chatParticipantDo) Unscoped() IChatParticipantDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chatParticipantDo) Create(values ...*entity.ChatParticipant) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chatParticipantDo) CreateInBatches(values []*entity.ChatParticipant, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chatParticipantDo) Save(values ...*entity.ChatParticipant) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chatParticipantDo) First() (*entity.ChatParticipant, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatParticipant), nil
	}
}

func (c chatParticipantDo) Take() (*entity.ChatParticipant, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatParticipant), nil
	}
}

func (c chatParticipantDo) Last() (*entity.ChatParticipant, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatParticipant), nil
	}
}

func (c chatParticipantDo) Find() ([]*entity.ChatParticipant, error) {
	result, err := c.DO.Find()
	return result.([]*entity.ChatParticipant), err
}

func (c chatParticipantDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.ChatParticipant, err error) {
	buf := make([]*entity.ChatParticipant, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chatParticipantDo) FindInBatches(result *[]*entity.ChatParticipant, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chatParticipantDo) Attrs(attrs ...field.AssignExpr) IChatParticipantDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chatParticipantDo) Assign(attrs ...field.AssignExpr) IChatParticipantDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chatParticipantDo) Joins(fields ...field.RelationField) IChatParticipantDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chatParticipantDo) Preload(fields ...field.RelationField) IChatParticipantDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chatParticipantDo) FirstOrInit() (*entity.ChatParticipant, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatParticipant), nil
	}
}

func (c chatParticipantDo) FirstOrCreate() (*entity.ChatParticipant, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ChatParticipant), nil
	}
}

func (c chatParticipantDo) FindByPage(offset int, limit int) (result []*entity.ChatParticipant, count int64, err error) {
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

func (c chatParticipantDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chatParticipantDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chatParticipantDo) Delete(models ...*entity.ChatParticipant) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chatParticipantDo) withDO(do gen.Dao) *chatParticipantDo {
	c.DO = *do.(*gen.DO)
	return c
}
