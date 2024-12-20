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

func newStoreCart(db *gorm.DB, opts ...gen.DOOption) storeCart {
	_storeCart := storeCart{}

	_storeCart.storeCartDo.UseDB(db, opts...)
	_storeCart.storeCartDo.UseModel(&model.StoreCart{})

	tableName := _storeCart.storeCartDo.TableName()
	_storeCart.ALL = field.NewAsterisk(tableName)
	_storeCart.ID = field.NewInt64(tableName, "id")
	_storeCart.UID = field.NewInt64(tableName, "uid")
	_storeCart.Type = field.NewString(tableName, "type")
	_storeCart.ProductID = field.NewInt64(tableName, "product_id")
	_storeCart.ProductAttrUnique = field.NewString(tableName, "product_attr_unique")
	_storeCart.CartNum = field.NewInt64(tableName, "cart_num")
	_storeCart.IsNew = field.NewInt64(tableName, "is_new")
	_storeCart.CombinationID = field.NewInt64(tableName, "combination_id")
	_storeCart.SeckillID = field.NewInt64(tableName, "seckill_id")
	_storeCart.BargainID = field.NewInt64(tableName, "bargain_id")
	_storeCart.Status = field.NewInt64(tableName, "status")
	_storeCart.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeCart.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeCart.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeCart.fillFieldMap()

	return _storeCart
}

// storeCart 购物车表
type storeCart struct {
	storeCartDo storeCartDo

	ALL               field.Asterisk
	ID                field.Int64  // 购物车表ID
	UID               field.Int64  // 用户ID
	Type              field.String // 类型
	ProductID         field.Int64  // 商品ID
	ProductAttrUnique field.String // 商品属性
	CartNum           field.Int64  // 商品数量
	IsNew             field.Int64  // 是否为立即购买
	CombinationID     field.Int64  // 拼团id
	SeckillID         field.Int64  // 秒杀商品ID
	BargainID         field.Int64  // 砍价id
	Status            field.Int64  // 购物车状态
	CreatedAt         field.Int64
	UpdatedAt         field.Int64
	DeletedAt         field.Field

	fieldMap map[string]field.Expr
}

func (s storeCart) Table(newTableName string) *storeCart {
	s.storeCartDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeCart) As(alias string) *storeCart {
	s.storeCartDo.DO = *(s.storeCartDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeCart) updateTableName(table string) *storeCart {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UID = field.NewInt64(table, "uid")
	s.Type = field.NewString(table, "type")
	s.ProductID = field.NewInt64(table, "product_id")
	s.ProductAttrUnique = field.NewString(table, "product_attr_unique")
	s.CartNum = field.NewInt64(table, "cart_num")
	s.IsNew = field.NewInt64(table, "is_new")
	s.CombinationID = field.NewInt64(table, "combination_id")
	s.SeckillID = field.NewInt64(table, "seckill_id")
	s.BargainID = field.NewInt64(table, "bargain_id")
	s.Status = field.NewInt64(table, "status")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeCart) WithContext(ctx context.Context) IStoreCartDo {
	return s.storeCartDo.WithContext(ctx)
}

func (s storeCart) TableName() string { return s.storeCartDo.TableName() }

func (s storeCart) Alias() string { return s.storeCartDo.Alias() }

func (s storeCart) Columns(cols ...field.Expr) gen.Columns { return s.storeCartDo.Columns(cols...) }

func (s *storeCart) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeCart) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["uid"] = s.UID
	s.fieldMap["type"] = s.Type
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["product_attr_unique"] = s.ProductAttrUnique
	s.fieldMap["cart_num"] = s.CartNum
	s.fieldMap["is_new"] = s.IsNew
	s.fieldMap["combination_id"] = s.CombinationID
	s.fieldMap["seckill_id"] = s.SeckillID
	s.fieldMap["bargain_id"] = s.BargainID
	s.fieldMap["status"] = s.Status
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeCart) clone(db *gorm.DB) storeCart {
	s.storeCartDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeCart) replaceDB(db *gorm.DB) storeCart {
	s.storeCartDo.ReplaceDB(db)
	return s
}

type storeCartDo struct{ gen.DO }

type IStoreCartDo interface {
	gen.SubQuery
	Debug() IStoreCartDo
	WithContext(ctx context.Context) IStoreCartDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreCartDo
	WriteDB() IStoreCartDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreCartDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreCartDo
	Not(conds ...gen.Condition) IStoreCartDo
	Or(conds ...gen.Condition) IStoreCartDo
	Select(conds ...field.Expr) IStoreCartDo
	Where(conds ...gen.Condition) IStoreCartDo
	Order(conds ...field.Expr) IStoreCartDo
	Distinct(cols ...field.Expr) IStoreCartDo
	Omit(cols ...field.Expr) IStoreCartDo
	Join(table schema.Tabler, on ...field.Expr) IStoreCartDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreCartDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreCartDo
	Group(cols ...field.Expr) IStoreCartDo
	Having(conds ...gen.Condition) IStoreCartDo
	Limit(limit int) IStoreCartDo
	Offset(offset int) IStoreCartDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreCartDo
	Unscoped() IStoreCartDo
	Create(values ...*model.StoreCart) error
	CreateInBatches(values []*model.StoreCart, batchSize int) error
	Save(values ...*model.StoreCart) error
	First() (*model.StoreCart, error)
	Take() (*model.StoreCart, error)
	Last() (*model.StoreCart, error)
	Find() ([]*model.StoreCart, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreCart, err error)
	FindInBatches(result *[]*model.StoreCart, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreCart) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreCartDo
	Assign(attrs ...field.AssignExpr) IStoreCartDo
	Joins(fields ...field.RelationField) IStoreCartDo
	Preload(fields ...field.RelationField) IStoreCartDo
	FirstOrInit() (*model.StoreCart, error)
	FirstOrCreate() (*model.StoreCart, error)
	FindByPage(offset int, limit int) (result []*model.StoreCart, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreCartDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeCartDo) Debug() IStoreCartDo {
	return s.withDO(s.DO.Debug())
}

func (s storeCartDo) WithContext(ctx context.Context) IStoreCartDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeCartDo) ReadDB() IStoreCartDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeCartDo) WriteDB() IStoreCartDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeCartDo) Session(config *gorm.Session) IStoreCartDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeCartDo) Clauses(conds ...clause.Expression) IStoreCartDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeCartDo) Returning(value interface{}, columns ...string) IStoreCartDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeCartDo) Not(conds ...gen.Condition) IStoreCartDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeCartDo) Or(conds ...gen.Condition) IStoreCartDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeCartDo) Select(conds ...field.Expr) IStoreCartDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeCartDo) Where(conds ...gen.Condition) IStoreCartDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeCartDo) Order(conds ...field.Expr) IStoreCartDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeCartDo) Distinct(cols ...field.Expr) IStoreCartDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeCartDo) Omit(cols ...field.Expr) IStoreCartDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeCartDo) Join(table schema.Tabler, on ...field.Expr) IStoreCartDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeCartDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreCartDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeCartDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreCartDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeCartDo) Group(cols ...field.Expr) IStoreCartDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeCartDo) Having(conds ...gen.Condition) IStoreCartDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeCartDo) Limit(limit int) IStoreCartDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeCartDo) Offset(offset int) IStoreCartDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeCartDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreCartDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeCartDo) Unscoped() IStoreCartDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeCartDo) Create(values ...*model.StoreCart) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeCartDo) CreateInBatches(values []*model.StoreCart, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeCartDo) Save(values ...*model.StoreCart) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeCartDo) First() (*model.StoreCart, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCart), nil
	}
}

func (s storeCartDo) Take() (*model.StoreCart, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCart), nil
	}
}

func (s storeCartDo) Last() (*model.StoreCart, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCart), nil
	}
}

func (s storeCartDo) Find() ([]*model.StoreCart, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreCart), err
}

func (s storeCartDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreCart, err error) {
	buf := make([]*model.StoreCart, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeCartDo) FindInBatches(result *[]*model.StoreCart, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeCartDo) Attrs(attrs ...field.AssignExpr) IStoreCartDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeCartDo) Assign(attrs ...field.AssignExpr) IStoreCartDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeCartDo) Joins(fields ...field.RelationField) IStoreCartDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeCartDo) Preload(fields ...field.RelationField) IStoreCartDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeCartDo) FirstOrInit() (*model.StoreCart, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCart), nil
	}
}

func (s storeCartDo) FirstOrCreate() (*model.StoreCart, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCart), nil
	}
}

func (s storeCartDo) FindByPage(offset int, limit int) (result []*model.StoreCart, count int64, err error) {
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

func (s storeCartDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeCartDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeCartDo) Delete(models ...*model.StoreCart) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeCartDo) withDO(do gen.Dao) *storeCartDo {
	s.DO = *do.(*gen.DO)
	return s
}
