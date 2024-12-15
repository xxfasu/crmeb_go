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

func newStoreProductReply(db *gorm.DB, opts ...gen.DOOption) storeProductReply {
	_storeProductReply := storeProductReply{}

	_storeProductReply.storeProductReplyDo.UseDB(db, opts...)
	_storeProductReply.storeProductReplyDo.UseModel(&model.StoreProductReply{})

	tableName := _storeProductReply.storeProductReplyDo.TableName()
	_storeProductReply.ALL = field.NewAsterisk(tableName)
	_storeProductReply.ID = field.NewInt64(tableName, "id")
	_storeProductReply.UID = field.NewInt64(tableName, "uid")
	_storeProductReply.Oid = field.NewInt64(tableName, "oid")
	_storeProductReply.Unique = field.NewString(tableName, "unique")
	_storeProductReply.ProductID = field.NewInt64(tableName, "product_id")
	_storeProductReply.ReplyType = field.NewString(tableName, "reply_type")
	_storeProductReply.ProductScore = field.NewInt64(tableName, "product_score")
	_storeProductReply.ServiceScore = field.NewInt64(tableName, "service_score")
	_storeProductReply.Comment = field.NewString(tableName, "comment")
	_storeProductReply.Pics = field.NewString(tableName, "pics")
	_storeProductReply.MerchantReplyContent = field.NewString(tableName, "merchant_reply_content")
	_storeProductReply.MerchantReplyTime = field.NewInt64(tableName, "merchant_reply_time")
	_storeProductReply.IsReply = field.NewInt64(tableName, "is_reply")
	_storeProductReply.Nickname = field.NewString(tableName, "nickname")
	_storeProductReply.Avatar = field.NewString(tableName, "avatar")
	_storeProductReply.Sku = field.NewString(tableName, "sku")
	_storeProductReply.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeProductReply.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeProductReply.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeProductReply.fillFieldMap()

	return _storeProductReply
}

// storeProductReply 评论表
type storeProductReply struct {
	storeProductReplyDo storeProductReplyDo

	ALL                  field.Asterisk
	ID                   field.Int64  // 评论ID
	UID                  field.Int64  // 用户ID
	Oid                  field.Int64  // 订单ID
	Unique               field.String // 商品唯一id
	ProductID            field.Int64  // 商品id
	ReplyType            field.String // 某种商品类型(普通商品、秒杀商品）
	ProductScore         field.Int64  // 商品分数
	ServiceScore         field.Int64  // 服务分数
	Comment              field.String // 评论内容
	Pics                 field.String // 评论图片
	MerchantReplyContent field.String // 管理员回复内容
	MerchantReplyTime    field.Int64  // 管理员回复时间
	IsReply              field.Int64  // 0未回复1已回复
	Nickname             field.String // 用户名称
	Avatar               field.String // 用户头像
	Sku                  field.String // 商品规格属性值,多个,号隔开
	CreatedAt            field.Int64
	UpdatedAt            field.Int64
	DeletedAt            field.Field

	fieldMap map[string]field.Expr
}

func (s storeProductReply) Table(newTableName string) *storeProductReply {
	s.storeProductReplyDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeProductReply) As(alias string) *storeProductReply {
	s.storeProductReplyDo.DO = *(s.storeProductReplyDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeProductReply) updateTableName(table string) *storeProductReply {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UID = field.NewInt64(table, "uid")
	s.Oid = field.NewInt64(table, "oid")
	s.Unique = field.NewString(table, "unique")
	s.ProductID = field.NewInt64(table, "product_id")
	s.ReplyType = field.NewString(table, "reply_type")
	s.ProductScore = field.NewInt64(table, "product_score")
	s.ServiceScore = field.NewInt64(table, "service_score")
	s.Comment = field.NewString(table, "comment")
	s.Pics = field.NewString(table, "pics")
	s.MerchantReplyContent = field.NewString(table, "merchant_reply_content")
	s.MerchantReplyTime = field.NewInt64(table, "merchant_reply_time")
	s.IsReply = field.NewInt64(table, "is_reply")
	s.Nickname = field.NewString(table, "nickname")
	s.Avatar = field.NewString(table, "avatar")
	s.Sku = field.NewString(table, "sku")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeProductReply) WithContext(ctx context.Context) IStoreProductReplyDo {
	return s.storeProductReplyDo.WithContext(ctx)
}

func (s storeProductReply) TableName() string { return s.storeProductReplyDo.TableName() }

func (s storeProductReply) Alias() string { return s.storeProductReplyDo.Alias() }

func (s storeProductReply) Columns(cols ...field.Expr) gen.Columns {
	return s.storeProductReplyDo.Columns(cols...)
}

func (s *storeProductReply) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeProductReply) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 19)
	s.fieldMap["id"] = s.ID
	s.fieldMap["uid"] = s.UID
	s.fieldMap["oid"] = s.Oid
	s.fieldMap["unique"] = s.Unique
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["reply_type"] = s.ReplyType
	s.fieldMap["product_score"] = s.ProductScore
	s.fieldMap["service_score"] = s.ServiceScore
	s.fieldMap["comment"] = s.Comment
	s.fieldMap["pics"] = s.Pics
	s.fieldMap["merchant_reply_content"] = s.MerchantReplyContent
	s.fieldMap["merchant_reply_time"] = s.MerchantReplyTime
	s.fieldMap["is_reply"] = s.IsReply
	s.fieldMap["nickname"] = s.Nickname
	s.fieldMap["avatar"] = s.Avatar
	s.fieldMap["sku"] = s.Sku
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeProductReply) clone(db *gorm.DB) storeProductReply {
	s.storeProductReplyDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeProductReply) replaceDB(db *gorm.DB) storeProductReply {
	s.storeProductReplyDo.ReplaceDB(db)
	return s
}

type storeProductReplyDo struct{ gen.DO }

type IStoreProductReplyDo interface {
	gen.SubQuery
	Debug() IStoreProductReplyDo
	WithContext(ctx context.Context) IStoreProductReplyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreProductReplyDo
	WriteDB() IStoreProductReplyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreProductReplyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreProductReplyDo
	Not(conds ...gen.Condition) IStoreProductReplyDo
	Or(conds ...gen.Condition) IStoreProductReplyDo
	Select(conds ...field.Expr) IStoreProductReplyDo
	Where(conds ...gen.Condition) IStoreProductReplyDo
	Order(conds ...field.Expr) IStoreProductReplyDo
	Distinct(cols ...field.Expr) IStoreProductReplyDo
	Omit(cols ...field.Expr) IStoreProductReplyDo
	Join(table schema.Tabler, on ...field.Expr) IStoreProductReplyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductReplyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductReplyDo
	Group(cols ...field.Expr) IStoreProductReplyDo
	Having(conds ...gen.Condition) IStoreProductReplyDo
	Limit(limit int) IStoreProductReplyDo
	Offset(offset int) IStoreProductReplyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductReplyDo
	Unscoped() IStoreProductReplyDo
	Create(values ...*model.StoreProductReply) error
	CreateInBatches(values []*model.StoreProductReply, batchSize int) error
	Save(values ...*model.StoreProductReply) error
	First() (*model.StoreProductReply, error)
	Take() (*model.StoreProductReply, error)
	Last() (*model.StoreProductReply, error)
	Find() ([]*model.StoreProductReply, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductReply, err error)
	FindInBatches(result *[]*model.StoreProductReply, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreProductReply) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreProductReplyDo
	Assign(attrs ...field.AssignExpr) IStoreProductReplyDo
	Joins(fields ...field.RelationField) IStoreProductReplyDo
	Preload(fields ...field.RelationField) IStoreProductReplyDo
	FirstOrInit() (*model.StoreProductReply, error)
	FirstOrCreate() (*model.StoreProductReply, error)
	FindByPage(offset int, limit int) (result []*model.StoreProductReply, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreProductReplyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeProductReplyDo) Debug() IStoreProductReplyDo {
	return s.withDO(s.DO.Debug())
}

func (s storeProductReplyDo) WithContext(ctx context.Context) IStoreProductReplyDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeProductReplyDo) ReadDB() IStoreProductReplyDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeProductReplyDo) WriteDB() IStoreProductReplyDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeProductReplyDo) Session(config *gorm.Session) IStoreProductReplyDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeProductReplyDo) Clauses(conds ...clause.Expression) IStoreProductReplyDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeProductReplyDo) Returning(value interface{}, columns ...string) IStoreProductReplyDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeProductReplyDo) Not(conds ...gen.Condition) IStoreProductReplyDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeProductReplyDo) Or(conds ...gen.Condition) IStoreProductReplyDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeProductReplyDo) Select(conds ...field.Expr) IStoreProductReplyDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeProductReplyDo) Where(conds ...gen.Condition) IStoreProductReplyDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeProductReplyDo) Order(conds ...field.Expr) IStoreProductReplyDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeProductReplyDo) Distinct(cols ...field.Expr) IStoreProductReplyDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeProductReplyDo) Omit(cols ...field.Expr) IStoreProductReplyDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeProductReplyDo) Join(table schema.Tabler, on ...field.Expr) IStoreProductReplyDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeProductReplyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductReplyDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeProductReplyDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductReplyDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeProductReplyDo) Group(cols ...field.Expr) IStoreProductReplyDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeProductReplyDo) Having(conds ...gen.Condition) IStoreProductReplyDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeProductReplyDo) Limit(limit int) IStoreProductReplyDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeProductReplyDo) Offset(offset int) IStoreProductReplyDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeProductReplyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductReplyDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeProductReplyDo) Unscoped() IStoreProductReplyDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeProductReplyDo) Create(values ...*model.StoreProductReply) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeProductReplyDo) CreateInBatches(values []*model.StoreProductReply, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeProductReplyDo) Save(values ...*model.StoreProductReply) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeProductReplyDo) First() (*model.StoreProductReply, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductReply), nil
	}
}

func (s storeProductReplyDo) Take() (*model.StoreProductReply, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductReply), nil
	}
}

func (s storeProductReplyDo) Last() (*model.StoreProductReply, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductReply), nil
	}
}

func (s storeProductReplyDo) Find() ([]*model.StoreProductReply, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreProductReply), err
}

func (s storeProductReplyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductReply, err error) {
	buf := make([]*model.StoreProductReply, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeProductReplyDo) FindInBatches(result *[]*model.StoreProductReply, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeProductReplyDo) Attrs(attrs ...field.AssignExpr) IStoreProductReplyDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeProductReplyDo) Assign(attrs ...field.AssignExpr) IStoreProductReplyDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeProductReplyDo) Joins(fields ...field.RelationField) IStoreProductReplyDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeProductReplyDo) Preload(fields ...field.RelationField) IStoreProductReplyDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeProductReplyDo) FirstOrInit() (*model.StoreProductReply, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductReply), nil
	}
}

func (s storeProductReplyDo) FirstOrCreate() (*model.StoreProductReply, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductReply), nil
	}
}

func (s storeProductReplyDo) FindByPage(offset int, limit int) (result []*model.StoreProductReply, count int64, err error) {
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

func (s storeProductReplyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeProductReplyDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeProductReplyDo) Delete(models ...*model.StoreProductReply) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeProductReplyDo) withDO(do gen.Dao) *storeProductReplyDo {
	s.DO = *do.(*gen.DO)
	return s
}
