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

	"project/internal/model"
)

func newRGroupDevice(db *gorm.DB, opts ...gen.DOOption) rGroupDevice {
	_rGroupDevice := rGroupDevice{}

	_rGroupDevice.rGroupDeviceDo.UseDB(db, opts...)
	_rGroupDevice.rGroupDeviceDo.UseModel(&model.RGroupDevice{})

	tableName := _rGroupDevice.rGroupDeviceDo.TableName()
	_rGroupDevice.ALL = field.NewAsterisk(tableName)
	_rGroupDevice.GroupID = field.NewString(tableName, "group_id")
	_rGroupDevice.DeviceID = field.NewString(tableName, "device_id")
	_rGroupDevice.TenantID = field.NewString(tableName, "tenant_id")

	_rGroupDevice.fillFieldMap()

	return _rGroupDevice
}

type rGroupDevice struct {
	rGroupDeviceDo

	ALL      field.Asterisk
	GroupID  field.String
	DeviceID field.String
	TenantID field.String

	fieldMap map[string]field.Expr
}

func (r rGroupDevice) Table(newTableName string) *rGroupDevice {
	r.rGroupDeviceDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r rGroupDevice) As(alias string) *rGroupDevice {
	r.rGroupDeviceDo.DO = *(r.rGroupDeviceDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *rGroupDevice) updateTableName(table string) *rGroupDevice {
	r.ALL = field.NewAsterisk(table)
	r.GroupID = field.NewString(table, "group_id")
	r.DeviceID = field.NewString(table, "device_id")
	r.TenantID = field.NewString(table, "tenant_id")

	r.fillFieldMap()

	return r
}

func (r *rGroupDevice) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *rGroupDevice) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 3)
	r.fieldMap["group_id"] = r.GroupID
	r.fieldMap["device_id"] = r.DeviceID
	r.fieldMap["tenant_id"] = r.TenantID
}

func (r rGroupDevice) clone(db *gorm.DB) rGroupDevice {
	r.rGroupDeviceDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r rGroupDevice) replaceDB(db *gorm.DB) rGroupDevice {
	r.rGroupDeviceDo.ReplaceDB(db)
	return r
}

type rGroupDeviceDo struct{ gen.DO }

type IRGroupDeviceDo interface {
	gen.SubQuery
	Debug() IRGroupDeviceDo
	WithContext(ctx context.Context) IRGroupDeviceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRGroupDeviceDo
	WriteDB() IRGroupDeviceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRGroupDeviceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRGroupDeviceDo
	Not(conds ...gen.Condition) IRGroupDeviceDo
	Or(conds ...gen.Condition) IRGroupDeviceDo
	Select(conds ...field.Expr) IRGroupDeviceDo
	Where(conds ...gen.Condition) IRGroupDeviceDo
	Order(conds ...field.Expr) IRGroupDeviceDo
	Distinct(cols ...field.Expr) IRGroupDeviceDo
	Omit(cols ...field.Expr) IRGroupDeviceDo
	Join(table schema.Tabler, on ...field.Expr) IRGroupDeviceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRGroupDeviceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRGroupDeviceDo
	Group(cols ...field.Expr) IRGroupDeviceDo
	Having(conds ...gen.Condition) IRGroupDeviceDo
	Limit(limit int) IRGroupDeviceDo
	Offset(offset int) IRGroupDeviceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRGroupDeviceDo
	Unscoped() IRGroupDeviceDo
	Create(values ...*model.RGroupDevice) error
	CreateInBatches(values []*model.RGroupDevice, batchSize int) error
	Save(values ...*model.RGroupDevice) error
	First() (*model.RGroupDevice, error)
	Take() (*model.RGroupDevice, error)
	Last() (*model.RGroupDevice, error)
	Find() ([]*model.RGroupDevice, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RGroupDevice, err error)
	FindInBatches(result *[]*model.RGroupDevice, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RGroupDevice) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRGroupDeviceDo
	Assign(attrs ...field.AssignExpr) IRGroupDeviceDo
	Joins(fields ...field.RelationField) IRGroupDeviceDo
	Preload(fields ...field.RelationField) IRGroupDeviceDo
	FirstOrInit() (*model.RGroupDevice, error)
	FirstOrCreate() (*model.RGroupDevice, error)
	FindByPage(offset int, limit int) (result []*model.RGroupDevice, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRGroupDeviceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r rGroupDeviceDo) Debug() IRGroupDeviceDo {
	return r.withDO(r.DO.Debug())
}

func (r rGroupDeviceDo) WithContext(ctx context.Context) IRGroupDeviceDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r rGroupDeviceDo) ReadDB() IRGroupDeviceDo {
	return r.Clauses(dbresolver.Read)
}

func (r rGroupDeviceDo) WriteDB() IRGroupDeviceDo {
	return r.Clauses(dbresolver.Write)
}

func (r rGroupDeviceDo) Session(config *gorm.Session) IRGroupDeviceDo {
	return r.withDO(r.DO.Session(config))
}

func (r rGroupDeviceDo) Clauses(conds ...clause.Expression) IRGroupDeviceDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r rGroupDeviceDo) Returning(value interface{}, columns ...string) IRGroupDeviceDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r rGroupDeviceDo) Not(conds ...gen.Condition) IRGroupDeviceDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r rGroupDeviceDo) Or(conds ...gen.Condition) IRGroupDeviceDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r rGroupDeviceDo) Select(conds ...field.Expr) IRGroupDeviceDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r rGroupDeviceDo) Where(conds ...gen.Condition) IRGroupDeviceDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r rGroupDeviceDo) Order(conds ...field.Expr) IRGroupDeviceDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r rGroupDeviceDo) Distinct(cols ...field.Expr) IRGroupDeviceDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r rGroupDeviceDo) Omit(cols ...field.Expr) IRGroupDeviceDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r rGroupDeviceDo) Join(table schema.Tabler, on ...field.Expr) IRGroupDeviceDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r rGroupDeviceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRGroupDeviceDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r rGroupDeviceDo) RightJoin(table schema.Tabler, on ...field.Expr) IRGroupDeviceDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r rGroupDeviceDo) Group(cols ...field.Expr) IRGroupDeviceDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r rGroupDeviceDo) Having(conds ...gen.Condition) IRGroupDeviceDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r rGroupDeviceDo) Limit(limit int) IRGroupDeviceDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r rGroupDeviceDo) Offset(offset int) IRGroupDeviceDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r rGroupDeviceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRGroupDeviceDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r rGroupDeviceDo) Unscoped() IRGroupDeviceDo {
	return r.withDO(r.DO.Unscoped())
}

func (r rGroupDeviceDo) Create(values ...*model.RGroupDevice) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r rGroupDeviceDo) CreateInBatches(values []*model.RGroupDevice, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r rGroupDeviceDo) Save(values ...*model.RGroupDevice) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r rGroupDeviceDo) First() (*model.RGroupDevice, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RGroupDevice), nil
	}
}

func (r rGroupDeviceDo) Take() (*model.RGroupDevice, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RGroupDevice), nil
	}
}

func (r rGroupDeviceDo) Last() (*model.RGroupDevice, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RGroupDevice), nil
	}
}

func (r rGroupDeviceDo) Find() ([]*model.RGroupDevice, error) {
	result, err := r.DO.Find()
	return result.([]*model.RGroupDevice), err
}

func (r rGroupDeviceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RGroupDevice, err error) {
	buf := make([]*model.RGroupDevice, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r rGroupDeviceDo) FindInBatches(result *[]*model.RGroupDevice, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r rGroupDeviceDo) Attrs(attrs ...field.AssignExpr) IRGroupDeviceDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r rGroupDeviceDo) Assign(attrs ...field.AssignExpr) IRGroupDeviceDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r rGroupDeviceDo) Joins(fields ...field.RelationField) IRGroupDeviceDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r rGroupDeviceDo) Preload(fields ...field.RelationField) IRGroupDeviceDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r rGroupDeviceDo) FirstOrInit() (*model.RGroupDevice, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RGroupDevice), nil
	}
}

func (r rGroupDeviceDo) FirstOrCreate() (*model.RGroupDevice, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RGroupDevice), nil
	}
}

func (r rGroupDeviceDo) FindByPage(offset int, limit int) (result []*model.RGroupDevice, count int64, err error) {
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

func (r rGroupDeviceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r rGroupDeviceDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r rGroupDeviceDo) Delete(models ...*model.RGroupDevice) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *rGroupDeviceDo) withDO(do gen.Dao) *rGroupDeviceDo {
	r.DO = *do.(*gen.DO)
	return r
}
