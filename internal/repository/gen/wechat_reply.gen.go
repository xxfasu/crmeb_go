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

func newWechatReply(db *gorm.DB, opts ...gen.DOOption) wechatReply {
	_wechatReply := wechatReply{}

	_wechatReply.wechatReplyDo.UseDB(db, opts...)
	_wechatReply.wechatReplyDo.UseModel(&model.WechatReply{})

	tableName := _wechatReply.wechatReplyDo.TableName()
	_wechatReply.ALL = field.NewAsterisk(tableName)
	_wechatReply.ID = field.NewInt64(tableName, "id")
	_wechatReply.Keywords = field.NewString(tableName, "keywords")
	_wechatReply.Type = field.NewString(tableName, "type")
	_wechatReply.Data = field.NewString(tableName, "data")
	_wechatReply.Status = field.NewInt32(tableName, "status")
	_wechatReply.CreatedAt = field.NewInt64(tableName, "created_at")
	_wechatReply.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_wechatReply.DeletedAt = field.NewField(tableName, "deleted_at")

	_wechatReply.fillFieldMap()

	return _wechatReply
}

// wechatReply 微信关键字回复表
type wechatReply struct {
	wechatReplyDo wechatReplyDo

	ALL       field.Asterisk
	ID        field.Int64  // 微信关键字回复id
	Keywords  field.String // 关键字
	Type      field.String // 回复类型
	Data      field.String // 回复数据
	Status    field.Int32  // 回复状态 0=不可用  1 =可用
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (w wechatReply) Table(newTableName string) *wechatReply {
	w.wechatReplyDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wechatReply) As(alias string) *wechatReply {
	w.wechatReplyDo.DO = *(w.wechatReplyDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wechatReply) updateTableName(table string) *wechatReply {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Keywords = field.NewString(table, "keywords")
	w.Type = field.NewString(table, "type")
	w.Data = field.NewString(table, "data")
	w.Status = field.NewInt32(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")

	w.fillFieldMap()

	return w
}

func (w *wechatReply) WithContext(ctx context.Context) IWechatReplyDo {
	return w.wechatReplyDo.WithContext(ctx)
}

func (w wechatReply) TableName() string { return w.wechatReplyDo.TableName() }

func (w wechatReply) Alias() string { return w.wechatReplyDo.Alias() }

func (w wechatReply) Columns(cols ...field.Expr) gen.Columns { return w.wechatReplyDo.Columns(cols...) }

func (w *wechatReply) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wechatReply) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["keywords"] = w.Keywords
	w.fieldMap["type"] = w.Type
	w.fieldMap["data"] = w.Data
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
}

func (w wechatReply) clone(db *gorm.DB) wechatReply {
	w.wechatReplyDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wechatReply) replaceDB(db *gorm.DB) wechatReply {
	w.wechatReplyDo.ReplaceDB(db)
	return w
}

type wechatReplyDo struct{ gen.DO }

type IWechatReplyDo interface {
	gen.SubQuery
	Debug() IWechatReplyDo
	WithContext(ctx context.Context) IWechatReplyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWechatReplyDo
	WriteDB() IWechatReplyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWechatReplyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWechatReplyDo
	Not(conds ...gen.Condition) IWechatReplyDo
	Or(conds ...gen.Condition) IWechatReplyDo
	Select(conds ...field.Expr) IWechatReplyDo
	Where(conds ...gen.Condition) IWechatReplyDo
	Order(conds ...field.Expr) IWechatReplyDo
	Distinct(cols ...field.Expr) IWechatReplyDo
	Omit(cols ...field.Expr) IWechatReplyDo
	Join(table schema.Tabler, on ...field.Expr) IWechatReplyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWechatReplyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWechatReplyDo
	Group(cols ...field.Expr) IWechatReplyDo
	Having(conds ...gen.Condition) IWechatReplyDo
	Limit(limit int) IWechatReplyDo
	Offset(offset int) IWechatReplyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWechatReplyDo
	Unscoped() IWechatReplyDo
	Create(values ...*model.WechatReply) error
	CreateInBatches(values []*model.WechatReply, batchSize int) error
	Save(values ...*model.WechatReply) error
	First() (*model.WechatReply, error)
	Take() (*model.WechatReply, error)
	Last() (*model.WechatReply, error)
	Find() ([]*model.WechatReply, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WechatReply, err error)
	FindInBatches(result *[]*model.WechatReply, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WechatReply) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWechatReplyDo
	Assign(attrs ...field.AssignExpr) IWechatReplyDo
	Joins(fields ...field.RelationField) IWechatReplyDo
	Preload(fields ...field.RelationField) IWechatReplyDo
	FirstOrInit() (*model.WechatReply, error)
	FirstOrCreate() (*model.WechatReply, error)
	FindByPage(offset int, limit int) (result []*model.WechatReply, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWechatReplyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wechatReplyDo) Debug() IWechatReplyDo {
	return w.withDO(w.DO.Debug())
}

func (w wechatReplyDo) WithContext(ctx context.Context) IWechatReplyDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wechatReplyDo) ReadDB() IWechatReplyDo {
	return w.Clauses(dbresolver.Read)
}

func (w wechatReplyDo) WriteDB() IWechatReplyDo {
	return w.Clauses(dbresolver.Write)
}

func (w wechatReplyDo) Session(config *gorm.Session) IWechatReplyDo {
	return w.withDO(w.DO.Session(config))
}

func (w wechatReplyDo) Clauses(conds ...clause.Expression) IWechatReplyDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wechatReplyDo) Returning(value interface{}, columns ...string) IWechatReplyDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wechatReplyDo) Not(conds ...gen.Condition) IWechatReplyDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wechatReplyDo) Or(conds ...gen.Condition) IWechatReplyDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wechatReplyDo) Select(conds ...field.Expr) IWechatReplyDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wechatReplyDo) Where(conds ...gen.Condition) IWechatReplyDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wechatReplyDo) Order(conds ...field.Expr) IWechatReplyDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wechatReplyDo) Distinct(cols ...field.Expr) IWechatReplyDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wechatReplyDo) Omit(cols ...field.Expr) IWechatReplyDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wechatReplyDo) Join(table schema.Tabler, on ...field.Expr) IWechatReplyDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wechatReplyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWechatReplyDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wechatReplyDo) RightJoin(table schema.Tabler, on ...field.Expr) IWechatReplyDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wechatReplyDo) Group(cols ...field.Expr) IWechatReplyDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wechatReplyDo) Having(conds ...gen.Condition) IWechatReplyDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wechatReplyDo) Limit(limit int) IWechatReplyDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wechatReplyDo) Offset(offset int) IWechatReplyDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wechatReplyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWechatReplyDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wechatReplyDo) Unscoped() IWechatReplyDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wechatReplyDo) Create(values ...*model.WechatReply) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wechatReplyDo) CreateInBatches(values []*model.WechatReply, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wechatReplyDo) Save(values ...*model.WechatReply) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wechatReplyDo) First() (*model.WechatReply, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatReply), nil
	}
}

func (w wechatReplyDo) Take() (*model.WechatReply, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatReply), nil
	}
}

func (w wechatReplyDo) Last() (*model.WechatReply, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatReply), nil
	}
}

func (w wechatReplyDo) Find() ([]*model.WechatReply, error) {
	result, err := w.DO.Find()
	return result.([]*model.WechatReply), err
}

func (w wechatReplyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WechatReply, err error) {
	buf := make([]*model.WechatReply, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wechatReplyDo) FindInBatches(result *[]*model.WechatReply, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wechatReplyDo) Attrs(attrs ...field.AssignExpr) IWechatReplyDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wechatReplyDo) Assign(attrs ...field.AssignExpr) IWechatReplyDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wechatReplyDo) Joins(fields ...field.RelationField) IWechatReplyDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wechatReplyDo) Preload(fields ...field.RelationField) IWechatReplyDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wechatReplyDo) FirstOrInit() (*model.WechatReply, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatReply), nil
	}
}

func (w wechatReplyDo) FirstOrCreate() (*model.WechatReply, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatReply), nil
	}
}

func (w wechatReplyDo) FindByPage(offset int, limit int) (result []*model.WechatReply, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wechatReplyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wechatReplyDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wechatReplyDo) Delete(models ...*model.WechatReply) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wechatReplyDo) withDO(do gen.Dao) *wechatReplyDo {
	w.DO = *do.(*gen.DO)
	return w
}
