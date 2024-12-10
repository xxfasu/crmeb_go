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

func newStoreBargain(db *gorm.DB, opts ...gen.DOOption) storeBargain {
	_storeBargain := storeBargain{}

	_storeBargain.storeBargainDo.UseDB(db, opts...)
	_storeBargain.storeBargainDo.UseModel(&model.StoreBargain{})

	tableName := _storeBargain.storeBargainDo.TableName()
	_storeBargain.ALL = field.NewAsterisk(tableName)
	_storeBargain.ID = field.NewInt64(tableName, "id")
	_storeBargain.ProductID = field.NewInt64(tableName, "product_id")
	_storeBargain.Title = field.NewString(tableName, "title")
	_storeBargain.Image = field.NewString(tableName, "image")
	_storeBargain.UnitName = field.NewString(tableName, "unit_name")
	_storeBargain.Stock = field.NewInt64(tableName, "stock")
	_storeBargain.Sales = field.NewInt64(tableName, "sales")
	_storeBargain.Images = field.NewString(tableName, "images")
	_storeBargain.StartTime = field.NewInt64(tableName, "start_time")
	_storeBargain.StopTime = field.NewInt64(tableName, "stop_time")
	_storeBargain.StoreName = field.NewString(tableName, "store_name")
	_storeBargain.Price = field.NewField(tableName, "price")
	_storeBargain.MinPrice = field.NewField(tableName, "min_price")
	_storeBargain.Num = field.NewInt64(tableName, "num")
	_storeBargain.BargainMaxPrice = field.NewField(tableName, "bargain_max_price")
	_storeBargain.BargainMinPrice = field.NewField(tableName, "bargain_min_price")
	_storeBargain.BargainNum = field.NewInt64(tableName, "bargain_num")
	_storeBargain.Status = field.NewInt64(tableName, "status")
	_storeBargain.GiveIntegral = field.NewInt64(tableName, "give_integral")
	_storeBargain.Info = field.NewString(tableName, "info")
	_storeBargain.Cost = field.NewField(tableName, "cost")
	_storeBargain.Sort = field.NewInt64(tableName, "sort")
	_storeBargain.IsHot = field.NewInt64(tableName, "is_hot")
	_storeBargain.IsDel = field.NewInt64(tableName, "is_del")
	_storeBargain.AddTime = field.NewInt64(tableName, "add_time")
	_storeBargain.IsPostage = field.NewInt64(tableName, "is_postage")
	_storeBargain.Postage = field.NewField(tableName, "postage")
	_storeBargain.Rule = field.NewString(tableName, "rule")
	_storeBargain.Look = field.NewInt64(tableName, "look")
	_storeBargain.Share = field.NewInt64(tableName, "share")
	_storeBargain.TempID = field.NewInt64(tableName, "temp_id")
	_storeBargain.Weight = field.NewField(tableName, "weight")
	_storeBargain.Volume = field.NewField(tableName, "volume")
	_storeBargain.Quota = field.NewInt64(tableName, "quota")
	_storeBargain.QuotaShow = field.NewInt64(tableName, "quota_show")
	_storeBargain.PeopleNum = field.NewInt64(tableName, "people_num")
	_storeBargain.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeBargain.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeBargain.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeBargain.fillFieldMap()

	return _storeBargain
}

// storeBargain 砍价表
type storeBargain struct {
	storeBargainDo storeBargainDo

	ALL             field.Asterisk
	ID              field.Int64  // 砍价商品ID
	ProductID       field.Int64  // 关联商品ID
	Title           field.String // 砍价活动名称
	Image           field.String // 砍价活动图片
	UnitName        field.String // 单位名称
	Stock           field.Int64  // 库存
	Sales           field.Int64  // 销量
	Images          field.String // 砍价商品轮播图
	StartTime       field.Int64  // 砍价开启时间
	StopTime        field.Int64  // 砍价结束时间
	StoreName       field.String // 砍价商品名称
	Price           field.Field  // 砍价金额
	MinPrice        field.Field  // 砍价商品最低价
	Num             field.Int64  // 购买数量限制——单个活动每个用户发起砍价次数限制
	BargainMaxPrice field.Field  // 用户每次砍价的最大金额
	BargainMinPrice field.Field  // 用户每次砍价的最小金额
	BargainNum      field.Int64  // 帮砍次数——单个商品用户可以帮砍的次数
	Status          field.Int64  // 砍价状态 0(到砍价时间不自动开启)  1(到砍价时间自动开启时间)
	GiveIntegral    field.Int64  // 反多少积分
	Info            field.String // 砍价活动简介
	Cost            field.Field  // 成本价
	Sort            field.Int64  // 排序
	IsHot           field.Int64  // 是否推荐0不推荐1推荐
	IsDel           field.Int64  // 是否删除 0未删除 1删除
	AddTime         field.Int64  // 添加时间
	IsPostage       field.Int64  // 是否包邮 0不包邮 1包邮
	Postage         field.Field  // 邮费
	Rule            field.String // 砍价规则
	Look            field.Int64  // 砍价商品浏览量
	Share           field.Int64  // 砍价商品分享量
	TempID          field.Int64  // 运费模板ID
	Weight          field.Field  // 重量
	Volume          field.Field  // 体积
	Quota           field.Int64  // 限购总数
	QuotaShow       field.Int64  // 限量总数显示
	PeopleNum       field.Int64  // 砍价人数——需要多少人砍价成功
	CreatedAt       field.Int64
	UpdatedAt       field.Int64
	DeletedAt       field.Field

	fieldMap map[string]field.Expr
}

func (s storeBargain) Table(newTableName string) *storeBargain {
	s.storeBargainDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeBargain) As(alias string) *storeBargain {
	s.storeBargainDo.DO = *(s.storeBargainDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeBargain) updateTableName(table string) *storeBargain {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.ProductID = field.NewInt64(table, "product_id")
	s.Title = field.NewString(table, "title")
	s.Image = field.NewString(table, "image")
	s.UnitName = field.NewString(table, "unit_name")
	s.Stock = field.NewInt64(table, "stock")
	s.Sales = field.NewInt64(table, "sales")
	s.Images = field.NewString(table, "images")
	s.StartTime = field.NewInt64(table, "start_time")
	s.StopTime = field.NewInt64(table, "stop_time")
	s.StoreName = field.NewString(table, "store_name")
	s.Price = field.NewField(table, "price")
	s.MinPrice = field.NewField(table, "min_price")
	s.Num = field.NewInt64(table, "num")
	s.BargainMaxPrice = field.NewField(table, "bargain_max_price")
	s.BargainMinPrice = field.NewField(table, "bargain_min_price")
	s.BargainNum = field.NewInt64(table, "bargain_num")
	s.Status = field.NewInt64(table, "status")
	s.GiveIntegral = field.NewInt64(table, "give_integral")
	s.Info = field.NewString(table, "info")
	s.Cost = field.NewField(table, "cost")
	s.Sort = field.NewInt64(table, "sort")
	s.IsHot = field.NewInt64(table, "is_hot")
	s.IsDel = field.NewInt64(table, "is_del")
	s.AddTime = field.NewInt64(table, "add_time")
	s.IsPostage = field.NewInt64(table, "is_postage")
	s.Postage = field.NewField(table, "postage")
	s.Rule = field.NewString(table, "rule")
	s.Look = field.NewInt64(table, "look")
	s.Share = field.NewInt64(table, "share")
	s.TempID = field.NewInt64(table, "temp_id")
	s.Weight = field.NewField(table, "weight")
	s.Volume = field.NewField(table, "volume")
	s.Quota = field.NewInt64(table, "quota")
	s.QuotaShow = field.NewInt64(table, "quota_show")
	s.PeopleNum = field.NewInt64(table, "people_num")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeBargain) WithContext(ctx context.Context) IStoreBargainDo {
	return s.storeBargainDo.WithContext(ctx)
}

func (s storeBargain) TableName() string { return s.storeBargainDo.TableName() }

func (s storeBargain) Alias() string { return s.storeBargainDo.Alias() }

func (s storeBargain) Columns(cols ...field.Expr) gen.Columns {
	return s.storeBargainDo.Columns(cols...)
}

func (s *storeBargain) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeBargain) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 39)
	s.fieldMap["id"] = s.ID
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["title"] = s.Title
	s.fieldMap["image"] = s.Image
	s.fieldMap["unit_name"] = s.UnitName
	s.fieldMap["stock"] = s.Stock
	s.fieldMap["sales"] = s.Sales
	s.fieldMap["images"] = s.Images
	s.fieldMap["start_time"] = s.StartTime
	s.fieldMap["stop_time"] = s.StopTime
	s.fieldMap["store_name"] = s.StoreName
	s.fieldMap["price"] = s.Price
	s.fieldMap["min_price"] = s.MinPrice
	s.fieldMap["num"] = s.Num
	s.fieldMap["bargain_max_price"] = s.BargainMaxPrice
	s.fieldMap["bargain_min_price"] = s.BargainMinPrice
	s.fieldMap["bargain_num"] = s.BargainNum
	s.fieldMap["status"] = s.Status
	s.fieldMap["give_integral"] = s.GiveIntegral
	s.fieldMap["info"] = s.Info
	s.fieldMap["cost"] = s.Cost
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["is_hot"] = s.IsHot
	s.fieldMap["is_del"] = s.IsDel
	s.fieldMap["add_time"] = s.AddTime
	s.fieldMap["is_postage"] = s.IsPostage
	s.fieldMap["postage"] = s.Postage
	s.fieldMap["rule"] = s.Rule
	s.fieldMap["look"] = s.Look
	s.fieldMap["share"] = s.Share
	s.fieldMap["temp_id"] = s.TempID
	s.fieldMap["weight"] = s.Weight
	s.fieldMap["volume"] = s.Volume
	s.fieldMap["quota"] = s.Quota
	s.fieldMap["quota_show"] = s.QuotaShow
	s.fieldMap["people_num"] = s.PeopleNum
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeBargain) clone(db *gorm.DB) storeBargain {
	s.storeBargainDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeBargain) replaceDB(db *gorm.DB) storeBargain {
	s.storeBargainDo.ReplaceDB(db)
	return s
}

type storeBargainDo struct{ gen.DO }

type IStoreBargainDo interface {
	gen.SubQuery
	Debug() IStoreBargainDo
	WithContext(ctx context.Context) IStoreBargainDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreBargainDo
	WriteDB() IStoreBargainDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreBargainDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreBargainDo
	Not(conds ...gen.Condition) IStoreBargainDo
	Or(conds ...gen.Condition) IStoreBargainDo
	Select(conds ...field.Expr) IStoreBargainDo
	Where(conds ...gen.Condition) IStoreBargainDo
	Order(conds ...field.Expr) IStoreBargainDo
	Distinct(cols ...field.Expr) IStoreBargainDo
	Omit(cols ...field.Expr) IStoreBargainDo
	Join(table schema.Tabler, on ...field.Expr) IStoreBargainDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreBargainDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreBargainDo
	Group(cols ...field.Expr) IStoreBargainDo
	Having(conds ...gen.Condition) IStoreBargainDo
	Limit(limit int) IStoreBargainDo
	Offset(offset int) IStoreBargainDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreBargainDo
	Unscoped() IStoreBargainDo
	Create(values ...*model.StoreBargain) error
	CreateInBatches(values []*model.StoreBargain, batchSize int) error
	Save(values ...*model.StoreBargain) error
	First() (*model.StoreBargain, error)
	Take() (*model.StoreBargain, error)
	Last() (*model.StoreBargain, error)
	Find() ([]*model.StoreBargain, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreBargain, err error)
	FindInBatches(result *[]*model.StoreBargain, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreBargain) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreBargainDo
	Assign(attrs ...field.AssignExpr) IStoreBargainDo
	Joins(fields ...field.RelationField) IStoreBargainDo
	Preload(fields ...field.RelationField) IStoreBargainDo
	FirstOrInit() (*model.StoreBargain, error)
	FirstOrCreate() (*model.StoreBargain, error)
	FindByPage(offset int, limit int) (result []*model.StoreBargain, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreBargainDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeBargainDo) Debug() IStoreBargainDo {
	return s.withDO(s.DO.Debug())
}

func (s storeBargainDo) WithContext(ctx context.Context) IStoreBargainDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeBargainDo) ReadDB() IStoreBargainDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeBargainDo) WriteDB() IStoreBargainDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeBargainDo) Session(config *gorm.Session) IStoreBargainDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeBargainDo) Clauses(conds ...clause.Expression) IStoreBargainDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeBargainDo) Returning(value interface{}, columns ...string) IStoreBargainDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeBargainDo) Not(conds ...gen.Condition) IStoreBargainDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeBargainDo) Or(conds ...gen.Condition) IStoreBargainDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeBargainDo) Select(conds ...field.Expr) IStoreBargainDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeBargainDo) Where(conds ...gen.Condition) IStoreBargainDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeBargainDo) Order(conds ...field.Expr) IStoreBargainDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeBargainDo) Distinct(cols ...field.Expr) IStoreBargainDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeBargainDo) Omit(cols ...field.Expr) IStoreBargainDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeBargainDo) Join(table schema.Tabler, on ...field.Expr) IStoreBargainDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeBargainDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreBargainDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeBargainDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreBargainDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeBargainDo) Group(cols ...field.Expr) IStoreBargainDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeBargainDo) Having(conds ...gen.Condition) IStoreBargainDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeBargainDo) Limit(limit int) IStoreBargainDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeBargainDo) Offset(offset int) IStoreBargainDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeBargainDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreBargainDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeBargainDo) Unscoped() IStoreBargainDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeBargainDo) Create(values ...*model.StoreBargain) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeBargainDo) CreateInBatches(values []*model.StoreBargain, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeBargainDo) Save(values ...*model.StoreBargain) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeBargainDo) First() (*model.StoreBargain, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargain), nil
	}
}

func (s storeBargainDo) Take() (*model.StoreBargain, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargain), nil
	}
}

func (s storeBargainDo) Last() (*model.StoreBargain, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargain), nil
	}
}

func (s storeBargainDo) Find() ([]*model.StoreBargain, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreBargain), err
}

func (s storeBargainDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreBargain, err error) {
	buf := make([]*model.StoreBargain, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeBargainDo) FindInBatches(result *[]*model.StoreBargain, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeBargainDo) Attrs(attrs ...field.AssignExpr) IStoreBargainDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeBargainDo) Assign(attrs ...field.AssignExpr) IStoreBargainDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeBargainDo) Joins(fields ...field.RelationField) IStoreBargainDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeBargainDo) Preload(fields ...field.RelationField) IStoreBargainDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeBargainDo) FirstOrInit() (*model.StoreBargain, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargain), nil
	}
}

func (s storeBargainDo) FirstOrCreate() (*model.StoreBargain, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargain), nil
	}
}

func (s storeBargainDo) FindByPage(offset int, limit int) (result []*model.StoreBargain, count int64, err error) {
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

func (s storeBargainDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeBargainDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeBargainDo) Delete(models ...*model.StoreBargain) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeBargainDo) withDO(do gen.Dao) *storeBargainDo {
	s.DO = *do.(*gen.DO)
	return s
}
