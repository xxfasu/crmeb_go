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

func newUserToken(db *gorm.DB, opts ...gen.DOOption) userToken {
	_userToken := userToken{}

	_userToken.userTokenDo.UseDB(db, opts...)
	_userToken.userTokenDo.UseModel(&model.UserToken{})

	tableName := _userToken.userTokenDo.TableName()
	_userToken.ALL = field.NewAsterisk(tableName)
	_userToken.ID = field.NewInt64(tableName, "id")
	_userToken.UID = field.NewInt64(tableName, "uid")
	_userToken.Token = field.NewString(tableName, "token")
	_userToken.Type = field.NewInt64(tableName, "type")
	_userToken.ExpiresTime = field.NewTime(tableName, "expires_time")
	_userToken.LoginIP = field.NewString(tableName, "login_ip")
	_userToken.CreatedAt = field.NewInt64(tableName, "created_at")
	_userToken.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_userToken.DeletedAt = field.NewField(tableName, "deleted_at")

	_userToken.fillFieldMap()

	return _userToken
}

type userToken struct {
	userTokenDo userTokenDo

	ALL         field.Asterisk
	ID          field.Int64
	UID         field.Int64  // 用户 id
	Token       field.String // token
	Type        field.Int64  // 类型，1 公众号， 2 小程序, 3 unionid, 5AppIos,6AppAndroid,7ios
	ExpiresTime field.Time   // 到期时间
	LoginIP     field.String // 登录ip
	CreatedAt   field.Int64
	UpdatedAt   field.Int64
	DeletedAt   field.Field

	fieldMap map[string]field.Expr
}

func (u userToken) Table(newTableName string) *userToken {
	u.userTokenDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userToken) As(alias string) *userToken {
	u.userTokenDo.DO = *(u.userTokenDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userToken) updateTableName(table string) *userToken {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UID = field.NewInt64(table, "uid")
	u.Token = field.NewString(table, "token")
	u.Type = field.NewInt64(table, "type")
	u.ExpiresTime = field.NewTime(table, "expires_time")
	u.LoginIP = field.NewString(table, "login_ip")
	u.CreatedAt = field.NewInt64(table, "created_at")
	u.UpdatedAt = field.NewInt64(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userToken) WithContext(ctx context.Context) IUserTokenDo {
	return u.userTokenDo.WithContext(ctx)
}

func (u userToken) TableName() string { return u.userTokenDo.TableName() }

func (u userToken) Alias() string { return u.userTokenDo.Alias() }

func (u userToken) Columns(cols ...field.Expr) gen.Columns { return u.userTokenDo.Columns(cols...) }

func (u *userToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userToken) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["token"] = u.Token
	u.fieldMap["type"] = u.Type
	u.fieldMap["expires_time"] = u.ExpiresTime
	u.fieldMap["login_ip"] = u.LoginIP
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userToken) clone(db *gorm.DB) userToken {
	u.userTokenDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userToken) replaceDB(db *gorm.DB) userToken {
	u.userTokenDo.ReplaceDB(db)
	return u
}

type userTokenDo struct{ gen.DO }

type IUserTokenDo interface {
	gen.SubQuery
	Debug() IUserTokenDo
	WithContext(ctx context.Context) IUserTokenDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserTokenDo
	WriteDB() IUserTokenDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserTokenDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserTokenDo
	Not(conds ...gen.Condition) IUserTokenDo
	Or(conds ...gen.Condition) IUserTokenDo
	Select(conds ...field.Expr) IUserTokenDo
	Where(conds ...gen.Condition) IUserTokenDo
	Order(conds ...field.Expr) IUserTokenDo
	Distinct(cols ...field.Expr) IUserTokenDo
	Omit(cols ...field.Expr) IUserTokenDo
	Join(table schema.Tabler, on ...field.Expr) IUserTokenDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserTokenDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserTokenDo
	Group(cols ...field.Expr) IUserTokenDo
	Having(conds ...gen.Condition) IUserTokenDo
	Limit(limit int) IUserTokenDo
	Offset(offset int) IUserTokenDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTokenDo
	Unscoped() IUserTokenDo
	Create(values ...*model.UserToken) error
	CreateInBatches(values []*model.UserToken, batchSize int) error
	Save(values ...*model.UserToken) error
	First() (*model.UserToken, error)
	Take() (*model.UserToken, error)
	Last() (*model.UserToken, error)
	Find() ([]*model.UserToken, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserToken, err error)
	FindInBatches(result *[]*model.UserToken, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserToken) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserTokenDo
	Assign(attrs ...field.AssignExpr) IUserTokenDo
	Joins(fields ...field.RelationField) IUserTokenDo
	Preload(fields ...field.RelationField) IUserTokenDo
	FirstOrInit() (*model.UserToken, error)
	FirstOrCreate() (*model.UserToken, error)
	FindByPage(offset int, limit int) (result []*model.UserToken, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserTokenDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userTokenDo) Debug() IUserTokenDo {
	return u.withDO(u.DO.Debug())
}

func (u userTokenDo) WithContext(ctx context.Context) IUserTokenDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userTokenDo) ReadDB() IUserTokenDo {
	return u.Clauses(dbresolver.Read)
}

func (u userTokenDo) WriteDB() IUserTokenDo {
	return u.Clauses(dbresolver.Write)
}

func (u userTokenDo) Session(config *gorm.Session) IUserTokenDo {
	return u.withDO(u.DO.Session(config))
}

func (u userTokenDo) Clauses(conds ...clause.Expression) IUserTokenDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userTokenDo) Returning(value interface{}, columns ...string) IUserTokenDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userTokenDo) Not(conds ...gen.Condition) IUserTokenDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userTokenDo) Or(conds ...gen.Condition) IUserTokenDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userTokenDo) Select(conds ...field.Expr) IUserTokenDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userTokenDo) Where(conds ...gen.Condition) IUserTokenDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userTokenDo) Order(conds ...field.Expr) IUserTokenDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userTokenDo) Distinct(cols ...field.Expr) IUserTokenDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userTokenDo) Omit(cols ...field.Expr) IUserTokenDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userTokenDo) Join(table schema.Tabler, on ...field.Expr) IUserTokenDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserTokenDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserTokenDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userTokenDo) Group(cols ...field.Expr) IUserTokenDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userTokenDo) Having(conds ...gen.Condition) IUserTokenDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userTokenDo) Limit(limit int) IUserTokenDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userTokenDo) Offset(offset int) IUserTokenDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTokenDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userTokenDo) Unscoped() IUserTokenDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userTokenDo) Create(values ...*model.UserToken) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userTokenDo) CreateInBatches(values []*model.UserToken, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userTokenDo) Save(values ...*model.UserToken) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userTokenDo) First() (*model.UserToken, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserToken), nil
	}
}

func (u userTokenDo) Take() (*model.UserToken, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserToken), nil
	}
}

func (u userTokenDo) Last() (*model.UserToken, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserToken), nil
	}
}

func (u userTokenDo) Find() ([]*model.UserToken, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserToken), err
}

func (u userTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserToken, err error) {
	buf := make([]*model.UserToken, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userTokenDo) FindInBatches(result *[]*model.UserToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userTokenDo) Attrs(attrs ...field.AssignExpr) IUserTokenDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userTokenDo) Assign(attrs ...field.AssignExpr) IUserTokenDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userTokenDo) Joins(fields ...field.RelationField) IUserTokenDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userTokenDo) Preload(fields ...field.RelationField) IUserTokenDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userTokenDo) FirstOrInit() (*model.UserToken, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserToken), nil
	}
}

func (u userTokenDo) FirstOrCreate() (*model.UserToken, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserToken), nil
	}
}

func (u userTokenDo) FindByPage(offset int, limit int) (result []*model.UserToken, count int64, err error) {
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

func (u userTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userTokenDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userTokenDo) Delete(models ...*model.UserToken) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userTokenDo) withDO(do gen.Dao) *userTokenDo {
	u.DO = *do.(*gen.DO)
	return u
}
