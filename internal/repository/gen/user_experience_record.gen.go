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

func newUserExperienceRecord(db *gorm.DB, opts ...gen.DOOption) userExperienceRecord {
	_userExperienceRecord := userExperienceRecord{}

	_userExperienceRecord.userExperienceRecordDo.UseDB(db, opts...)
	_userExperienceRecord.userExperienceRecordDo.UseModel(&model.UserExperienceRecord{})

	tableName := _userExperienceRecord.userExperienceRecordDo.TableName()
	_userExperienceRecord.ALL = field.NewAsterisk(tableName)
	_userExperienceRecord.ID = field.NewInt64(tableName, "id")
	_userExperienceRecord.UID = field.NewInt64(tableName, "uid")
	_userExperienceRecord.LinkID = field.NewString(tableName, "link_id")
	_userExperienceRecord.LinkType = field.NewString(tableName, "link_type")
	_userExperienceRecord.Type = field.NewInt64(tableName, "type")
	_userExperienceRecord.Title = field.NewString(tableName, "title")
	_userExperienceRecord.Experience = field.NewInt64(tableName, "experience")
	_userExperienceRecord.Balance = field.NewInt64(tableName, "balance")
	_userExperienceRecord.Mark = field.NewString(tableName, "mark")
	_userExperienceRecord.Status = field.NewInt64(tableName, "status")
	_userExperienceRecord.CreatedAt = field.NewInt64(tableName, "created_at")
	_userExperienceRecord.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_userExperienceRecord.DeletedAt = field.NewField(tableName, "deleted_at")

	_userExperienceRecord.fillFieldMap()

	return _userExperienceRecord
}

// userExperienceRecord 用户经验记录表
type userExperienceRecord struct {
	userExperienceRecordDo userExperienceRecordDo

	ALL        field.Asterisk
	ID         field.Int64  // 记录id
	UID        field.Int64  // 用户uid
	LinkID     field.String // 关联id-orderNo,(sign,system默认为0）
	LinkType   field.String // 关联类型（order,sign,system）
	Type       field.Int64  // 类型：1-增加，2-扣减
	Title      field.String // 标题
	Experience field.Int64  // 经验
	Balance    field.Int64  // 剩余
	Mark       field.String // 备注
	Status     field.Int64  // 状态：1-成功（保留字段）
	CreatedAt  field.Int64
	UpdatedAt  field.Int64
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (u userExperienceRecord) Table(newTableName string) *userExperienceRecord {
	u.userExperienceRecordDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userExperienceRecord) As(alias string) *userExperienceRecord {
	u.userExperienceRecordDo.DO = *(u.userExperienceRecordDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userExperienceRecord) updateTableName(table string) *userExperienceRecord {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UID = field.NewInt64(table, "uid")
	u.LinkID = field.NewString(table, "link_id")
	u.LinkType = field.NewString(table, "link_type")
	u.Type = field.NewInt64(table, "type")
	u.Title = field.NewString(table, "title")
	u.Experience = field.NewInt64(table, "experience")
	u.Balance = field.NewInt64(table, "balance")
	u.Mark = field.NewString(table, "mark")
	u.Status = field.NewInt64(table, "status")
	u.CreatedAt = field.NewInt64(table, "created_at")
	u.UpdatedAt = field.NewInt64(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userExperienceRecord) WithContext(ctx context.Context) IUserExperienceRecordDo {
	return u.userExperienceRecordDo.WithContext(ctx)
}

func (u userExperienceRecord) TableName() string { return u.userExperienceRecordDo.TableName() }

func (u userExperienceRecord) Alias() string { return u.userExperienceRecordDo.Alias() }

func (u userExperienceRecord) Columns(cols ...field.Expr) gen.Columns {
	return u.userExperienceRecordDo.Columns(cols...)
}

func (u *userExperienceRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userExperienceRecord) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 13)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["link_id"] = u.LinkID
	u.fieldMap["link_type"] = u.LinkType
	u.fieldMap["type"] = u.Type
	u.fieldMap["title"] = u.Title
	u.fieldMap["experience"] = u.Experience
	u.fieldMap["balance"] = u.Balance
	u.fieldMap["mark"] = u.Mark
	u.fieldMap["status"] = u.Status
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userExperienceRecord) clone(db *gorm.DB) userExperienceRecord {
	u.userExperienceRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userExperienceRecord) replaceDB(db *gorm.DB) userExperienceRecord {
	u.userExperienceRecordDo.ReplaceDB(db)
	return u
}

type userExperienceRecordDo struct{ gen.DO }

type IUserExperienceRecordDo interface {
	gen.SubQuery
	Debug() IUserExperienceRecordDo
	WithContext(ctx context.Context) IUserExperienceRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserExperienceRecordDo
	WriteDB() IUserExperienceRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserExperienceRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserExperienceRecordDo
	Not(conds ...gen.Condition) IUserExperienceRecordDo
	Or(conds ...gen.Condition) IUserExperienceRecordDo
	Select(conds ...field.Expr) IUserExperienceRecordDo
	Where(conds ...gen.Condition) IUserExperienceRecordDo
	Order(conds ...field.Expr) IUserExperienceRecordDo
	Distinct(cols ...field.Expr) IUserExperienceRecordDo
	Omit(cols ...field.Expr) IUserExperienceRecordDo
	Join(table schema.Tabler, on ...field.Expr) IUserExperienceRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserExperienceRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserExperienceRecordDo
	Group(cols ...field.Expr) IUserExperienceRecordDo
	Having(conds ...gen.Condition) IUserExperienceRecordDo
	Limit(limit int) IUserExperienceRecordDo
	Offset(offset int) IUserExperienceRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserExperienceRecordDo
	Unscoped() IUserExperienceRecordDo
	Create(values ...*model.UserExperienceRecord) error
	CreateInBatches(values []*model.UserExperienceRecord, batchSize int) error
	Save(values ...*model.UserExperienceRecord) error
	First() (*model.UserExperienceRecord, error)
	Take() (*model.UserExperienceRecord, error)
	Last() (*model.UserExperienceRecord, error)
	Find() ([]*model.UserExperienceRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserExperienceRecord, err error)
	FindInBatches(result *[]*model.UserExperienceRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserExperienceRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserExperienceRecordDo
	Assign(attrs ...field.AssignExpr) IUserExperienceRecordDo
	Joins(fields ...field.RelationField) IUserExperienceRecordDo
	Preload(fields ...field.RelationField) IUserExperienceRecordDo
	FirstOrInit() (*model.UserExperienceRecord, error)
	FirstOrCreate() (*model.UserExperienceRecord, error)
	FindByPage(offset int, limit int) (result []*model.UserExperienceRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserExperienceRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userExperienceRecordDo) Debug() IUserExperienceRecordDo {
	return u.withDO(u.DO.Debug())
}

func (u userExperienceRecordDo) WithContext(ctx context.Context) IUserExperienceRecordDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userExperienceRecordDo) ReadDB() IUserExperienceRecordDo {
	return u.Clauses(dbresolver.Read)
}

func (u userExperienceRecordDo) WriteDB() IUserExperienceRecordDo {
	return u.Clauses(dbresolver.Write)
}

func (u userExperienceRecordDo) Session(config *gorm.Session) IUserExperienceRecordDo {
	return u.withDO(u.DO.Session(config))
}

func (u userExperienceRecordDo) Clauses(conds ...clause.Expression) IUserExperienceRecordDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userExperienceRecordDo) Returning(value interface{}, columns ...string) IUserExperienceRecordDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userExperienceRecordDo) Not(conds ...gen.Condition) IUserExperienceRecordDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userExperienceRecordDo) Or(conds ...gen.Condition) IUserExperienceRecordDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userExperienceRecordDo) Select(conds ...field.Expr) IUserExperienceRecordDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userExperienceRecordDo) Where(conds ...gen.Condition) IUserExperienceRecordDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userExperienceRecordDo) Order(conds ...field.Expr) IUserExperienceRecordDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userExperienceRecordDo) Distinct(cols ...field.Expr) IUserExperienceRecordDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userExperienceRecordDo) Omit(cols ...field.Expr) IUserExperienceRecordDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userExperienceRecordDo) Join(table schema.Tabler, on ...field.Expr) IUserExperienceRecordDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userExperienceRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserExperienceRecordDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userExperienceRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserExperienceRecordDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userExperienceRecordDo) Group(cols ...field.Expr) IUserExperienceRecordDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userExperienceRecordDo) Having(conds ...gen.Condition) IUserExperienceRecordDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userExperienceRecordDo) Limit(limit int) IUserExperienceRecordDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userExperienceRecordDo) Offset(offset int) IUserExperienceRecordDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userExperienceRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserExperienceRecordDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userExperienceRecordDo) Unscoped() IUserExperienceRecordDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userExperienceRecordDo) Create(values ...*model.UserExperienceRecord) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userExperienceRecordDo) CreateInBatches(values []*model.UserExperienceRecord, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userExperienceRecordDo) Save(values ...*model.UserExperienceRecord) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userExperienceRecordDo) First() (*model.UserExperienceRecord, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExperienceRecord), nil
	}
}

func (u userExperienceRecordDo) Take() (*model.UserExperienceRecord, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExperienceRecord), nil
	}
}

func (u userExperienceRecordDo) Last() (*model.UserExperienceRecord, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExperienceRecord), nil
	}
}

func (u userExperienceRecordDo) Find() ([]*model.UserExperienceRecord, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserExperienceRecord), err
}

func (u userExperienceRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserExperienceRecord, err error) {
	buf := make([]*model.UserExperienceRecord, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userExperienceRecordDo) FindInBatches(result *[]*model.UserExperienceRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userExperienceRecordDo) Attrs(attrs ...field.AssignExpr) IUserExperienceRecordDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userExperienceRecordDo) Assign(attrs ...field.AssignExpr) IUserExperienceRecordDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userExperienceRecordDo) Joins(fields ...field.RelationField) IUserExperienceRecordDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userExperienceRecordDo) Preload(fields ...field.RelationField) IUserExperienceRecordDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userExperienceRecordDo) FirstOrInit() (*model.UserExperienceRecord, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExperienceRecord), nil
	}
}

func (u userExperienceRecordDo) FirstOrCreate() (*model.UserExperienceRecord, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExperienceRecord), nil
	}
}

func (u userExperienceRecordDo) FindByPage(offset int, limit int) (result []*model.UserExperienceRecord, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userExperienceRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userExperienceRecordDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userExperienceRecordDo) Delete(models ...*model.UserExperienceRecord) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userExperienceRecordDo) withDO(do gen.Dao) *userExperienceRecordDo {
	u.DO = *do.(*gen.DO)
	return u
}
