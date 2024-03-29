// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"snake/internal/models"
)

func newRecord(db *gorm.DB, opts ...gen.DOOption) record {
	_record := record{}

	_record.recordDo.UseDB(db, opts...)
	_record.recordDo.UseModel(&models.Record{})

	tableName := _record.recordDo.TableName()
	_record.ALL = field.NewAsterisk(tableName)
	_record.ID = field.NewInt32(tableName, "id")
	_record.AID = field.NewInt32(tableName, "a_id")
	_record.ASx = field.NewInt32(tableName, "a_sx")
	_record.ASy = field.NewInt32(tableName, "a_sy")
	_record.BID = field.NewInt32(tableName, "b_id")
	_record.BSx = field.NewInt32(tableName, "b_sx")
	_record.BSy = field.NewInt32(tableName, "b_sy")
	_record.ASteps = field.NewString(tableName, "a_steps")
	_record.BSteps = field.NewString(tableName, "b_steps")
	_record.Map = field.NewString(tableName, "map")
	_record.Loser = field.NewString(tableName, "loser")
	_record.Createtime = field.NewTime(tableName, "createtime")

	_record.fillFieldMap()

	return _record
}

type record struct {
	recordDo

	ALL        field.Asterisk
	ID         field.Int32
	AID        field.Int32
	ASx        field.Int32
	ASy        field.Int32
	BID        field.Int32
	BSx        field.Int32
	BSy        field.Int32
	ASteps     field.String
	BSteps     field.String
	Map        field.String
	Loser      field.String
	Createtime field.Time

	fieldMap map[string]field.Expr
}

func (r record) Table(newTableName string) *record {
	r.recordDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r record) As(alias string) *record {
	r.recordDo.DO = *(r.recordDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *record) updateTableName(table string) *record {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.AID = field.NewInt32(table, "a_id")
	r.ASx = field.NewInt32(table, "a_sx")
	r.ASy = field.NewInt32(table, "a_sy")
	r.BID = field.NewInt32(table, "b_id")
	r.BSx = field.NewInt32(table, "b_sx")
	r.BSy = field.NewInt32(table, "b_sy")
	r.ASteps = field.NewString(table, "a_steps")
	r.BSteps = field.NewString(table, "b_steps")
	r.Map = field.NewString(table, "map")
	r.Loser = field.NewString(table, "loser")
	r.Createtime = field.NewTime(table, "createtime")

	r.fillFieldMap()

	return r
}

func (r *record) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *record) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 12)
	r.fieldMap["id"] = r.ID
	r.fieldMap["a_id"] = r.AID
	r.fieldMap["a_sx"] = r.ASx
	r.fieldMap["a_sy"] = r.ASy
	r.fieldMap["b_id"] = r.BID
	r.fieldMap["b_sx"] = r.BSx
	r.fieldMap["b_sy"] = r.BSy
	r.fieldMap["a_steps"] = r.ASteps
	r.fieldMap["b_steps"] = r.BSteps
	r.fieldMap["map"] = r.Map
	r.fieldMap["loser"] = r.Loser
	r.fieldMap["createtime"] = r.Createtime
}

func (r record) clone(db *gorm.DB) record {
	r.recordDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r record) replaceDB(db *gorm.DB) record {
	r.recordDo.ReplaceDB(db)
	return r
}

type recordDo struct{ gen.DO }

type IRecordDo interface {
	gen.SubQuery
	Debug() IRecordDo
	WithContext(ctx context.Context) IRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRecordDo
	WriteDB() IRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRecordDo
	Not(conds ...gen.Condition) IRecordDo
	Or(conds ...gen.Condition) IRecordDo
	Select(conds ...field.Expr) IRecordDo
	Where(conds ...gen.Condition) IRecordDo
	Order(conds ...field.Expr) IRecordDo
	Distinct(cols ...field.Expr) IRecordDo
	Omit(cols ...field.Expr) IRecordDo
	Join(table schema.Tabler, on ...field.Expr) IRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRecordDo
	Group(cols ...field.Expr) IRecordDo
	Having(conds ...gen.Condition) IRecordDo
	Limit(limit int) IRecordDo
	Offset(offset int) IRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRecordDo
	Unscoped() IRecordDo
	Create(values ...*models.Record) error
	CreateInBatches(values []*models.Record, batchSize int) error
	Save(values ...*models.Record) error
	First() (*models.Record, error)
	Take() (*models.Record, error)
	Last() (*models.Record, error)
	Find() ([]*models.Record, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Record, err error)
	FindInBatches(result *[]*models.Record, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Record) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRecordDo
	Assign(attrs ...field.AssignExpr) IRecordDo
	Joins(fields ...field.RelationField) IRecordDo
	Preload(fields ...field.RelationField) IRecordDo
	FirstOrInit() (*models.Record, error)
	FirstOrCreate() (*models.Record, error)
	FindByPage(offset int, limit int) (result []*models.Record, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r recordDo) Debug() IRecordDo {
	return r.withDO(r.DO.Debug())
}

func (r recordDo) WithContext(ctx context.Context) IRecordDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r recordDo) ReadDB() IRecordDo {
	return r.Clauses(dbresolver.Read)
}

func (r recordDo) WriteDB() IRecordDo {
	return r.Clauses(dbresolver.Write)
}

func (r recordDo) Session(config *gorm.Session) IRecordDo {
	return r.withDO(r.DO.Session(config))
}

func (r recordDo) Clauses(conds ...clause.Expression) IRecordDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r recordDo) Returning(value interface{}, columns ...string) IRecordDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r recordDo) Not(conds ...gen.Condition) IRecordDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r recordDo) Or(conds ...gen.Condition) IRecordDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r recordDo) Select(conds ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r recordDo) Where(conds ...gen.Condition) IRecordDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r recordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IRecordDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r recordDo) Order(conds ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r recordDo) Distinct(cols ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r recordDo) Omit(cols ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r recordDo) Join(table schema.Tabler, on ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r recordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRecordDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r recordDo) RightJoin(table schema.Tabler, on ...field.Expr) IRecordDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r recordDo) Group(cols ...field.Expr) IRecordDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r recordDo) Having(conds ...gen.Condition) IRecordDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r recordDo) Limit(limit int) IRecordDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r recordDo) Offset(offset int) IRecordDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r recordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRecordDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r recordDo) Unscoped() IRecordDo {
	return r.withDO(r.DO.Unscoped())
}

func (r recordDo) Create(values ...*models.Record) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r recordDo) CreateInBatches(values []*models.Record, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r recordDo) Save(values ...*models.Record) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r recordDo) First() (*models.Record, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Record), nil
	}
}

func (r recordDo) Take() (*models.Record, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Record), nil
	}
}

func (r recordDo) Last() (*models.Record, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Record), nil
	}
}

func (r recordDo) Find() ([]*models.Record, error) {
	result, err := r.DO.Find()
	return result.([]*models.Record), err
}

func (r recordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Record, err error) {
	buf := make([]*models.Record, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r recordDo) FindInBatches(result *[]*models.Record, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r recordDo) Attrs(attrs ...field.AssignExpr) IRecordDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r recordDo) Assign(attrs ...field.AssignExpr) IRecordDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r recordDo) Joins(fields ...field.RelationField) IRecordDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r recordDo) Preload(fields ...field.RelationField) IRecordDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r recordDo) FirstOrInit() (*models.Record, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Record), nil
	}
}

func (r recordDo) FirstOrCreate() (*models.Record, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Record), nil
	}
}

func (r recordDo) FindByPage(offset int, limit int) (result []*models.Record, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r recordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r recordDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r recordDo) Delete(models ...*models.Record) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *recordDo) withDO(do gen.Dao) *recordDo {
	r.DO = *do.(*gen.DO)
	return r
}
