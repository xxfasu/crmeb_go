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

func newSystemGroupDatum(db *gorm.DB, opts ...gen.DOOption) systemGroupDatum {
	_systemGroupDatum := systemGroupDatum{}

	_systemGroupDatum.systemGroupDatumDo.UseDB(db, opts...)
	_systemGroupDatum.systemGroupDatumDo.UseModel(&model.SystemGroupDatum{})

	tableName := _systemGroupDatum.systemGroupDatumDo.TableName()
	_systemGroupDatum.ALL = field.NewAsterisk(tableName)
	_systemGroupDatum.ID = field.NewInt64(tableName, "id")
	_systemGroupDatum.Gid = field.NewInt64(tableName, "gid")
	_systemGroupDatum.Value = field.NewString(tableName, "value")
	_systemGroupDatum.Sort = field.NewInt64(tableName, "sort")
	_systemGroupDatum.Status = field.NewBool(tableName, "status")
	_systemGroupDatum.CreatedAt = field.NewInt64(tableName, "created_at")
	_systemGroupDatum.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_systemGroupDatum.DeletedAt = field.NewField(tableName, "deleted_at")

	_systemGroupDatum.fillFieldMap()

	return _systemGroupDatum
}

// systemGroupDatum 组合数据详情表
type systemGroupDatum struct {
	systemGroupDatumDo systemGroupDatumDo

	ALL       field.Asterisk
	ID        field.Int64  // 组合数据详情ID
	Gid       field.Int64  // 对应的数据组id
	Value     field.String // 数据组对应的数据值（json数据）
	Sort      field.Int64  // 数据排序
	Status    field.Bool   // 状态（1：开启；2：关闭；）
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (s systemGroupDatum) Table(newTableName string) *systemGroupDatum {
	s.systemGroupDatumDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s systemGroupDatum) As(alias string) *systemGroupDatum {
	s.systemGroupDatumDo.DO = *(s.systemGroupDatumDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *systemGroupDatum) updateTableName(table string) *systemGroupDatum {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Gid = field.NewInt64(table, "gid")
	s.Value = field.NewString(table, "value")
	s.Sort = field.NewInt64(table, "sort")
	s.Status = field.NewBool(table, "status")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *systemGroupDatum) WithContext(ctx context.Context) ISystemGroupDatumDo {
	return s.systemGroupDatumDo.WithContext(ctx)
}

func (s systemGroupDatum) TableName() string { return s.systemGroupDatumDo.TableName() }

func (s systemGroupDatum) Alias() string { return s.systemGroupDatumDo.Alias() }

func (s systemGroupDatum) Columns(cols ...field.Expr) gen.Columns {
	return s.systemGroupDatumDo.Columns(cols...)
}

func (s *systemGroupDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *systemGroupDatum) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["gid"] = s.Gid
	s.fieldMap["value"] = s.Value
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["status"] = s.Status
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s systemGroupDatum) clone(db *gorm.DB) systemGroupDatum {
	s.systemGroupDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s systemGroupDatum) replaceDB(db *gorm.DB) systemGroupDatum {
	s.systemGroupDatumDo.ReplaceDB(db)
	return s
}

type systemGroupDatumDo struct{ gen.DO }

type ISystemGroupDatumDo interface {
	gen.SubQuery
	Debug() ISystemGroupDatumDo
	WithContext(ctx context.Context) ISystemGroupDatumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISystemGroupDatumDo
	WriteDB() ISystemGroupDatumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISystemGroupDatumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISystemGroupDatumDo
	Not(conds ...gen.Condition) ISystemGroupDatumDo
	Or(conds ...gen.Condition) ISystemGroupDatumDo
	Select(conds ...field.Expr) ISystemGroupDatumDo
	Where(conds ...gen.Condition) ISystemGroupDatumDo
	Order(conds ...field.Expr) ISystemGroupDatumDo
	Distinct(cols ...field.Expr) ISystemGroupDatumDo
	Omit(cols ...field.Expr) ISystemGroupDatumDo
	Join(table schema.Tabler, on ...field.Expr) ISystemGroupDatumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISystemGroupDatumDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISystemGroupDatumDo
	Group(cols ...field.Expr) ISystemGroupDatumDo
	Having(conds ...gen.Condition) ISystemGroupDatumDo
	Limit(limit int) ISystemGroupDatumDo
	Offset(offset int) ISystemGroupDatumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemGroupDatumDo
	Unscoped() ISystemGroupDatumDo
	Create(values ...*model.SystemGroupDatum) error
	CreateInBatches(values []*model.SystemGroupDatum, batchSize int) error
	Save(values ...*model.SystemGroupDatum) error
	First() (*model.SystemGroupDatum, error)
	Take() (*model.SystemGroupDatum, error)
	Last() (*model.SystemGroupDatum, error)
	Find() ([]*model.SystemGroupDatum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemGroupDatum, err error)
	FindInBatches(result *[]*model.SystemGroupDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SystemGroupDatum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISystemGroupDatumDo
	Assign(attrs ...field.AssignExpr) ISystemGroupDatumDo
	Joins(fields ...field.RelationField) ISystemGroupDatumDo
	Preload(fields ...field.RelationField) ISystemGroupDatumDo
	FirstOrInit() (*model.SystemGroupDatum, error)
	FirstOrCreate() (*model.SystemGroupDatum, error)
	FindByPage(offset int, limit int) (result []*model.SystemGroupDatum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISystemGroupDatumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s systemGroupDatumDo) Debug() ISystemGroupDatumDo {
	return s.withDO(s.DO.Debug())
}

func (s systemGroupDatumDo) WithContext(ctx context.Context) ISystemGroupDatumDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemGroupDatumDo) ReadDB() ISystemGroupDatumDo {
	return s.Clauses(dbresolver.Read)
}

func (s systemGroupDatumDo) WriteDB() ISystemGroupDatumDo {
	return s.Clauses(dbresolver.Write)
}

func (s systemGroupDatumDo) Session(config *gorm.Session) ISystemGroupDatumDo {
	return s.withDO(s.DO.Session(config))
}

func (s systemGroupDatumDo) Clauses(conds ...clause.Expression) ISystemGroupDatumDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemGroupDatumDo) Returning(value interface{}, columns ...string) ISystemGroupDatumDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemGroupDatumDo) Not(conds ...gen.Condition) ISystemGroupDatumDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemGroupDatumDo) Or(conds ...gen.Condition) ISystemGroupDatumDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemGroupDatumDo) Select(conds ...field.Expr) ISystemGroupDatumDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemGroupDatumDo) Where(conds ...gen.Condition) ISystemGroupDatumDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemGroupDatumDo) Order(conds ...field.Expr) ISystemGroupDatumDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemGroupDatumDo) Distinct(cols ...field.Expr) ISystemGroupDatumDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemGroupDatumDo) Omit(cols ...field.Expr) ISystemGroupDatumDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemGroupDatumDo) Join(table schema.Tabler, on ...field.Expr) ISystemGroupDatumDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemGroupDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISystemGroupDatumDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemGroupDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) ISystemGroupDatumDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemGroupDatumDo) Group(cols ...field.Expr) ISystemGroupDatumDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemGroupDatumDo) Having(conds ...gen.Condition) ISystemGroupDatumDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemGroupDatumDo) Limit(limit int) ISystemGroupDatumDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemGroupDatumDo) Offset(offset int) ISystemGroupDatumDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemGroupDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemGroupDatumDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemGroupDatumDo) Unscoped() ISystemGroupDatumDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemGroupDatumDo) Create(values ...*model.SystemGroupDatum) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemGroupDatumDo) CreateInBatches(values []*model.SystemGroupDatum, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemGroupDatumDo) Save(values ...*model.SystemGroupDatum) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemGroupDatumDo) First() (*model.SystemGroupDatum, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroupDatum), nil
	}
}

func (s systemGroupDatumDo) Take() (*model.SystemGroupDatum, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroupDatum), nil
	}
}

func (s systemGroupDatumDo) Last() (*model.SystemGroupDatum, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroupDatum), nil
	}
}

func (s systemGroupDatumDo) Find() ([]*model.SystemGroupDatum, error) {
	result, err := s.DO.Find()
	return result.([]*model.SystemGroupDatum), err
}

func (s systemGroupDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemGroupDatum, err error) {
	buf := make([]*model.SystemGroupDatum, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemGroupDatumDo) FindInBatches(result *[]*model.SystemGroupDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemGroupDatumDo) Attrs(attrs ...field.AssignExpr) ISystemGroupDatumDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemGroupDatumDo) Assign(attrs ...field.AssignExpr) ISystemGroupDatumDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemGroupDatumDo) Joins(fields ...field.RelationField) ISystemGroupDatumDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemGroupDatumDo) Preload(fields ...field.RelationField) ISystemGroupDatumDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemGroupDatumDo) FirstOrInit() (*model.SystemGroupDatum, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroupDatum), nil
	}
}

func (s systemGroupDatumDo) FirstOrCreate() (*model.SystemGroupDatum, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroupDatum), nil
	}
}

func (s systemGroupDatumDo) FindByPage(offset int, limit int) (result []*model.SystemGroupDatum, count int64, err error) {
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

func (s systemGroupDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s systemGroupDatumDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s systemGroupDatumDo) Delete(models ...*model.SystemGroupDatum) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *systemGroupDatumDo) withDO(do gen.Dao) *systemGroupDatumDo {
	s.DO = *do.(*gen.DO)
	return s
}
