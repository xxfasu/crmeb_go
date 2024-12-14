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

func newSystemAdmin(db *gorm.DB, opts ...gen.DOOption) systemAdmin {
	_systemAdmin := systemAdmin{}

	_systemAdmin.systemAdminDo.UseDB(db, opts...)
	_systemAdmin.systemAdminDo.UseModel(&model.SystemAdmin{})

	tableName := _systemAdmin.systemAdminDo.TableName()
	_systemAdmin.ALL = field.NewAsterisk(tableName)
	_systemAdmin.ID = field.NewInt64(tableName, "id")
	_systemAdmin.Account = field.NewString(tableName, "account")
	_systemAdmin.Pwd = field.NewString(tableName, "pwd")
	_systemAdmin.RealName = field.NewString(tableName, "real_name")
	_systemAdmin.Roles = field.NewString(tableName, "roles")
	_systemAdmin.LastIP = field.NewString(tableName, "last_ip")
	_systemAdmin.LoginCount = field.NewInt64(tableName, "login_count")
	_systemAdmin.Level = field.NewInt32(tableName, "level")
	_systemAdmin.Status = field.NewInt32(tableName, "status")
	_systemAdmin.Phone = field.NewString(tableName, "phone")
	_systemAdmin.IsSms = field.NewInt32(tableName, "is_sms")
	_systemAdmin.CreatedAt = field.NewInt64(tableName, "created_at")
	_systemAdmin.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_systemAdmin.DeletedAt = field.NewField(tableName, "deleted_at")

	_systemAdmin.fillFieldMap()

	return _systemAdmin
}

// systemAdmin 后台管理员表
type systemAdmin struct {
	systemAdminDo systemAdminDo

	ALL        field.Asterisk
	ID         field.Int64  // 后台管理员表ID
	Account    field.String // 后台管理员账号
	Pwd        field.String // 后台管理员密码
	RealName   field.String // 后台管理员姓名
	Roles      field.String // 后台管理员权限(menus_id)
	LastIP     field.String // 后台管理员最后一次登录ip
	LoginCount field.Int64  // 登录次数
	Level      field.Int32  // 后台管理员级别
	Status     field.Int32  // 后台管理员状态 1有效0无效
	Phone      field.String // 手机号码
	IsSms      field.Int32  // 是否接收短信
	CreatedAt  field.Int64
	UpdatedAt  field.Int64
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (s systemAdmin) Table(newTableName string) *systemAdmin {
	s.systemAdminDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s systemAdmin) As(alias string) *systemAdmin {
	s.systemAdminDo.DO = *(s.systemAdminDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *systemAdmin) updateTableName(table string) *systemAdmin {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Account = field.NewString(table, "account")
	s.Pwd = field.NewString(table, "pwd")
	s.RealName = field.NewString(table, "real_name")
	s.Roles = field.NewString(table, "roles")
	s.LastIP = field.NewString(table, "last_ip")
	s.LoginCount = field.NewInt64(table, "login_count")
	s.Level = field.NewInt32(table, "level")
	s.Status = field.NewInt32(table, "status")
	s.Phone = field.NewString(table, "phone")
	s.IsSms = field.NewInt32(table, "is_sms")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *systemAdmin) WithContext(ctx context.Context) ISystemAdminDo {
	return s.systemAdminDo.WithContext(ctx)
}

func (s systemAdmin) TableName() string { return s.systemAdminDo.TableName() }

func (s systemAdmin) Alias() string { return s.systemAdminDo.Alias() }

func (s systemAdmin) Columns(cols ...field.Expr) gen.Columns { return s.systemAdminDo.Columns(cols...) }

func (s *systemAdmin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *systemAdmin) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["account"] = s.Account
	s.fieldMap["pwd"] = s.Pwd
	s.fieldMap["real_name"] = s.RealName
	s.fieldMap["roles"] = s.Roles
	s.fieldMap["last_ip"] = s.LastIP
	s.fieldMap["login_count"] = s.LoginCount
	s.fieldMap["level"] = s.Level
	s.fieldMap["status"] = s.Status
	s.fieldMap["phone"] = s.Phone
	s.fieldMap["is_sms"] = s.IsSms
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s systemAdmin) clone(db *gorm.DB) systemAdmin {
	s.systemAdminDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s systemAdmin) replaceDB(db *gorm.DB) systemAdmin {
	s.systemAdminDo.ReplaceDB(db)
	return s
}

type systemAdminDo struct{ gen.DO }

type ISystemAdminDo interface {
	gen.SubQuery
	Debug() ISystemAdminDo
	WithContext(ctx context.Context) ISystemAdminDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISystemAdminDo
	WriteDB() ISystemAdminDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISystemAdminDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISystemAdminDo
	Not(conds ...gen.Condition) ISystemAdminDo
	Or(conds ...gen.Condition) ISystemAdminDo
	Select(conds ...field.Expr) ISystemAdminDo
	Where(conds ...gen.Condition) ISystemAdminDo
	Order(conds ...field.Expr) ISystemAdminDo
	Distinct(cols ...field.Expr) ISystemAdminDo
	Omit(cols ...field.Expr) ISystemAdminDo
	Join(table schema.Tabler, on ...field.Expr) ISystemAdminDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISystemAdminDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISystemAdminDo
	Group(cols ...field.Expr) ISystemAdminDo
	Having(conds ...gen.Condition) ISystemAdminDo
	Limit(limit int) ISystemAdminDo
	Offset(offset int) ISystemAdminDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemAdminDo
	Unscoped() ISystemAdminDo
	Create(values ...*model.SystemAdmin) error
	CreateInBatches(values []*model.SystemAdmin, batchSize int) error
	Save(values ...*model.SystemAdmin) error
	First() (*model.SystemAdmin, error)
	Take() (*model.SystemAdmin, error)
	Last() (*model.SystemAdmin, error)
	Find() ([]*model.SystemAdmin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemAdmin, err error)
	FindInBatches(result *[]*model.SystemAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SystemAdmin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISystemAdminDo
	Assign(attrs ...field.AssignExpr) ISystemAdminDo
	Joins(fields ...field.RelationField) ISystemAdminDo
	Preload(fields ...field.RelationField) ISystemAdminDo
	FirstOrInit() (*model.SystemAdmin, error)
	FirstOrCreate() (*model.SystemAdmin, error)
	FindByPage(offset int, limit int) (result []*model.SystemAdmin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISystemAdminDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s systemAdminDo) Debug() ISystemAdminDo {
	return s.withDO(s.DO.Debug())
}

func (s systemAdminDo) WithContext(ctx context.Context) ISystemAdminDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemAdminDo) ReadDB() ISystemAdminDo {
	return s.Clauses(dbresolver.Read)
}

func (s systemAdminDo) WriteDB() ISystemAdminDo {
	return s.Clauses(dbresolver.Write)
}

func (s systemAdminDo) Session(config *gorm.Session) ISystemAdminDo {
	return s.withDO(s.DO.Session(config))
}

func (s systemAdminDo) Clauses(conds ...clause.Expression) ISystemAdminDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemAdminDo) Returning(value interface{}, columns ...string) ISystemAdminDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemAdminDo) Not(conds ...gen.Condition) ISystemAdminDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemAdminDo) Or(conds ...gen.Condition) ISystemAdminDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemAdminDo) Select(conds ...field.Expr) ISystemAdminDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemAdminDo) Where(conds ...gen.Condition) ISystemAdminDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemAdminDo) Order(conds ...field.Expr) ISystemAdminDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemAdminDo) Distinct(cols ...field.Expr) ISystemAdminDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemAdminDo) Omit(cols ...field.Expr) ISystemAdminDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemAdminDo) Join(table schema.Tabler, on ...field.Expr) ISystemAdminDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemAdminDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISystemAdminDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemAdminDo) RightJoin(table schema.Tabler, on ...field.Expr) ISystemAdminDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemAdminDo) Group(cols ...field.Expr) ISystemAdminDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemAdminDo) Having(conds ...gen.Condition) ISystemAdminDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemAdminDo) Limit(limit int) ISystemAdminDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemAdminDo) Offset(offset int) ISystemAdminDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemAdminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemAdminDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemAdminDo) Unscoped() ISystemAdminDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemAdminDo) Create(values ...*model.SystemAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemAdminDo) CreateInBatches(values []*model.SystemAdmin, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemAdminDo) Save(values ...*model.SystemAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemAdminDo) First() (*model.SystemAdmin, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAdmin), nil
	}
}

func (s systemAdminDo) Take() (*model.SystemAdmin, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAdmin), nil
	}
}

func (s systemAdminDo) Last() (*model.SystemAdmin, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAdmin), nil
	}
}

func (s systemAdminDo) Find() ([]*model.SystemAdmin, error) {
	result, err := s.DO.Find()
	return result.([]*model.SystemAdmin), err
}

func (s systemAdminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemAdmin, err error) {
	buf := make([]*model.SystemAdmin, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemAdminDo) FindInBatches(result *[]*model.SystemAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemAdminDo) Attrs(attrs ...field.AssignExpr) ISystemAdminDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemAdminDo) Assign(attrs ...field.AssignExpr) ISystemAdminDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemAdminDo) Joins(fields ...field.RelationField) ISystemAdminDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemAdminDo) Preload(fields ...field.RelationField) ISystemAdminDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemAdminDo) FirstOrInit() (*model.SystemAdmin, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAdmin), nil
	}
}

func (s systemAdminDo) FirstOrCreate() (*model.SystemAdmin, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemAdmin), nil
	}
}

func (s systemAdminDo) FindByPage(offset int, limit int) (result []*model.SystemAdmin, count int64, err error) {
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

func (s systemAdminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s systemAdminDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s systemAdminDo) Delete(models ...*model.SystemAdmin) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *systemAdminDo) withDO(do gen.Dao) *systemAdminDo {
	s.DO = *do.(*gen.DO)
	return s
}
