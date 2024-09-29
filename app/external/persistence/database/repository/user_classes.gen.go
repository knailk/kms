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

func newUserClass(db *gorm.DB, opts ...gen.DOOption) userClass {
	_userClass := userClass{}

	_userClass.userClassDo.UseDB(db, opts...)
	_userClass.userClassDo.UseModel(&entity.UserClass{})

	tableName := _userClass.userClassDo.TableName()
	_userClass.ALL = field.NewAsterisk(tableName)
	_userClass.ID = field.NewField(tableName, "id")
	_userClass.Username = field.NewString(tableName, "username")
	_userClass.ClassID = field.NewField(tableName, "class_id")
	_userClass.Status = field.NewString(tableName, "status")
	_userClass.CreatedAt = field.NewTime(tableName, "created_at")
	_userClass.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userClass.User = userClassHasOneUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "entity.User"),
	}

	_userClass.CheckInOuts = userClassHasManyCheckInOuts{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("CheckInOuts", "entity.CheckInOut"),
	}

	_userClass.fillFieldMap()

	return _userClass
}

type userClass struct {
	userClassDo

	ALL       field.Asterisk
	ID        field.Field
	Username  field.String
	ClassID   field.Field
	Status    field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	User      userClassHasOneUser

	CheckInOuts userClassHasManyCheckInOuts

	fieldMap map[string]field.Expr
}

func (u userClass) Table(newTableName string) *userClass {
	u.userClassDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userClass) As(alias string) *userClass {
	u.userClassDo.DO = *(u.userClassDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userClass) updateTableName(table string) *userClass {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewField(table, "id")
	u.Username = field.NewString(table, "username")
	u.ClassID = field.NewField(table, "class_id")
	u.Status = field.NewString(table, "status")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")

	u.fillFieldMap()

	return u
}

func (u *userClass) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userClass) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["id"] = u.ID
	u.fieldMap["username"] = u.Username
	u.fieldMap["class_id"] = u.ClassID
	u.fieldMap["status"] = u.Status
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt

}

func (u userClass) clone(db *gorm.DB) userClass {
	u.userClassDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userClass) replaceDB(db *gorm.DB) userClass {
	u.userClassDo.ReplaceDB(db)
	return u
}

type userClassHasOneUser struct {
	db *gorm.DB

	field.RelationField
}

func (a userClassHasOneUser) Where(conds ...field.Expr) *userClassHasOneUser {
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

func (a userClassHasOneUser) WithContext(ctx context.Context) *userClassHasOneUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userClassHasOneUser) Session(session *gorm.Session) *userClassHasOneUser {
	a.db = a.db.Session(session)
	return &a
}

func (a userClassHasOneUser) Model(m *entity.UserClass) *userClassHasOneUserTx {
	return &userClassHasOneUserTx{a.db.Model(m).Association(a.Name())}
}

type userClassHasOneUserTx struct{ tx *gorm.Association }

func (a userClassHasOneUserTx) Find() (result *entity.User, err error) {
	return result, a.tx.Find(&result)
}

func (a userClassHasOneUserTx) Append(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userClassHasOneUserTx) Replace(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userClassHasOneUserTx) Delete(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userClassHasOneUserTx) Clear() error {
	return a.tx.Clear()
}

func (a userClassHasOneUserTx) Count() int64 {
	return a.tx.Count()
}

type userClassHasManyCheckInOuts struct {
	db *gorm.DB

	field.RelationField
}

func (a userClassHasManyCheckInOuts) Where(conds ...field.Expr) *userClassHasManyCheckInOuts {
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

func (a userClassHasManyCheckInOuts) WithContext(ctx context.Context) *userClassHasManyCheckInOuts {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userClassHasManyCheckInOuts) Session(session *gorm.Session) *userClassHasManyCheckInOuts {
	a.db = a.db.Session(session)
	return &a
}

func (a userClassHasManyCheckInOuts) Model(m *entity.UserClass) *userClassHasManyCheckInOutsTx {
	return &userClassHasManyCheckInOutsTx{a.db.Model(m).Association(a.Name())}
}

type userClassHasManyCheckInOutsTx struct{ tx *gorm.Association }

func (a userClassHasManyCheckInOutsTx) Find() (result []*entity.CheckInOut, err error) {
	return result, a.tx.Find(&result)
}

func (a userClassHasManyCheckInOutsTx) Append(values ...*entity.CheckInOut) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userClassHasManyCheckInOutsTx) Replace(values ...*entity.CheckInOut) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userClassHasManyCheckInOutsTx) Delete(values ...*entity.CheckInOut) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userClassHasManyCheckInOutsTx) Clear() error {
	return a.tx.Clear()
}

func (a userClassHasManyCheckInOutsTx) Count() int64 {
	return a.tx.Count()
}

type userClassDo struct{ gen.DO }

type IUserClassDo interface {
	gen.SubQuery
	Debug() IUserClassDo
	WithContext(ctx context.Context) IUserClassDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserClassDo
	WriteDB() IUserClassDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserClassDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserClassDo
	Not(conds ...gen.Condition) IUserClassDo
	Or(conds ...gen.Condition) IUserClassDo
	Select(conds ...field.Expr) IUserClassDo
	Where(conds ...gen.Condition) IUserClassDo
	Order(conds ...field.Expr) IUserClassDo
	Distinct(cols ...field.Expr) IUserClassDo
	Omit(cols ...field.Expr) IUserClassDo
	Join(table schema.Tabler, on ...field.Expr) IUserClassDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserClassDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserClassDo
	Group(cols ...field.Expr) IUserClassDo
	Having(conds ...gen.Condition) IUserClassDo
	Limit(limit int) IUserClassDo
	Offset(offset int) IUserClassDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserClassDo
	Unscoped() IUserClassDo
	Create(values ...*entity.UserClass) error
	CreateInBatches(values []*entity.UserClass, batchSize int) error
	Save(values ...*entity.UserClass) error
	First() (*entity.UserClass, error)
	Take() (*entity.UserClass, error)
	Last() (*entity.UserClass, error)
	Find() ([]*entity.UserClass, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.UserClass, err error)
	FindInBatches(result *[]*entity.UserClass, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.UserClass) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserClassDo
	Assign(attrs ...field.AssignExpr) IUserClassDo
	Joins(fields ...field.RelationField) IUserClassDo
	Preload(fields ...field.RelationField) IUserClassDo
	FirstOrInit() (*entity.UserClass, error)
	FirstOrCreate() (*entity.UserClass, error)
	FindByPage(offset int, limit int) (result []*entity.UserClass, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserClassDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userClassDo) Debug() IUserClassDo {
	return u.withDO(u.DO.Debug())
}

func (u userClassDo) WithContext(ctx context.Context) IUserClassDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userClassDo) ReadDB() IUserClassDo {
	return u.Clauses(dbresolver.Read)
}

func (u userClassDo) WriteDB() IUserClassDo {
	return u.Clauses(dbresolver.Write)
}

func (u userClassDo) Session(config *gorm.Session) IUserClassDo {
	return u.withDO(u.DO.Session(config))
}

func (u userClassDo) Clauses(conds ...clause.Expression) IUserClassDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userClassDo) Returning(value interface{}, columns ...string) IUserClassDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userClassDo) Not(conds ...gen.Condition) IUserClassDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userClassDo) Or(conds ...gen.Condition) IUserClassDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userClassDo) Select(conds ...field.Expr) IUserClassDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userClassDo) Where(conds ...gen.Condition) IUserClassDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userClassDo) Order(conds ...field.Expr) IUserClassDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userClassDo) Distinct(cols ...field.Expr) IUserClassDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userClassDo) Omit(cols ...field.Expr) IUserClassDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userClassDo) Join(table schema.Tabler, on ...field.Expr) IUserClassDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userClassDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserClassDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userClassDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserClassDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userClassDo) Group(cols ...field.Expr) IUserClassDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userClassDo) Having(conds ...gen.Condition) IUserClassDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userClassDo) Limit(limit int) IUserClassDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userClassDo) Offset(offset int) IUserClassDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userClassDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserClassDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userClassDo) Unscoped() IUserClassDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userClassDo) Create(values ...*entity.UserClass) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userClassDo) CreateInBatches(values []*entity.UserClass, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userClassDo) Save(values ...*entity.UserClass) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userClassDo) First() (*entity.UserClass, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.UserClass), nil
	}
}

func (u userClassDo) Take() (*entity.UserClass, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.UserClass), nil
	}
}

func (u userClassDo) Last() (*entity.UserClass, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.UserClass), nil
	}
}

func (u userClassDo) Find() ([]*entity.UserClass, error) {
	result, err := u.DO.Find()
	return result.([]*entity.UserClass), err
}

func (u userClassDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.UserClass, err error) {
	buf := make([]*entity.UserClass, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userClassDo) FindInBatches(result *[]*entity.UserClass, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userClassDo) Attrs(attrs ...field.AssignExpr) IUserClassDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userClassDo) Assign(attrs ...field.AssignExpr) IUserClassDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userClassDo) Joins(fields ...field.RelationField) IUserClassDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userClassDo) Preload(fields ...field.RelationField) IUserClassDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userClassDo) FirstOrInit() (*entity.UserClass, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.UserClass), nil
	}
}

func (u userClassDo) FirstOrCreate() (*entity.UserClass, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.UserClass), nil
	}
}

func (u userClassDo) FindByPage(offset int, limit int) (result []*entity.UserClass, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userClassDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userClassDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userClassDo) Delete(models ...*entity.UserClass) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userClassDo) withDO(do gen.Dao) *userClassDo {
	u.DO = *do.(*gen.DO)
	return u
}
