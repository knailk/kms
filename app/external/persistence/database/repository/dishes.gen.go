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

func newDish(db *gorm.DB, opts ...gen.DOOption) dish {
	_dish := dish{}

	_dish.dishDo.UseDB(db, opts...)
	_dish.dishDo.UseModel(&entity.Dish{})

	tableName := _dish.dishDo.TableName()
	_dish.ALL = field.NewAsterisk(tableName)
	_dish.ID = field.NewField(tableName, "id")
	_dish.DayOfWeek = field.NewString(tableName, "day_of_week")
	_dish.Date = field.NewInt(tableName, "date")
	_dish.Breakfast = field.NewString(tableName, "breakfast")
	_dish.EatLightly = field.NewString(tableName, "eat_lightly")
	_dish.Lunch = field.NewString(tableName, "lunch")
	_dish.AfternoonSnack = field.NewString(tableName, "afternoon_snack")
	_dish.Dinner = field.NewString(tableName, "dinner")
	_dish.CreatedAt = field.NewTime(tableName, "created_at")
	_dish.UpdatedAt = field.NewTime(tableName, "updated_at")

	_dish.fillFieldMap()

	return _dish
}

type dish struct {
	dishDo

	ALL            field.Asterisk
	ID             field.Field
	DayOfWeek      field.String
	Date           field.Int
	Breakfast      field.String
	EatLightly     field.String
	Lunch          field.String
	AfternoonSnack field.String
	Dinner         field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time

	fieldMap map[string]field.Expr
}

func (d dish) Table(newTableName string) *dish {
	d.dishDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dish) As(alias string) *dish {
	d.dishDo.DO = *(d.dishDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dish) updateTableName(table string) *dish {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewField(table, "id")
	d.DayOfWeek = field.NewString(table, "day_of_week")
	d.Date = field.NewInt(table, "date")
	d.Breakfast = field.NewString(table, "breakfast")
	d.EatLightly = field.NewString(table, "eat_lightly")
	d.Lunch = field.NewString(table, "lunch")
	d.AfternoonSnack = field.NewString(table, "afternoon_snack")
	d.Dinner = field.NewString(table, "dinner")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")

	d.fillFieldMap()

	return d
}

func (d *dish) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dish) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 10)
	d.fieldMap["id"] = d.ID
	d.fieldMap["day_of_week"] = d.DayOfWeek
	d.fieldMap["date"] = d.Date
	d.fieldMap["breakfast"] = d.Breakfast
	d.fieldMap["eat_lightly"] = d.EatLightly
	d.fieldMap["lunch"] = d.Lunch
	d.fieldMap["afternoon_snack"] = d.AfternoonSnack
	d.fieldMap["dinner"] = d.Dinner
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
}

func (d dish) clone(db *gorm.DB) dish {
	d.dishDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dish) replaceDB(db *gorm.DB) dish {
	d.dishDo.ReplaceDB(db)
	return d
}

type dishDo struct{ gen.DO }

type IDishDo interface {
	gen.SubQuery
	Debug() IDishDo
	WithContext(ctx context.Context) IDishDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDishDo
	WriteDB() IDishDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDishDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDishDo
	Not(conds ...gen.Condition) IDishDo
	Or(conds ...gen.Condition) IDishDo
	Select(conds ...field.Expr) IDishDo
	Where(conds ...gen.Condition) IDishDo
	Order(conds ...field.Expr) IDishDo
	Distinct(cols ...field.Expr) IDishDo
	Omit(cols ...field.Expr) IDishDo
	Join(table schema.Tabler, on ...field.Expr) IDishDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDishDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDishDo
	Group(cols ...field.Expr) IDishDo
	Having(conds ...gen.Condition) IDishDo
	Limit(limit int) IDishDo
	Offset(offset int) IDishDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDishDo
	Unscoped() IDishDo
	Create(values ...*entity.Dish) error
	CreateInBatches(values []*entity.Dish, batchSize int) error
	Save(values ...*entity.Dish) error
	First() (*entity.Dish, error)
	Take() (*entity.Dish, error)
	Last() (*entity.Dish, error)
	Find() ([]*entity.Dish, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Dish, err error)
	FindInBatches(result *[]*entity.Dish, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Dish) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDishDo
	Assign(attrs ...field.AssignExpr) IDishDo
	Joins(fields ...field.RelationField) IDishDo
	Preload(fields ...field.RelationField) IDishDo
	FirstOrInit() (*entity.Dish, error)
	FirstOrCreate() (*entity.Dish, error)
	FindByPage(offset int, limit int) (result []*entity.Dish, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDishDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dishDo) Debug() IDishDo {
	return d.withDO(d.DO.Debug())
}

func (d dishDo) WithContext(ctx context.Context) IDishDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dishDo) ReadDB() IDishDo {
	return d.Clauses(dbresolver.Read)
}

func (d dishDo) WriteDB() IDishDo {
	return d.Clauses(dbresolver.Write)
}

func (d dishDo) Session(config *gorm.Session) IDishDo {
	return d.withDO(d.DO.Session(config))
}

func (d dishDo) Clauses(conds ...clause.Expression) IDishDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dishDo) Returning(value interface{}, columns ...string) IDishDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dishDo) Not(conds ...gen.Condition) IDishDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dishDo) Or(conds ...gen.Condition) IDishDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dishDo) Select(conds ...field.Expr) IDishDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dishDo) Where(conds ...gen.Condition) IDishDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dishDo) Order(conds ...field.Expr) IDishDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dishDo) Distinct(cols ...field.Expr) IDishDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dishDo) Omit(cols ...field.Expr) IDishDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dishDo) Join(table schema.Tabler, on ...field.Expr) IDishDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dishDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDishDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dishDo) RightJoin(table schema.Tabler, on ...field.Expr) IDishDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dishDo) Group(cols ...field.Expr) IDishDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dishDo) Having(conds ...gen.Condition) IDishDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dishDo) Limit(limit int) IDishDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dishDo) Offset(offset int) IDishDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dishDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDishDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dishDo) Unscoped() IDishDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dishDo) Create(values ...*entity.Dish) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dishDo) CreateInBatches(values []*entity.Dish, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dishDo) Save(values ...*entity.Dish) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dishDo) First() (*entity.Dish, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Dish), nil
	}
}

func (d dishDo) Take() (*entity.Dish, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Dish), nil
	}
}

func (d dishDo) Last() (*entity.Dish, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Dish), nil
	}
}

func (d dishDo) Find() ([]*entity.Dish, error) {
	result, err := d.DO.Find()
	return result.([]*entity.Dish), err
}

func (d dishDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Dish, err error) {
	buf := make([]*entity.Dish, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dishDo) FindInBatches(result *[]*entity.Dish, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dishDo) Attrs(attrs ...field.AssignExpr) IDishDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dishDo) Assign(attrs ...field.AssignExpr) IDishDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dishDo) Joins(fields ...field.RelationField) IDishDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dishDo) Preload(fields ...field.RelationField) IDishDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dishDo) FirstOrInit() (*entity.Dish, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Dish), nil
	}
}

func (d dishDo) FirstOrCreate() (*entity.Dish, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Dish), nil
	}
}

func (d dishDo) FindByPage(offset int, limit int) (result []*entity.Dish, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dishDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dishDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dishDo) Delete(models ...*entity.Dish) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dishDo) withDO(do gen.Dao) *dishDo {
	d.DO = *do.(*gen.DO)
	return d
}