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

func newSystemAttachment(db *gorm.DB, opts ...gen.DOOption) systemAttachment {
	_systemAttachment := systemAttachment{}

	_systemAttachment.systemAttachmentDo.UseDB(db, opts...)
	_systemAttachment.systemAttachmentDo.UseModel(&model.SystemAttachment{})

	tableName := _systemAttachment.systemAttachmentDo.TableName()
	_systemAttachment.ALL = field.NewAsterisk(tableName)
	_systemAttachment.AttID = field.NewInt64(tableName, "att_id")
	_systemAttachment.Name = field.NewString(tableName, "name")
	_systemAttachment.AttDir = field.NewString(tableName, "att_dir")
	_systemAttachment.SattDir = field.NewString(tableName, "satt_dir")
	_systemAttachment.AttSize = field.NewString(tableName, "att_size")
	_systemAttachment.AttType = field.NewString(tableName, "att_type")
	_systemAttachment.Pid = field.NewInt64(tableName, "pid")
	_systemAttachment.ImageType = field.NewInt32(tableName, "image_type")
	_systemAttachment.CreatedAt = field.NewInt64(tableName, "created_at")
	_systemAttachment.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_systemAttachment.DeletedAt = field.NewField(tableName, "deleted_at")

	_systemAttachment.fillFieldMap()

	return _systemAttachment
}

// systemAttachment 附件管理表
type systemAttachment struct {
	systemAttachmentDo systemAttachmentDo

	ALL       field.Asterisk
	AttID     field.Int64
	Name      field.String // 附件名称
	AttDir    field.String // 附件路径
	SattDir   field.String // 压缩图片路径
	AttSize   field.String // 附件大小
	AttType   field.String // 附件类型
	Pid       field.Int64  // 分类ID0编辑器,1商品图片,2拼团图片,3砍价图片,4秒杀图片,5文章图片,6组合数据图， 7前台用户
	ImageType field.Int32  // 图片上传类型 1本地 2七牛云 3OSS 4COS
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (s systemAttachment) Table(newTableName string) *systemAttachment {
	s.systemAttachmentDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s systemAttachment) As(alias string) *systemAttachment {
	s.systemAttachmentDo.DO = *(s.systemAttachmentDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *systemAttachment) updateTableName(table string) *systemAttachment {
	s.ALL = field.NewAsterisk(table)
	s.AttID = field.NewInt64(table, "att_id")
	s.Name = field.NewString(table, "name")
	s.AttDir = field.NewString(table, "att_dir")
	s.SattDir = field.NewString(table, "satt_dir")
	s.AttSize = field.NewString(table, "att_size")
	s.AttType = field.NewString(table, "att_type")
	s.Pid = field.NewInt64(table, "pid")
	s.ImageType = field.NewInt32(table, "image_type")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *systemAttachment) WithContext(ctx context.Context) ISystemAttachmentDo {
	return s.systemAttachmentDo.WithContext(ctx)
}

func (s systemAttachment) TableName() string { return s.systemAttachmentDo.TableName() }

func (s systemAttachment) Alias() string { return s.systemAttachmentDo.Alias() }

func (s systemAttachment) Columns(cols ...field.Expr) gen.Columns {
	return s.systemAttachmentDo.Columns(cols...)
}

func (s *systemAttachment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *systemAttachment) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["att_id"] = s.AttID
	s.fieldMap["name"] = s.Name
	s.fieldMap["att_dir"] = s.AttDir
	s.fieldMap["satt_dir"] = s.SattDir
	s.fieldMap["att_size"] = s.AttSize
	s.fieldMap["att_type"] = s.AttType
	s.fieldMap["pid"] = s.Pid
	s.fieldMap["image_type"] = s.ImageType
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s systemAttachment) clone(db *gorm.DB) systemAttachment {
	s.systemAttachmentDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s systemAttachment) replaceDB(db *gorm.DB) systemAttachment {
	s.systemAttachmentDo.ReplaceDB(db)
	return s
}

type systemAttachmentDo struct{ gen.DO }

type ISystemAttachmentDo interface {
	gen.SubQuery
	Debug() ISystemAttachmentDo
	WithContext(ctx context.Context) ISystemAttachmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISystemAttachmentDo
	WriteDB() ISystemAttachmentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISystemAttachmentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISystemAttachmentDo
	Not(conds ...gen.Condition) ISystemAttachmentDo
	Or(conds ...gen.Condition) ISystemAttachmentDo
	Select(conds ...field.Expr) ISystemAttachmentDo
	Where(conds ...gen.Condition) ISystemAttachmentDo
	Order(conds ...field.Expr) ISystemAttachmentDo
	Distinct(cols ...field.Expr) ISystemAttachmentDo
	Omit(cols ...field.Expr) ISystemAttachmentDo
	Join(table schema.Tabler, on ...field.Expr) ISystemAttachmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISystemAttachmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISystemAttachmentDo
	Group(cols ...field.Expr) ISystemAttachmentDo
	Having(conds ...gen.Condition) ISystemAttachmentDo
	Limit(limit int) ISystemAttachmentDo
	Offset(offset int) ISystemAttachmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemAttachmentDo
	Unscoped() ISystemAttachmentDo
	Create(values ...*model.SystemAttachment) error
	CreateInBatches(values []*model.SystemAttachment, batchSize int) error
	Save(values ...*model.SystemAttachment) error
	First() (*model.SystemAttachment, error)
	Take() (*model.SystemAttachment, error)
	Last() (*model.SystemAttachment, error)
	Find() ([]*model.SystemAttachment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemAttachment, err error)
	FindInBatches(result *[]*model.SystemAttachment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SystemAttachment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISystemAttachmentDo
	Assign(attrs ...field.AssignExpr) ISystemAttachmentDo
	Joins(fields ...field.RelationField) ISystemAttachmentDo
	Preload(fields ...field.RelationField) ISystemAttachmentDo
	FirstOrInit() (*model.SystemAttachment, error)
	FirstOrCreate() (*model.SystemAttachment, error)
	FindByPage(offset int, limit int) (result []*model.SystemAttachment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISystemAttachmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s systemAttachmentDo) Debug() ISystemAttachmentDo {
	return s.withDO(s.DO.Debug())
}

func (s systemAttachmentDo) WithContext(ctx context.Context) ISystemAttachmentDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemAttachmentDo) ReadDB() ISystemAttachmentDo {
	return s.Clauses(dbresolver.Read)
}

func (s systemAttachmentDo) WriteDB() ISystemAttachmentDo {
	return s.Clauses(dbresolver.Write)
}

func (s systemAttachmentDo) Session(config *gorm.Session) ISystemAttachmentDo {
	return s.withDO(s.DO.Session(config))
}

func (s systemAttachmentDo) Clauses(conds ...clause.Expression) ISystemAttachmentDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemAttachmentDo) Returning(value interface{}, columns ...string) ISystemAttachmentDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemAttachmentDo) Not(conds ...gen.Condition) ISystemAttachmentDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemAttachmentDo) Or(conds ...gen.Condition) ISystemAttachmentDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemAttachmentDo) Select(conds ...field.Expr) ISystemAttachmentDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemAttachmentDo) Where(conds ...gen.Condition) ISystemAttachmentDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemAttachmentDo) Order(conds ...field.Expr) ISystemAttachmentDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemAttachmentDo) Distinct(cols ...field.Expr) ISystemAttachmentDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemAttachmentDo) Omit(cols ...field.Expr) ISystemAttachmentDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemAttachmentDo) Join(table schema.Tabler, on ...field.Expr) ISystemAttachmentDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemAttachmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISystemAttachmentDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemAttachmentDo) RightJoin(table schema.Tabler, on ...field.Expr) ISystemAttachmentDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemAttachmentDo) Group(cols ...field.Expr) ISystemAttachmentDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemAttachmentDo) Having(conds ...gen.Condition) ISystemAttachmentDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemAttachmentDo) Limit(limit int) ISystemAttachmentDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemAttachmentDo) Offset(offset int) ISystemAttachmentDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemAttachmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemAttachmentDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemAttachmentDo) Unscoped() ISystemAttachmentDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemAttachmentDo) Create(values ...*model.SystemAttachment) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemAttachmentDo) CreateInBatches(values []*model.SystemAttachment, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemAttachmentDo) Save(values ...*model.SystemAttachment) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemAttachmentDo) First() (*model.SystemAttachment, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAttachment), nil
	}
}

func (s systemAttachmentDo) Take() (*model.SystemAttachment, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAttachment), nil
	}
}

func (s systemAttachmentDo) Last() (*model.SystemAttachment, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAttachment), nil
	}
}

func (s systemAttachmentDo) Find() ([]*model.SystemAttachment, error) {
	result, err := s.DO.Find()
	return result.([]*model.SystemAttachment), err
}

func (s systemAttachmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemAttachment, err error) {
	buf := make([]*model.SystemAttachment, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemAttachmentDo) FindInBatches(result *[]*model.SystemAttachment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemAttachmentDo) Attrs(attrs ...field.AssignExpr) ISystemAttachmentDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemAttachmentDo) Assign(attrs ...field.AssignExpr) ISystemAttachmentDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemAttachmentDo) Joins(fields ...field.RelationField) ISystemAttachmentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemAttachmentDo) Preload(fields ...field.RelationField) ISystemAttachmentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemAttachmentDo) FirstOrInit() (*model.SystemAttachment, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAttachment), nil
	}
}

func (s systemAttachmentDo) FirstOrCreate() (*model.SystemAttachment, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAttachment), nil
	}
}

func (s systemAttachmentDo) FindByPage(offset int, limit int) (result []*model.SystemAttachment, count int64, err error) {
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

func (s systemAttachmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s systemAttachmentDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s systemAttachmentDo) Delete(models ...*model.SystemAttachment) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *systemAttachmentDo) withDO(do gen.Dao) *systemAttachmentDo {
	s.DO = *do.(*gen.DO)
	return s
}
