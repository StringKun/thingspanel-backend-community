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

func newDevice(db *gorm.DB, opts ...gen.DOOption) device {
	_device := device{}

	_device.deviceDo.UseDB(db, opts...)
	_device.deviceDo.UseModel(&model.Device{})

	tableName := _device.deviceDo.TableName()
	_device.ALL = field.NewAsterisk(tableName)
	_device.ID = field.NewString(tableName, "id")
	_device.Name = field.NewString(tableName, "name")
	_device.Voucher = field.NewString(tableName, "voucher")
	_device.TenantID = field.NewString(tableName, "tenant_id")
	_device.IsEnabled = field.NewString(tableName, "is_enabled")
	_device.ActivateFlag = field.NewString(tableName, "activate_flag")
	_device.CreatedAt = field.NewTime(tableName, "created_at")
	_device.UpdateAt = field.NewTime(tableName, "update_at")
	_device.DeviceNumber = field.NewString(tableName, "device_number")
	_device.ProductID = field.NewString(tableName, "product_id")
	_device.ParentID = field.NewString(tableName, "parent_id")
	_device.Protocol = field.NewString(tableName, "protocol")
	_device.Label = field.NewString(tableName, "label")
	_device.Location = field.NewString(tableName, "location")
	_device.SubDeviceAddr = field.NewString(tableName, "sub_device_addr")
	_device.CurrentVersion = field.NewString(tableName, "current_version")
	_device.AdditionalInfo = field.NewString(tableName, "additional_info")
	_device.ProtocolConfig = field.NewString(tableName, "protocol_config")
	_device.Remark1 = field.NewString(tableName, "remark1")
	_device.Remark2 = field.NewString(tableName, "remark2")
	_device.Remark3 = field.NewString(tableName, "remark3")
	_device.DeviceConfigID = field.NewString(tableName, "device_config_id")
	_device.BatchNumber = field.NewString(tableName, "batch_number")
	_device.ActivateAt = field.NewTime(tableName, "activate_at")
	_device.IsOnline = field.NewInt16(tableName, "is_online")
	_device.AccessWay = field.NewString(tableName, "access_way")
	_device.Description = field.NewString(tableName, "description")
	_device.ServiceAccessID = field.NewString(tableName, "service_access_id")

	_device.fillFieldMap()

	return _device
}

type device struct {
	deviceDo

	ALL            field.Asterisk
	ID             field.String // Id
	Name           field.String // 设备名称
	Voucher        field.String // 凭证
	TenantID       field.String // 租户id，外键，删除时阻止
	IsEnabled      field.String // 启用/禁用 enabled-启用 disabled-禁用 默认禁用，激活后默认启用
	ActivateFlag   field.String // 激活标志inactive-未激活 active-已激活
	CreatedAt      field.Time   // 创建时间
	UpdateAt       field.Time   // 更新时间
	DeviceNumber   field.String // 设备编号 没送默认和token一样
	ProductID      field.String // 产品id 外键，删除时阻止
	ParentID       field.String // 子设备的网关id
	Protocol       field.String // 通讯协议
	Label          field.String // 标签 单标签，英文逗号隔开
	Location       field.String // 地理位置
	SubDeviceAddr  field.String // 子设备地址
	CurrentVersion field.String // 当前固件版本
	AdditionalInfo field.String // 其他信息 阈值、图片等
	ProtocolConfig field.String // 协议表单配置
	Remark1        field.String
	Remark2        field.String
	Remark3        field.String
	/*
		设备配置id（外键）


	*/
	DeviceConfigID field.String
	/*
		批次编号

	*/
	BatchNumber     field.String
	ActivateAt      field.Time   // 激活日期
	IsOnline        field.Int16  // 是否在线 1-在线 0-离线
	AccessWay       field.String // 接入方式A-通过协议 B通过服务
	Description     field.String // 描述
	ServiceAccessID field.String

	fieldMap map[string]field.Expr
}

func (d device) Table(newTableName string) *device {
	d.deviceDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d device) As(alias string) *device {
	d.deviceDo.DO = *(d.deviceDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *device) updateTableName(table string) *device {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.Name = field.NewString(table, "name")
	d.Voucher = field.NewString(table, "voucher")
	d.TenantID = field.NewString(table, "tenant_id")
	d.IsEnabled = field.NewString(table, "is_enabled")
	d.ActivateFlag = field.NewString(table, "activate_flag")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdateAt = field.NewTime(table, "update_at")
	d.DeviceNumber = field.NewString(table, "device_number")
	d.ProductID = field.NewString(table, "product_id")
	d.ParentID = field.NewString(table, "parent_id")
	d.Protocol = field.NewString(table, "protocol")
	d.Label = field.NewString(table, "label")
	d.Location = field.NewString(table, "location")
	d.SubDeviceAddr = field.NewString(table, "sub_device_addr")
	d.CurrentVersion = field.NewString(table, "current_version")
	d.AdditionalInfo = field.NewString(table, "additional_info")
	d.ProtocolConfig = field.NewString(table, "protocol_config")
	d.Remark1 = field.NewString(table, "remark1")
	d.Remark2 = field.NewString(table, "remark2")
	d.Remark3 = field.NewString(table, "remark3")
	d.DeviceConfigID = field.NewString(table, "device_config_id")
	d.BatchNumber = field.NewString(table, "batch_number")
	d.ActivateAt = field.NewTime(table, "activate_at")
	d.IsOnline = field.NewInt16(table, "is_online")
	d.AccessWay = field.NewString(table, "access_way")
	d.Description = field.NewString(table, "description")
	d.ServiceAccessID = field.NewString(table, "service_access_id")

	d.fillFieldMap()

	return d
}

func (d *device) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *device) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 28)
	d.fieldMap["id"] = d.ID
	d.fieldMap["name"] = d.Name
	d.fieldMap["voucher"] = d.Voucher
	d.fieldMap["tenant_id"] = d.TenantID
	d.fieldMap["is_enabled"] = d.IsEnabled
	d.fieldMap["activate_flag"] = d.ActivateFlag
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["update_at"] = d.UpdateAt
	d.fieldMap["device_number"] = d.DeviceNumber
	d.fieldMap["product_id"] = d.ProductID
	d.fieldMap["parent_id"] = d.ParentID
	d.fieldMap["protocol"] = d.Protocol
	d.fieldMap["label"] = d.Label
	d.fieldMap["location"] = d.Location
	d.fieldMap["sub_device_addr"] = d.SubDeviceAddr
	d.fieldMap["current_version"] = d.CurrentVersion
	d.fieldMap["additional_info"] = d.AdditionalInfo
	d.fieldMap["protocol_config"] = d.ProtocolConfig
	d.fieldMap["remark1"] = d.Remark1
	d.fieldMap["remark2"] = d.Remark2
	d.fieldMap["remark3"] = d.Remark3
	d.fieldMap["device_config_id"] = d.DeviceConfigID
	d.fieldMap["batch_number"] = d.BatchNumber
	d.fieldMap["activate_at"] = d.ActivateAt
	d.fieldMap["is_online"] = d.IsOnline
	d.fieldMap["access_way"] = d.AccessWay
	d.fieldMap["description"] = d.Description
	d.fieldMap["service_access_id"] = d.ServiceAccessID
}

func (d device) clone(db *gorm.DB) device {
	d.deviceDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d device) replaceDB(db *gorm.DB) device {
	d.deviceDo.ReplaceDB(db)
	return d
}

type deviceDo struct{ gen.DO }

type IDeviceDo interface {
	gen.SubQuery
	Debug() IDeviceDo
	WithContext(ctx context.Context) IDeviceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeviceDo
	WriteDB() IDeviceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeviceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeviceDo
	Not(conds ...gen.Condition) IDeviceDo
	Or(conds ...gen.Condition) IDeviceDo
	Select(conds ...field.Expr) IDeviceDo
	Where(conds ...gen.Condition) IDeviceDo
	Order(conds ...field.Expr) IDeviceDo
	Distinct(cols ...field.Expr) IDeviceDo
	Omit(cols ...field.Expr) IDeviceDo
	Join(table schema.Tabler, on ...field.Expr) IDeviceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeviceDo
	Group(cols ...field.Expr) IDeviceDo
	Having(conds ...gen.Condition) IDeviceDo
	Limit(limit int) IDeviceDo
	Offset(offset int) IDeviceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceDo
	Unscoped() IDeviceDo
	Create(values ...*model.Device) error
	CreateInBatches(values []*model.Device, batchSize int) error
	Save(values ...*model.Device) error
	First() (*model.Device, error)
	Take() (*model.Device, error)
	Last() (*model.Device, error)
	Find() ([]*model.Device, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Device, err error)
	FindInBatches(result *[]*model.Device, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Device) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeviceDo
	Assign(attrs ...field.AssignExpr) IDeviceDo
	Joins(fields ...field.RelationField) IDeviceDo
	Preload(fields ...field.RelationField) IDeviceDo
	FirstOrInit() (*model.Device, error)
	FirstOrCreate() (*model.Device, error)
	FindByPage(offset int, limit int) (result []*model.Device, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeviceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deviceDo) Debug() IDeviceDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceDo) WithContext(ctx context.Context) IDeviceDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceDo) ReadDB() IDeviceDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceDo) WriteDB() IDeviceDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceDo) Session(config *gorm.Session) IDeviceDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceDo) Clauses(conds ...clause.Expression) IDeviceDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceDo) Returning(value interface{}, columns ...string) IDeviceDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceDo) Not(conds ...gen.Condition) IDeviceDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceDo) Or(conds ...gen.Condition) IDeviceDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceDo) Select(conds ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceDo) Where(conds ...gen.Condition) IDeviceDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceDo) Order(conds ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceDo) Distinct(cols ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceDo) Omit(cols ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceDo) Join(table schema.Tabler, on ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceDo) Group(cols ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceDo) Having(conds ...gen.Condition) IDeviceDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceDo) Limit(limit int) IDeviceDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceDo) Offset(offset int) IDeviceDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceDo) Unscoped() IDeviceDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceDo) Create(values ...*model.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceDo) CreateInBatches(values []*model.Device, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceDo) Save(values ...*model.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceDo) First() (*model.Device, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) Take() (*model.Device, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) Last() (*model.Device, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) Find() ([]*model.Device, error) {
	result, err := d.DO.Find()
	return result.([]*model.Device), err
}

func (d deviceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Device, err error) {
	buf := make([]*model.Device, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceDo) FindInBatches(result *[]*model.Device, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceDo) Attrs(attrs ...field.AssignExpr) IDeviceDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceDo) Assign(attrs ...field.AssignExpr) IDeviceDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceDo) Joins(fields ...field.RelationField) IDeviceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceDo) Preload(fields ...field.RelationField) IDeviceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceDo) FirstOrInit() (*model.Device, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) FirstOrCreate() (*model.Device, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Device), nil
	}
}

func (d deviceDo) FindByPage(offset int, limit int) (result []*model.Device, count int64, err error) {
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

func (d deviceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceDo) Delete(models ...*model.Device) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceDo) withDO(do gen.Dao) *deviceDo {
	d.DO = *do.(*gen.DO)
	return d
}
