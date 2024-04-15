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

func newClass(db *gorm.DB, opts ...gen.DOOption) class {
	_class := class{}

	_class.classDo.UseDB(db, opts...)
	_class.classDo.UseModel(&entity.Class{})

	tableName := _class.classDo.TableName()
	_class.ALL = field.NewAsterisk(tableName)
	_class.ID = field.NewField(tableName, "id")
	_class.TeacherID = field.NewString(tableName, "teacher_id")
	_class.DriverID = field.NewString(tableName, "driver_id")
	_class.FromDate = field.NewInt64(tableName, "from_date")
	_class.ToDate = field.NewInt64(tableName, "to_date")
	_class.Status = field.NewString(tableName, "status")
	_class.ClassName = field.NewString(tableName, "class_name")
	_class.AgeGroup = field.NewInt(tableName, "age_group")
	_class.Price = field.NewFloat64(tableName, "price")
	_class.Currency = field.NewString(tableName, "currency")
	_class.CreatedAt = field.NewTime(tableName, "created_at")
	_class.UpdatedAt = field.NewTime(tableName, "updated_at")
	_class.IsDeleted = field.NewUint(tableName, "is_deleted")
	_class.Schedules = classHasManySchedules{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Schedules", "entity.Schedule"),
	}

	_class.User = classHasManyUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "entity.UserClass"),
	}

	_class.fillFieldMap()

	return _class
}

type class struct {
	classDo

	ALL       field.Asterisk
	ID        field.Field
	TeacherID field.String
	DriverID  field.String
	FromDate  field.Int64
	ToDate    field.Int64
	Status    field.String
	ClassName field.String
	AgeGroup  field.Int
	Price     field.Float64
	Currency  field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	IsDeleted field.Uint
	Schedules classHasManySchedules

	User classHasManyUser

	fieldMap map[string]field.Expr
}

func (c class) Table(newTableName string) *class {
	c.classDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c class) As(alias string) *class {
	c.classDo.DO = *(c.classDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *class) updateTableName(table string) *class {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.TeacherID = field.NewString(table, "teacher_id")
	c.DriverID = field.NewString(table, "driver_id")
	c.FromDate = field.NewInt64(table, "from_date")
	c.ToDate = field.NewInt64(table, "to_date")
	c.Status = field.NewString(table, "status")
	c.ClassName = field.NewString(table, "class_name")
	c.AgeGroup = field.NewInt(table, "age_group")
	c.Price = field.NewFloat64(table, "price")
	c.Currency = field.NewString(table, "currency")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.IsDeleted = field.NewUint(table, "is_deleted")

	c.fillFieldMap()

	return c
}

func (c *class) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *class) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 15)
	c.fieldMap["id"] = c.ID
	c.fieldMap["teacher_id"] = c.TeacherID
	c.fieldMap["driver_id"] = c.DriverID
	c.fieldMap["from_date"] = c.FromDate
	c.fieldMap["to_date"] = c.ToDate
	c.fieldMap["status"] = c.Status
	c.fieldMap["class_name"] = c.ClassName
	c.fieldMap["age_group"] = c.AgeGroup
	c.fieldMap["price"] = c.Price
	c.fieldMap["currency"] = c.Currency
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["is_deleted"] = c.IsDeleted

}

func (c class) clone(db *gorm.DB) class {
	c.classDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c class) replaceDB(db *gorm.DB) class {
	c.classDo.ReplaceDB(db)
	return c
}

type classHasManySchedules struct {
	db *gorm.DB

	field.RelationField
}

func (a classHasManySchedules) Where(conds ...field.Expr) *classHasManySchedules {
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

func (a classHasManySchedules) WithContext(ctx context.Context) *classHasManySchedules {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a classHasManySchedules) Session(session *gorm.Session) *classHasManySchedules {
	a.db = a.db.Session(session)
	return &a
}

func (a classHasManySchedules) Model(m *entity.Class) *classHasManySchedulesTx {
	return &classHasManySchedulesTx{a.db.Model(m).Association(a.Name())}
}

type classHasManySchedulesTx struct{ tx *gorm.Association }

func (a classHasManySchedulesTx) Find() (result []*entity.Schedule, err error) {
	return result, a.tx.Find(&result)
}

func (a classHasManySchedulesTx) Append(values ...*entity.Schedule) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a classHasManySchedulesTx) Replace(values ...*entity.Schedule) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a classHasManySchedulesTx) Delete(values ...*entity.Schedule) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a classHasManySchedulesTx) Clear() error {
	return a.tx.Clear()
}

func (a classHasManySchedulesTx) Count() int64 {
	return a.tx.Count()
}

type classHasManyUser struct {
	db *gorm.DB

	field.RelationField
}

func (a classHasManyUser) Where(conds ...field.Expr) *classHasManyUser {
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

func (a classHasManyUser) WithContext(ctx context.Context) *classHasManyUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a classHasManyUser) Session(session *gorm.Session) *classHasManyUser {
	a.db = a.db.Session(session)
	return &a
}

func (a classHasManyUser) Model(m *entity.Class) *classHasManyUserTx {
	return &classHasManyUserTx{a.db.Model(m).Association(a.Name())}
}

type classHasManyUserTx struct{ tx *gorm.Association }

func (a classHasManyUserTx) Find() (result []*entity.UserClass, err error) {
	return result, a.tx.Find(&result)
}

func (a classHasManyUserTx) Append(values ...*entity.UserClass) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a classHasManyUserTx) Replace(values ...*entity.UserClass) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a classHasManyUserTx) Delete(values ...*entity.UserClass) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a classHasManyUserTx) Clear() error {
	return a.tx.Clear()
}

func (a classHasManyUserTx) Count() int64 {
	return a.tx.Count()
}

type classDo struct{ gen.DO }

type IClassDo interface {
	gen.SubQuery
	Debug() IClassDo
	WithContext(ctx context.Context) IClassDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IClassDo
	WriteDB() IClassDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IClassDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IClassDo
	Not(conds ...gen.Condition) IClassDo
	Or(conds ...gen.Condition) IClassDo
	Select(conds ...field.Expr) IClassDo
	Where(conds ...gen.Condition) IClassDo
	Order(conds ...field.Expr) IClassDo
	Distinct(cols ...field.Expr) IClassDo
	Omit(cols ...field.Expr) IClassDo
	Join(table schema.Tabler, on ...field.Expr) IClassDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IClassDo
	RightJoin(table schema.Tabler, on ...field.Expr) IClassDo
	Group(cols ...field.Expr) IClassDo
	Having(conds ...gen.Condition) IClassDo
	Limit(limit int) IClassDo
	Offset(offset int) IClassDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IClassDo
	Unscoped() IClassDo
	Create(values ...*entity.Class) error
	CreateInBatches(values []*entity.Class, batchSize int) error
	Save(values ...*entity.Class) error
	First() (*entity.Class, error)
	Take() (*entity.Class, error)
	Last() (*entity.Class, error)
	Find() ([]*entity.Class, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Class, err error)
	FindInBatches(result *[]*entity.Class, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Class) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IClassDo
	Assign(attrs ...field.AssignExpr) IClassDo
	Joins(fields ...field.RelationField) IClassDo
	Preload(fields ...field.RelationField) IClassDo
	FirstOrInit() (*entity.Class, error)
	FirstOrCreate() (*entity.Class, error)
	FindByPage(offset int, limit int) (result []*entity.Class, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IClassDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c classDo) Debug() IClassDo {
	return c.withDO(c.DO.Debug())
}

func (c classDo) WithContext(ctx context.Context) IClassDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c classDo) ReadDB() IClassDo {
	return c.Clauses(dbresolver.Read)
}

func (c classDo) WriteDB() IClassDo {
	return c.Clauses(dbresolver.Write)
}

func (c classDo) Session(config *gorm.Session) IClassDo {
	return c.withDO(c.DO.Session(config))
}

func (c classDo) Clauses(conds ...clause.Expression) IClassDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c classDo) Returning(value interface{}, columns ...string) IClassDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c classDo) Not(conds ...gen.Condition) IClassDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c classDo) Or(conds ...gen.Condition) IClassDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c classDo) Select(conds ...field.Expr) IClassDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c classDo) Where(conds ...gen.Condition) IClassDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c classDo) Order(conds ...field.Expr) IClassDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c classDo) Distinct(cols ...field.Expr) IClassDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c classDo) Omit(cols ...field.Expr) IClassDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c classDo) Join(table schema.Tabler, on ...field.Expr) IClassDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c classDo) LeftJoin(table schema.Tabler, on ...field.Expr) IClassDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c classDo) RightJoin(table schema.Tabler, on ...field.Expr) IClassDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c classDo) Group(cols ...field.Expr) IClassDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c classDo) Having(conds ...gen.Condition) IClassDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c classDo) Limit(limit int) IClassDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c classDo) Offset(offset int) IClassDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c classDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IClassDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c classDo) Unscoped() IClassDo {
	return c.withDO(c.DO.Unscoped())
}

func (c classDo) Create(values ...*entity.Class) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c classDo) CreateInBatches(values []*entity.Class, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c classDo) Save(values ...*entity.Class) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c classDo) First() (*entity.Class, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Class), nil
	}
}

func (c classDo) Take() (*entity.Class, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Class), nil
	}
}

func (c classDo) Last() (*entity.Class, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Class), nil
	}
}

func (c classDo) Find() ([]*entity.Class, error) {
	result, err := c.DO.Find()
	return result.([]*entity.Class), err
}

func (c classDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Class, err error) {
	buf := make([]*entity.Class, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c classDo) FindInBatches(result *[]*entity.Class, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c classDo) Attrs(attrs ...field.AssignExpr) IClassDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c classDo) Assign(attrs ...field.AssignExpr) IClassDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c classDo) Joins(fields ...field.RelationField) IClassDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c classDo) Preload(fields ...field.RelationField) IClassDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c classDo) FirstOrInit() (*entity.Class, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Class), nil
	}
}

func (c classDo) FirstOrCreate() (*entity.Class, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Class), nil
	}
}

func (c classDo) FindByPage(offset int, limit int) (result []*entity.Class, count int64, err error) {
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

func (c classDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c classDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c classDo) Delete(models ...*entity.Class) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *classDo) withDO(do gen.Dao) *classDo {
	c.DO = *do.(*gen.DO)
	return c
}
