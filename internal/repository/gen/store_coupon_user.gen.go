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

func newStoreCouponUser(db *gorm.DB, opts ...gen.DOOption) storeCouponUser {
	_storeCouponUser := storeCouponUser{}

	_storeCouponUser.storeCouponUserDo.UseDB(db, opts...)
	_storeCouponUser.storeCouponUserDo.UseModel(&model.StoreCouponUser{})

	tableName := _storeCouponUser.storeCouponUserDo.TableName()
	_storeCouponUser.ALL = field.NewAsterisk(tableName)
	_storeCouponUser.ID = field.NewInt64(tableName, "id")
	_storeCouponUser.CouponID = field.NewInt64(tableName, "coupon_id")
	_storeCouponUser.Cid = field.NewInt64(tableName, "cid")
	_storeCouponUser.UID = field.NewInt64(tableName, "uid")
	_storeCouponUser.Name = field.NewString(tableName, "name")
	_storeCouponUser.Money = field.NewField(tableName, "money")
	_storeCouponUser.MinPrice = field.NewField(tableName, "min_price")
	_storeCouponUser.Type = field.NewString(tableName, "type")
	_storeCouponUser.Status = field.NewBool(tableName, "status")
	_storeCouponUser.StartTime = field.NewInt64(tableName, "start_time")
	_storeCouponUser.EndTime = field.NewInt64(tableName, "end_time")
	_storeCouponUser.UseTime = field.NewInt64(tableName, "use_time")
	_storeCouponUser.UseType = field.NewBool(tableName, "use_type")
	_storeCouponUser.PrimaryKey = field.NewString(tableName, "primary_key")
	_storeCouponUser.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeCouponUser.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeCouponUser.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeCouponUser.fillFieldMap()

	return _storeCouponUser
}

// storeCouponUser 优惠券记录表
type storeCouponUser struct {
	storeCouponUserDo storeCouponUserDo

	ALL        field.Asterisk
	ID         field.Int64  // id
	CouponID   field.Int64  // 优惠券发布id
	Cid        field.Int64  // 兑换的项目id
	UID        field.Int64  // 领取人id
	Name       field.String // 优惠券名称
	Money      field.Field  // 优惠券的面值
	MinPrice   field.Field  // 最低消费多少金额可用优惠券
	Type       field.String // 获取方式，send后台发放, 用户领取 get
	Status     field.Bool   // 状态（0：未使用，1：已使用, 2:已失效）
	StartTime  field.Int64
	EndTime    field.Int64
	UseTime    field.Int64
	UseType    field.Bool   // 使用类型 1 全场通用, 2 商品券, 3 品类券
	PrimaryKey field.String // 所属商品id / 分类id
	CreatedAt  field.Int64
	UpdatedAt  field.Int64
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (s storeCouponUser) Table(newTableName string) *storeCouponUser {
	s.storeCouponUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeCouponUser) As(alias string) *storeCouponUser {
	s.storeCouponUserDo.DO = *(s.storeCouponUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeCouponUser) updateTableName(table string) *storeCouponUser {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.CouponID = field.NewInt64(table, "coupon_id")
	s.Cid = field.NewInt64(table, "cid")
	s.UID = field.NewInt64(table, "uid")
	s.Name = field.NewString(table, "name")
	s.Money = field.NewField(table, "money")
	s.MinPrice = field.NewField(table, "min_price")
	s.Type = field.NewString(table, "type")
	s.Status = field.NewBool(table, "status")
	s.StartTime = field.NewInt64(table, "start_time")
	s.EndTime = field.NewInt64(table, "end_time")
	s.UseTime = field.NewInt64(table, "use_time")
	s.UseType = field.NewBool(table, "use_type")
	s.PrimaryKey = field.NewString(table, "primary_key")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeCouponUser) WithContext(ctx context.Context) IStoreCouponUserDo {
	return s.storeCouponUserDo.WithContext(ctx)
}

func (s storeCouponUser) TableName() string { return s.storeCouponUserDo.TableName() }

func (s storeCouponUser) Alias() string { return s.storeCouponUserDo.Alias() }

func (s storeCouponUser) Columns(cols ...field.Expr) gen.Columns {
	return s.storeCouponUserDo.Columns(cols...)
}

func (s *storeCouponUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeCouponUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 17)
	s.fieldMap["id"] = s.ID
	s.fieldMap["coupon_id"] = s.CouponID
	s.fieldMap["cid"] = s.Cid
	s.fieldMap["uid"] = s.UID
	s.fieldMap["name"] = s.Name
	s.fieldMap["money"] = s.Money
	s.fieldMap["min_price"] = s.MinPrice
	s.fieldMap["type"] = s.Type
	s.fieldMap["status"] = s.Status
	s.fieldMap["start_time"] = s.StartTime
	s.fieldMap["end_time"] = s.EndTime
	s.fieldMap["use_time"] = s.UseTime
	s.fieldMap["use_type"] = s.UseType
	s.fieldMap["primary_key"] = s.PrimaryKey
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeCouponUser) clone(db *gorm.DB) storeCouponUser {
	s.storeCouponUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeCouponUser) replaceDB(db *gorm.DB) storeCouponUser {
	s.storeCouponUserDo.ReplaceDB(db)
	return s
}

type storeCouponUserDo struct{ gen.DO }

type IStoreCouponUserDo interface {
	gen.SubQuery
	Debug() IStoreCouponUserDo
	WithContext(ctx context.Context) IStoreCouponUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreCouponUserDo
	WriteDB() IStoreCouponUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreCouponUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreCouponUserDo
	Not(conds ...gen.Condition) IStoreCouponUserDo
	Or(conds ...gen.Condition) IStoreCouponUserDo
	Select(conds ...field.Expr) IStoreCouponUserDo
	Where(conds ...gen.Condition) IStoreCouponUserDo
	Order(conds ...field.Expr) IStoreCouponUserDo
	Distinct(cols ...field.Expr) IStoreCouponUserDo
	Omit(cols ...field.Expr) IStoreCouponUserDo
	Join(table schema.Tabler, on ...field.Expr) IStoreCouponUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreCouponUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreCouponUserDo
	Group(cols ...field.Expr) IStoreCouponUserDo
	Having(conds ...gen.Condition) IStoreCouponUserDo
	Limit(limit int) IStoreCouponUserDo
	Offset(offset int) IStoreCouponUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreCouponUserDo
	Unscoped() IStoreCouponUserDo
	Create(values ...*model.StoreCouponUser) error
	CreateInBatches(values []*model.StoreCouponUser, batchSize int) error
	Save(values ...*model.StoreCouponUser) error
	First() (*model.StoreCouponUser, error)
	Take() (*model.StoreCouponUser, error)
	Last() (*model.StoreCouponUser, error)
	Find() ([]*model.StoreCouponUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreCouponUser, err error)
	FindInBatches(result *[]*model.StoreCouponUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreCouponUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreCouponUserDo
	Assign(attrs ...field.AssignExpr) IStoreCouponUserDo
	Joins(fields ...field.RelationField) IStoreCouponUserDo
	Preload(fields ...field.RelationField) IStoreCouponUserDo
	FirstOrInit() (*model.StoreCouponUser, error)
	FirstOrCreate() (*model.StoreCouponUser, error)
	FindByPage(offset int, limit int) (result []*model.StoreCouponUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreCouponUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeCouponUserDo) Debug() IStoreCouponUserDo {
	return s.withDO(s.DO.Debug())
}

func (s storeCouponUserDo) WithContext(ctx context.Context) IStoreCouponUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeCouponUserDo) ReadDB() IStoreCouponUserDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeCouponUserDo) WriteDB() IStoreCouponUserDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeCouponUserDo) Session(config *gorm.Session) IStoreCouponUserDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeCouponUserDo) Clauses(conds ...clause.Expression) IStoreCouponUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeCouponUserDo) Returning(value interface{}, columns ...string) IStoreCouponUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeCouponUserDo) Not(conds ...gen.Condition) IStoreCouponUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeCouponUserDo) Or(conds ...gen.Condition) IStoreCouponUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeCouponUserDo) Select(conds ...field.Expr) IStoreCouponUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeCouponUserDo) Where(conds ...gen.Condition) IStoreCouponUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeCouponUserDo) Order(conds ...field.Expr) IStoreCouponUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeCouponUserDo) Distinct(cols ...field.Expr) IStoreCouponUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeCouponUserDo) Omit(cols ...field.Expr) IStoreCouponUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeCouponUserDo) Join(table schema.Tabler, on ...field.Expr) IStoreCouponUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeCouponUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreCouponUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeCouponUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreCouponUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeCouponUserDo) Group(cols ...field.Expr) IStoreCouponUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeCouponUserDo) Having(conds ...gen.Condition) IStoreCouponUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeCouponUserDo) Limit(limit int) IStoreCouponUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeCouponUserDo) Offset(offset int) IStoreCouponUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeCouponUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreCouponUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeCouponUserDo) Unscoped() IStoreCouponUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeCouponUserDo) Create(values ...*model.StoreCouponUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeCouponUserDo) CreateInBatches(values []*model.StoreCouponUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeCouponUserDo) Save(values ...*model.StoreCouponUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeCouponUserDo) First() (*model.StoreCouponUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCouponUser), nil
	}
}

func (s storeCouponUserDo) Take() (*model.StoreCouponUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCouponUser), nil
	}
}

func (s storeCouponUserDo) Last() (*model.StoreCouponUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCouponUser), nil
	}
}

func (s storeCouponUserDo) Find() ([]*model.StoreCouponUser, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreCouponUser), err
}

func (s storeCouponUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreCouponUser, err error) {
	buf := make([]*model.StoreCouponUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeCouponUserDo) FindInBatches(result *[]*model.StoreCouponUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeCouponUserDo) Attrs(attrs ...field.AssignExpr) IStoreCouponUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeCouponUserDo) Assign(attrs ...field.AssignExpr) IStoreCouponUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeCouponUserDo) Joins(fields ...field.RelationField) IStoreCouponUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeCouponUserDo) Preload(fields ...field.RelationField) IStoreCouponUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeCouponUserDo) FirstOrInit() (*model.StoreCouponUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCouponUser), nil
	}
}

func (s storeCouponUserDo) FirstOrCreate() (*model.StoreCouponUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCouponUser), nil
	}
}

func (s storeCouponUserDo) FindByPage(offset int, limit int) (result []*model.StoreCouponUser, count int64, err error) {
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

func (s storeCouponUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeCouponUserDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeCouponUserDo) Delete(models ...*model.StoreCouponUser) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeCouponUserDo) withDO(do gen.Dao) *storeCouponUserDo {
	s.DO = *do.(*gen.DO)
	return s
}
