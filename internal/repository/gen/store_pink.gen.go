// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"crmeb_go/internal/model"
)

func newStorePink(db *gorm.DB, opts ...gen.DOOption) storePink {
	_storePink := storePink{}

	_storePink.storePinkDo.UseDB(db, opts...)
	_storePink.storePinkDo.UseModel(&model.StorePink{})

	tableName := _storePink.storePinkDo.TableName()
	_storePink.ALL = field.NewAsterisk(tableName)
	_storePink.ID = field.NewInt64(tableName, "id")
	_storePink.UID = field.NewInt64(tableName, "uid")
	_storePink.OrderID = field.NewString(tableName, "order_id")
	_storePink.OrderIDKey = field.NewInt64(tableName, "order_id_key")
	_storePink.TotalNum = field.NewInt64(tableName, "total_num")
	_storePink.TotalPrice = field.NewField(tableName, "total_price")
	_storePink.Cid = field.NewInt64(tableName, "cid")
	_storePink.Pid = field.NewInt64(tableName, "pid")
	_storePink.People = field.NewInt64(tableName, "people")
	_storePink.Price = field.NewField(tableName, "price")
	_storePink.AddTime = field.NewInt64(tableName, "add_time")
	_storePink.StopTime = field.NewInt64(tableName, "stop_time")
	_storePink.KID = field.NewInt64(tableName, "k_id")
	_storePink.IsTpl = field.NewInt64(tableName, "is_tpl")
	_storePink.IsRefund = field.NewInt64(tableName, "is_refund")
	_storePink.Status = field.NewInt64(tableName, "status")
	_storePink.IsVirtual = field.NewInt64(tableName, "is_virtual")
	_storePink.Nickname = field.NewString(tableName, "nickname")
	_storePink.Avatar = field.NewString(tableName, "avatar")
	_storePink.CreatedAt = field.NewInt64(tableName, "created_at")
	_storePink.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storePink.DeletedAt = field.NewField(tableName, "deleted_at")

	_storePink.fillFieldMap()

	return _storePink
}

// storePink 拼团表
type storePink struct {
	storePinkDo storePinkDo

	ALL        field.Asterisk
	ID         field.Int64  // 拼团ID
	UID        field.Int64  // 用户id
	OrderID    field.String // 订单id 生成
	OrderIDKey field.Int64  // 订单id  数据库
	TotalNum   field.Int64  // 购买商品个数
	TotalPrice field.Field  // 购买总金额
	Cid        field.Int64  // 拼团商品id
	Pid        field.Int64  // 商品id
	People     field.Int64  // 拼图总人数
	Price      field.Field  // 拼团商品单价
	AddTime    field.Int64  // 开始时间
	StopTime   field.Int64  // 结束时间
	KID        field.Int64  // 团长id 0为团长
	IsTpl      field.Int64  // 是否发送模板消息0未发送1已发送
	IsRefund   field.Int64  // 是否退款 0未退款 1已退款
	Status     field.Int64  // 状态1进行中2已完成3未完成
	IsVirtual  field.Int64  // 是否虚拟拼团
	Nickname   field.String // 用户昵称
	Avatar     field.String // 用户头像
	CreatedAt  field.Int64
	UpdatedAt  field.Int64
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (s storePink) Table(newTableName string) *storePink {
	s.storePinkDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storePink) As(alias string) *storePink {
	s.storePinkDo.DO = *(s.storePinkDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storePink) updateTableName(table string) *storePink {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UID = field.NewInt64(table, "uid")
	s.OrderID = field.NewString(table, "order_id")
	s.OrderIDKey = field.NewInt64(table, "order_id_key")
	s.TotalNum = field.NewInt64(table, "total_num")
	s.TotalPrice = field.NewField(table, "total_price")
	s.Cid = field.NewInt64(table, "cid")
	s.Pid = field.NewInt64(table, "pid")
	s.People = field.NewInt64(table, "people")
	s.Price = field.NewField(table, "price")
	s.AddTime = field.NewInt64(table, "add_time")
	s.StopTime = field.NewInt64(table, "stop_time")
	s.KID = field.NewInt64(table, "k_id")
	s.IsTpl = field.NewInt64(table, "is_tpl")
	s.IsRefund = field.NewInt64(table, "is_refund")
	s.Status = field.NewInt64(table, "status")
	s.IsVirtual = field.NewInt64(table, "is_virtual")
	s.Nickname = field.NewString(table, "nickname")
	s.Avatar = field.NewString(table, "avatar")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storePink) WithContext(ctx context.Context) IStorePinkDo {
	return s.storePinkDo.WithContext(ctx)
}

func (s storePink) TableName() string { return s.storePinkDo.TableName() }

func (s storePink) Alias() string { return s.storePinkDo.Alias() }

func (s storePink) Columns(cols ...field.Expr) gen.Columns { return s.storePinkDo.Columns(cols...) }

func (s *storePink) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storePink) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 22)
	s.fieldMap["id"] = s.ID
	s.fieldMap["uid"] = s.UID
	s.fieldMap["order_id"] = s.OrderID
	s.fieldMap["order_id_key"] = s.OrderIDKey
	s.fieldMap["total_num"] = s.TotalNum
	s.fieldMap["total_price"] = s.TotalPrice
	s.fieldMap["cid"] = s.Cid
	s.fieldMap["pid"] = s.Pid
	s.fieldMap["people"] = s.People
	s.fieldMap["price"] = s.Price
	s.fieldMap["add_time"] = s.AddTime
	s.fieldMap["stop_time"] = s.StopTime
	s.fieldMap["k_id"] = s.KID
	s.fieldMap["is_tpl"] = s.IsTpl
	s.fieldMap["is_refund"] = s.IsRefund
	s.fieldMap["status"] = s.Status
	s.fieldMap["is_virtual"] = s.IsVirtual
	s.fieldMap["nickname"] = s.Nickname
	s.fieldMap["avatar"] = s.Avatar
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storePink) clone(db *gorm.DB) storePink {
	s.storePinkDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storePink) replaceDB(db *gorm.DB) storePink {
	s.storePinkDo.ReplaceDB(db)
	return s
}

type storePinkDo struct{ gen.DO }

type IStorePinkDo interface {
	gen.SubQuery
	Debug() IStorePinkDo
	WithContext(ctx context.Context) IStorePinkDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStorePinkDo
	WriteDB() IStorePinkDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStorePinkDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStorePinkDo
	Not(conds ...gen.Condition) IStorePinkDo
	Or(conds ...gen.Condition) IStorePinkDo
	Select(conds ...field.Expr) IStorePinkDo
	Where(conds ...gen.Condition) IStorePinkDo
	Order(conds ...field.Expr) IStorePinkDo
	Distinct(cols ...field.Expr) IStorePinkDo
	Omit(cols ...field.Expr) IStorePinkDo
	Join(table schema.Tabler, on ...field.Expr) IStorePinkDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStorePinkDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStorePinkDo
	Group(cols ...field.Expr) IStorePinkDo
	Having(conds ...gen.Condition) IStorePinkDo
	Limit(limit int) IStorePinkDo
	Offset(offset int) IStorePinkDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStorePinkDo
	Unscoped() IStorePinkDo
	Create(values ...*model.StorePink) error
	CreateInBatches(values []*model.StorePink, batchSize int) error
	Save(values ...*model.StorePink) error
	First() (*model.StorePink, error)
	Take() (*model.StorePink, error)
	Last() (*model.StorePink, error)
	Find() ([]*model.StorePink, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StorePink, err error)
	FindInBatches(result *[]*model.StorePink, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StorePink) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStorePinkDo
	Assign(attrs ...field.AssignExpr) IStorePinkDo
	Joins(fields ...field.RelationField) IStorePinkDo
	Preload(fields ...field.RelationField) IStorePinkDo
	FirstOrInit() (*model.StorePink, error)
	FirstOrCreate() (*model.StorePink, error)
	FindByPage(offset int, limit int) (result []*model.StorePink, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStorePinkDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storePinkDo) Debug() IStorePinkDo {
	return s.withDO(s.DO.Debug())
}

func (s storePinkDo) WithContext(ctx context.Context) IStorePinkDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storePinkDo) ReadDB() IStorePinkDo {
	return s.Clauses(dbresolver.Read)
}

func (s storePinkDo) WriteDB() IStorePinkDo {
	return s.Clauses(dbresolver.Write)
}

func (s storePinkDo) Session(config *gorm.Session) IStorePinkDo {
	return s.withDO(s.DO.Session(config))
}

func (s storePinkDo) Clauses(conds ...clause.Expression) IStorePinkDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storePinkDo) Returning(value interface{}, columns ...string) IStorePinkDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storePinkDo) Not(conds ...gen.Condition) IStorePinkDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storePinkDo) Or(conds ...gen.Condition) IStorePinkDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storePinkDo) Select(conds ...field.Expr) IStorePinkDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storePinkDo) Where(conds ...gen.Condition) IStorePinkDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storePinkDo) Order(conds ...field.Expr) IStorePinkDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storePinkDo) Distinct(cols ...field.Expr) IStorePinkDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storePinkDo) Omit(cols ...field.Expr) IStorePinkDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storePinkDo) Join(table schema.Tabler, on ...field.Expr) IStorePinkDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storePinkDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStorePinkDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storePinkDo) RightJoin(table schema.Tabler, on ...field.Expr) IStorePinkDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storePinkDo) Group(cols ...field.Expr) IStorePinkDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storePinkDo) Having(conds ...gen.Condition) IStorePinkDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storePinkDo) Limit(limit int) IStorePinkDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storePinkDo) Offset(offset int) IStorePinkDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storePinkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStorePinkDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storePinkDo) Unscoped() IStorePinkDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storePinkDo) Create(values ...*model.StorePink) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storePinkDo) CreateInBatches(values []*model.StorePink, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storePinkDo) Save(values ...*model.StorePink) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storePinkDo) First() (*model.StorePink, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorePink), nil
	}
}

func (s storePinkDo) Take() (*model.StorePink, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorePink), nil
	}
}

func (s storePinkDo) Last() (*model.StorePink, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorePink), nil
	}
}

func (s storePinkDo) Find() ([]*model.StorePink, error) {
	result, err := s.DO.Find()
	return result.([]*model.StorePink), err
}

func (s storePinkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StorePink, err error) {
	buf := make([]*model.StorePink, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storePinkDo) FindInBatches(result *[]*model.StorePink, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storePinkDo) Attrs(attrs ...field.AssignExpr) IStorePinkDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storePinkDo) Assign(attrs ...field.AssignExpr) IStorePinkDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storePinkDo) Joins(fields ...field.RelationField) IStorePinkDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storePinkDo) Preload(fields ...field.RelationField) IStorePinkDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storePinkDo) FirstOrInit() (*model.StorePink, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorePink), nil
	}
}

func (s storePinkDo) FirstOrCreate() (*model.StorePink, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorePink), nil
	}
}

func (s storePinkDo) FindByPage(offset int, limit int) (result []*model.StorePink, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s storePinkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storePinkDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storePinkDo) Delete(models ...*model.StorePink) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storePinkDo) withDO(do gen.Dao) *storePinkDo {
	s.DO = *do.(*gen.DO)
	return s
}
