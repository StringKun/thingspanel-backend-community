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

func newDeviceModelTelemetry(db *gorm.DB, opts ...gen.DOOption) deviceModelTelemetry {
	_deviceModelTelemetry := deviceModelTelemetry{}

	_deviceModelTelemetry.deviceModelTelemetryDo.UseDB(db, opts...)
	_deviceModelTelemetry.deviceModelTelemetryDo.UseModel(&model.DeviceModelTelemetry{})

	tableName := _deviceModelTelemetry.deviceModelTelemetryDo.TableName()
	_deviceModelTelemetry.ALL = field.NewAsterisk(tableName)
	_deviceModelTelemetry.ID = field.NewString(tableName, "id")
	_deviceModelTelemetry.DeviceTemplateID = field.NewString(tableName, "device_template_id")
	_deviceModelTelemetry.DataName = field.NewString(tableName, "data_name")
	_deviceModelTelemetry.DataIdentifier = field.NewString(tableName, "data_identifier")
	_deviceModelTelemetry.ReadWriteFlag = field.NewString(tableName, "read_write_flag")
	_deviceModelTelemetry.DataType = field.NewString(tableName, "data_type")
	_deviceModelTelemetry.Unit = field.NewString(tableName, "unit")
	_deviceModelTelemetry.Description = field.NewString(tableName, "description")
	_deviceModelTelemetry.AdditionalInfo = field.NewString(tableName, "additional_info")
	_deviceModelTelemetry.CreatedAt = field.NewTime(tableName, "created_at")
	_deviceModelTelemetry.UpdatedAt = field.NewTime(tableName, "updated_at")
	_deviceModelTelemetry.Remark = field.NewString(tableName, "remark")
	_deviceModelTelemetry.TenantID = field.NewString(tableName, "tenant_id")

	_deviceModelTelemetry.fillFieldMap()

	return _deviceModelTelemetry
}

type deviceModelTelemetry struct {
	deviceModelTelemetryDo

	ALL              field.Asterisk
	ID               field.String // id
	DeviceTemplateID field.String // 设备模板id
	DataName         field.String // 数据名称
	DataIdentifier   field.String // 数据标识符
	ReadWriteFlag    field.String // 读写标志R-读 W-写 RW-读写
	DataType         field.String // 数据类型String Number Boolean
	Unit             field.String // 单位
	Description      field.String // 描述
	AdditionalInfo   field.String // 附加信息
	CreatedAt        field.Time   // 创建时间
	UpdatedAt        field.Time   // 更新时间
	Remark           field.String // 备注
	TenantID         field.String

	fieldMap map[string]field.Expr
}

func (d deviceModelTelemetry) Table(newTableName string) *deviceModelTelemetry {
	d.deviceModelTelemetryDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deviceModelTelemetry) As(alias string) *deviceModelTelemetry {
	d.deviceModelTelemetryDo.DO = *(d.deviceModelTelemetryDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deviceModelTelemetry) updateTableName(table string) *deviceModelTelemetry {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.DeviceTemplateID = field.NewString(table, "device_template_id")
	d.DataName = field.NewString(table, "data_name")
	d.DataIdentifier = field.NewString(table, "data_identifier")
	d.ReadWriteFlag = field.NewString(table, "read_write_flag")
	d.DataType = field.NewString(table, "data_type")
	d.Unit = field.NewString(table, "unit")
	d.Description = field.NewString(table, "description")
	d.AdditionalInfo = field.NewString(table, "additional_info")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.Remark = field.NewString(table, "remark")
	d.TenantID = field.NewString(table, "tenant_id")

	d.fillFieldMap()

	return d
}

func (d *deviceModelTelemetry) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deviceModelTelemetry) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 13)
	d.fieldMap["id"] = d.ID
	d.fieldMap["device_template_id"] = d.DeviceTemplateID
	d.fieldMap["data_name"] = d.DataName
	d.fieldMap["data_identifier"] = d.DataIdentifier
	d.fieldMap["read_write_flag"] = d.ReadWriteFlag
	d.fieldMap["data_type"] = d.DataType
	d.fieldMap["unit"] = d.Unit
	d.fieldMap["description"] = d.Description
	d.fieldMap["additional_info"] = d.AdditionalInfo
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["remark"] = d.Remark
	d.fieldMap["tenant_id"] = d.TenantID
}

func (d deviceModelTelemetry) clone(db *gorm.DB) deviceModelTelemetry {
	d.deviceModelTelemetryDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deviceModelTelemetry) replaceDB(db *gorm.DB) deviceModelTelemetry {
	d.deviceModelTelemetryDo.ReplaceDB(db)
	return d
}

type deviceModelTelemetryDo struct{ gen.DO }

type IDeviceModelTelemetryDo interface {
	gen.SubQuery
	Debug() IDeviceModelTelemetryDo
	WithContext(ctx context.Context) IDeviceModelTelemetryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeviceModelTelemetryDo
	WriteDB() IDeviceModelTelemetryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeviceModelTelemetryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeviceModelTelemetryDo
	Not(conds ...gen.Condition) IDeviceModelTelemetryDo
	Or(conds ...gen.Condition) IDeviceModelTelemetryDo
	Select(conds ...field.Expr) IDeviceModelTelemetryDo
	Where(conds ...gen.Condition) IDeviceModelTelemetryDo
	Order(conds ...field.Expr) IDeviceModelTelemetryDo
	Distinct(cols ...field.Expr) IDeviceModelTelemetryDo
	Omit(cols ...field.Expr) IDeviceModelTelemetryDo
	Join(table schema.Tabler, on ...field.Expr) IDeviceModelTelemetryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceModelTelemetryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeviceModelTelemetryDo
	Group(cols ...field.Expr) IDeviceModelTelemetryDo
	Having(conds ...gen.Condition) IDeviceModelTelemetryDo
	Limit(limit int) IDeviceModelTelemetryDo
	Offset(offset int) IDeviceModelTelemetryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceModelTelemetryDo
	Unscoped() IDeviceModelTelemetryDo
	Create(values ...*model.DeviceModelTelemetry) error
	CreateInBatches(values []*model.DeviceModelTelemetry, batchSize int) error
	Save(values ...*model.DeviceModelTelemetry) error
	First() (*model.DeviceModelTelemetry, error)
	Take() (*model.DeviceModelTelemetry, error)
	Last() (*model.DeviceModelTelemetry, error)
	Find() ([]*model.DeviceModelTelemetry, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeviceModelTelemetry, err error)
	FindInBatches(result *[]*model.DeviceModelTelemetry, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DeviceModelTelemetry) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeviceModelTelemetryDo
	Assign(attrs ...field.AssignExpr) IDeviceModelTelemetryDo
	Joins(fields ...field.RelationField) IDeviceModelTelemetryDo
	Preload(fields ...field.RelationField) IDeviceModelTelemetryDo
	FirstOrInit() (*model.DeviceModelTelemetry, error)
	FirstOrCreate() (*model.DeviceModelTelemetry, error)
	FindByPage(offset int, limit int) (result []*model.DeviceModelTelemetry, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeviceModelTelemetryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deviceModelTelemetryDo) Debug() IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceModelTelemetryDo) WithContext(ctx context.Context) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceModelTelemetryDo) ReadDB() IDeviceModelTelemetryDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceModelTelemetryDo) WriteDB() IDeviceModelTelemetryDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceModelTelemetryDo) Session(config *gorm.Session) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceModelTelemetryDo) Clauses(conds ...clause.Expression) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceModelTelemetryDo) Returning(value interface{}, columns ...string) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceModelTelemetryDo) Not(conds ...gen.Condition) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceModelTelemetryDo) Or(conds ...gen.Condition) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceModelTelemetryDo) Select(conds ...field.Expr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceModelTelemetryDo) Where(conds ...gen.Condition) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceModelTelemetryDo) Order(conds ...field.Expr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceModelTelemetryDo) Distinct(cols ...field.Expr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceModelTelemetryDo) Omit(cols ...field.Expr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceModelTelemetryDo) Join(table schema.Tabler, on ...field.Expr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceModelTelemetryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceModelTelemetryDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceModelTelemetryDo) Group(cols ...field.Expr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceModelTelemetryDo) Having(conds ...gen.Condition) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceModelTelemetryDo) Limit(limit int) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceModelTelemetryDo) Offset(offset int) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceModelTelemetryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceModelTelemetryDo) Unscoped() IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceModelTelemetryDo) Create(values ...*model.DeviceModelTelemetry) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceModelTelemetryDo) CreateInBatches(values []*model.DeviceModelTelemetry, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceModelTelemetryDo) Save(values ...*model.DeviceModelTelemetry) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceModelTelemetryDo) First() (*model.DeviceModelTelemetry, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelTelemetry), nil
	}
}

func (d deviceModelTelemetryDo) Take() (*model.DeviceModelTelemetry, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelTelemetry), nil
	}
}

func (d deviceModelTelemetryDo) Last() (*model.DeviceModelTelemetry, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelTelemetry), nil
	}
}

func (d deviceModelTelemetryDo) Find() ([]*model.DeviceModelTelemetry, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeviceModelTelemetry), err
}

func (d deviceModelTelemetryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeviceModelTelemetry, err error) {
	buf := make([]*model.DeviceModelTelemetry, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceModelTelemetryDo) FindInBatches(result *[]*model.DeviceModelTelemetry, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceModelTelemetryDo) Attrs(attrs ...field.AssignExpr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceModelTelemetryDo) Assign(attrs ...field.AssignExpr) IDeviceModelTelemetryDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceModelTelemetryDo) Joins(fields ...field.RelationField) IDeviceModelTelemetryDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceModelTelemetryDo) Preload(fields ...field.RelationField) IDeviceModelTelemetryDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceModelTelemetryDo) FirstOrInit() (*model.DeviceModelTelemetry, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelTelemetry), nil
	}
}

func (d deviceModelTelemetryDo) FirstOrCreate() (*model.DeviceModelTelemetry, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceModelTelemetry), nil
	}
}

func (d deviceModelTelemetryDo) FindByPage(offset int, limit int) (result []*model.DeviceModelTelemetry, count int64, err error) {
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

func (d deviceModelTelemetryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceModelTelemetryDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceModelTelemetryDo) Delete(models ...*model.DeviceModelTelemetry) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceModelTelemetryDo) withDO(do gen.Dao) *deviceModelTelemetryDo {
	d.DO = *do.(*gen.DO)
	return d
}
