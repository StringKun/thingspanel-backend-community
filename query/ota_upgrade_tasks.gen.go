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

func newOtaUpgradeTask(db *gorm.DB, opts ...gen.DOOption) otaUpgradeTask {
	_otaUpgradeTask := otaUpgradeTask{}

	_otaUpgradeTask.otaUpgradeTaskDo.UseDB(db, opts...)
	_otaUpgradeTask.otaUpgradeTaskDo.UseModel(&model.OtaUpgradeTask{})

	tableName := _otaUpgradeTask.otaUpgradeTaskDo.TableName()
	_otaUpgradeTask.ALL = field.NewAsterisk(tableName)
	_otaUpgradeTask.ID = field.NewString(tableName, "id")
	_otaUpgradeTask.Name = field.NewString(tableName, "name")
	_otaUpgradeTask.OtaUpgradePackageID = field.NewString(tableName, "ota_upgrade_package_id")
	_otaUpgradeTask.Description = field.NewString(tableName, "description")
	_otaUpgradeTask.CreatedAt = field.NewTime(tableName, "created_at")
	_otaUpgradeTask.Remark = field.NewString(tableName, "remark")

	_otaUpgradeTask.fillFieldMap()

	return _otaUpgradeTask
}

type otaUpgradeTask struct {
	otaUpgradeTaskDo

	ALL                 field.Asterisk
	ID                  field.String // Id
	Name                field.String // 任务名称
	OtaUpgradePackageID field.String // 升级包id（外键，关联删除）
	Description         field.String // 描述
	CreatedAt           field.Time   // 创建时间
	Remark              field.String // 备注

	fieldMap map[string]field.Expr
}

func (o otaUpgradeTask) Table(newTableName string) *otaUpgradeTask {
	o.otaUpgradeTaskDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o otaUpgradeTask) As(alias string) *otaUpgradeTask {
	o.otaUpgradeTaskDo.DO = *(o.otaUpgradeTaskDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *otaUpgradeTask) updateTableName(table string) *otaUpgradeTask {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewString(table, "id")
	o.Name = field.NewString(table, "name")
	o.OtaUpgradePackageID = field.NewString(table, "ota_upgrade_package_id")
	o.Description = field.NewString(table, "description")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.Remark = field.NewString(table, "remark")

	o.fillFieldMap()

	return o
}

func (o *otaUpgradeTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *otaUpgradeTask) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["id"] = o.ID
	o.fieldMap["name"] = o.Name
	o.fieldMap["ota_upgrade_package_id"] = o.OtaUpgradePackageID
	o.fieldMap["description"] = o.Description
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["remark"] = o.Remark
}

func (o otaUpgradeTask) clone(db *gorm.DB) otaUpgradeTask {
	o.otaUpgradeTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o otaUpgradeTask) replaceDB(db *gorm.DB) otaUpgradeTask {
	o.otaUpgradeTaskDo.ReplaceDB(db)
	return o
}

type otaUpgradeTaskDo struct{ gen.DO }

type IOtaUpgradeTaskDo interface {
	gen.SubQuery
	Debug() IOtaUpgradeTaskDo
	WithContext(ctx context.Context) IOtaUpgradeTaskDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOtaUpgradeTaskDo
	WriteDB() IOtaUpgradeTaskDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOtaUpgradeTaskDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOtaUpgradeTaskDo
	Not(conds ...gen.Condition) IOtaUpgradeTaskDo
	Or(conds ...gen.Condition) IOtaUpgradeTaskDo
	Select(conds ...field.Expr) IOtaUpgradeTaskDo
	Where(conds ...gen.Condition) IOtaUpgradeTaskDo
	Order(conds ...field.Expr) IOtaUpgradeTaskDo
	Distinct(cols ...field.Expr) IOtaUpgradeTaskDo
	Omit(cols ...field.Expr) IOtaUpgradeTaskDo
	Join(table schema.Tabler, on ...field.Expr) IOtaUpgradeTaskDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOtaUpgradeTaskDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOtaUpgradeTaskDo
	Group(cols ...field.Expr) IOtaUpgradeTaskDo
	Having(conds ...gen.Condition) IOtaUpgradeTaskDo
	Limit(limit int) IOtaUpgradeTaskDo
	Offset(offset int) IOtaUpgradeTaskDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOtaUpgradeTaskDo
	Unscoped() IOtaUpgradeTaskDo
	Create(values ...*model.OtaUpgradeTask) error
	CreateInBatches(values []*model.OtaUpgradeTask, batchSize int) error
	Save(values ...*model.OtaUpgradeTask) error
	First() (*model.OtaUpgradeTask, error)
	Take() (*model.OtaUpgradeTask, error)
	Last() (*model.OtaUpgradeTask, error)
	Find() ([]*model.OtaUpgradeTask, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OtaUpgradeTask, err error)
	FindInBatches(result *[]*model.OtaUpgradeTask, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OtaUpgradeTask) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOtaUpgradeTaskDo
	Assign(attrs ...field.AssignExpr) IOtaUpgradeTaskDo
	Joins(fields ...field.RelationField) IOtaUpgradeTaskDo
	Preload(fields ...field.RelationField) IOtaUpgradeTaskDo
	FirstOrInit() (*model.OtaUpgradeTask, error)
	FirstOrCreate() (*model.OtaUpgradeTask, error)
	FindByPage(offset int, limit int) (result []*model.OtaUpgradeTask, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOtaUpgradeTaskDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o otaUpgradeTaskDo) Debug() IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Debug())
}

func (o otaUpgradeTaskDo) WithContext(ctx context.Context) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o otaUpgradeTaskDo) ReadDB() IOtaUpgradeTaskDo {
	return o.Clauses(dbresolver.Read)
}

func (o otaUpgradeTaskDo) WriteDB() IOtaUpgradeTaskDo {
	return o.Clauses(dbresolver.Write)
}

func (o otaUpgradeTaskDo) Session(config *gorm.Session) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Session(config))
}

func (o otaUpgradeTaskDo) Clauses(conds ...clause.Expression) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o otaUpgradeTaskDo) Returning(value interface{}, columns ...string) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o otaUpgradeTaskDo) Not(conds ...gen.Condition) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o otaUpgradeTaskDo) Or(conds ...gen.Condition) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o otaUpgradeTaskDo) Select(conds ...field.Expr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o otaUpgradeTaskDo) Where(conds ...gen.Condition) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o otaUpgradeTaskDo) Order(conds ...field.Expr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o otaUpgradeTaskDo) Distinct(cols ...field.Expr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o otaUpgradeTaskDo) Omit(cols ...field.Expr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o otaUpgradeTaskDo) Join(table schema.Tabler, on ...field.Expr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o otaUpgradeTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o otaUpgradeTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o otaUpgradeTaskDo) Group(cols ...field.Expr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o otaUpgradeTaskDo) Having(conds ...gen.Condition) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o otaUpgradeTaskDo) Limit(limit int) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o otaUpgradeTaskDo) Offset(offset int) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o otaUpgradeTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o otaUpgradeTaskDo) Unscoped() IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Unscoped())
}

func (o otaUpgradeTaskDo) Create(values ...*model.OtaUpgradeTask) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o otaUpgradeTaskDo) CreateInBatches(values []*model.OtaUpgradeTask, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o otaUpgradeTaskDo) Save(values ...*model.OtaUpgradeTask) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o otaUpgradeTaskDo) First() (*model.OtaUpgradeTask, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OtaUpgradeTask), nil
	}
}

func (o otaUpgradeTaskDo) Take() (*model.OtaUpgradeTask, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OtaUpgradeTask), nil
	}
}

func (o otaUpgradeTaskDo) Last() (*model.OtaUpgradeTask, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OtaUpgradeTask), nil
	}
}

func (o otaUpgradeTaskDo) Find() ([]*model.OtaUpgradeTask, error) {
	result, err := o.DO.Find()
	return result.([]*model.OtaUpgradeTask), err
}

func (o otaUpgradeTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OtaUpgradeTask, err error) {
	buf := make([]*model.OtaUpgradeTask, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o otaUpgradeTaskDo) FindInBatches(result *[]*model.OtaUpgradeTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o otaUpgradeTaskDo) Attrs(attrs ...field.AssignExpr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o otaUpgradeTaskDo) Assign(attrs ...field.AssignExpr) IOtaUpgradeTaskDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o otaUpgradeTaskDo) Joins(fields ...field.RelationField) IOtaUpgradeTaskDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o otaUpgradeTaskDo) Preload(fields ...field.RelationField) IOtaUpgradeTaskDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o otaUpgradeTaskDo) FirstOrInit() (*model.OtaUpgradeTask, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OtaUpgradeTask), nil
	}
}

func (o otaUpgradeTaskDo) FirstOrCreate() (*model.OtaUpgradeTask, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OtaUpgradeTask), nil
	}
}

func (o otaUpgradeTaskDo) FindByPage(offset int, limit int) (result []*model.OtaUpgradeTask, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o otaUpgradeTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o otaUpgradeTaskDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o otaUpgradeTaskDo) Delete(models ...*model.OtaUpgradeTask) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *otaUpgradeTaskDo) withDO(do gen.Dao) *otaUpgradeTaskDo {
	o.DO = *do.(*gen.DO)
	return o
}
