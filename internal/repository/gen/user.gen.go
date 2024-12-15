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

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(&model.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewAsterisk(tableName)
	_user.UID = field.NewInt64(tableName, "uid")
	_user.Account = field.NewString(tableName, "account")
	_user.Pwd = field.NewString(tableName, "pwd")
	_user.RealName = field.NewString(tableName, "real_name")
	_user.Birthday = field.NewString(tableName, "birthday")
	_user.CardID = field.NewString(tableName, "card_id")
	_user.Mark = field.NewString(tableName, "mark")
	_user.PartnerID = field.NewInt64(tableName, "partner_id")
	_user.GroupID = field.NewString(tableName, "group_id")
	_user.TagID = field.NewString(tableName, "tag_id")
	_user.Nickname = field.NewString(tableName, "nickname")
	_user.Avatar = field.NewString(tableName, "avatar")
	_user.Phone = field.NewString(tableName, "phone")
	_user.AddIP = field.NewString(tableName, "add_ip")
	_user.LastIP = field.NewString(tableName, "last_ip")
	_user.NowMoney = field.NewField(tableName, "now_money")
	_user.BrokeragePrice = field.NewField(tableName, "brokerage_price")
	_user.Integral = field.NewInt64(tableName, "integral")
	_user.Experience = field.NewInt64(tableName, "experience")
	_user.SignNum = field.NewInt64(tableName, "sign_num")
	_user.Status = field.NewInt64(tableName, "status")
	_user.Level = field.NewInt64(tableName, "level")
	_user.SpreadUID = field.NewInt64(tableName, "spread_uid")
	_user.SpreadTime = field.NewInt64(tableName, "spread_time")
	_user.UserType = field.NewString(tableName, "user_type")
	_user.IsPromoter = field.NewInt64(tableName, "is_promoter")
	_user.PayCount = field.NewInt64(tableName, "pay_count")
	_user.SpreadCount = field.NewInt64(tableName, "spread_count")
	_user.Addres = field.NewString(tableName, "addres")
	_user.Adminid = field.NewInt64(tableName, "adminid")
	_user.LoginType = field.NewString(tableName, "login_type")
	_user.LastLoginTime = field.NewInt64(tableName, "last_login_time")
	_user.CleanTime = field.NewInt64(tableName, "clean_time")
	_user.Path = field.NewString(tableName, "path")
	_user.Subscribe = field.NewInt64(tableName, "subscribe")
	_user.SubscribeTime = field.NewInt64(tableName, "subscribe_time")
	_user.Sex = field.NewInt64(tableName, "sex")
	_user.Country = field.NewString(tableName, "country")
	_user.PromoterTime = field.NewInt64(tableName, "promoter_time")
	_user.CreatedAt = field.NewInt64(tableName, "created_at")
	_user.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")

	_user.fillFieldMap()

	return _user
}

// user 用户表
type user struct {
	userDo userDo

	ALL            field.Asterisk
	UID            field.Int64  // 用户id
	Account        field.String // 用户账号
	Pwd            field.String // 用户密码
	RealName       field.String // 真实姓名
	Birthday       field.String // 生日
	CardID         field.String // 身份证号码
	Mark           field.String // 用户备注
	PartnerID      field.Int64  // 合伙人id
	GroupID        field.String // 用户分组id
	TagID          field.String // 标签id
	Nickname       field.String // 用户昵称
	Avatar         field.String // 用户头像
	Phone          field.String // 手机号码
	AddIP          field.String // 添加ip
	LastIP         field.String // 最后一次登录ip
	NowMoney       field.Field  // 用户余额
	BrokeragePrice field.Field  // 佣金金额
	Integral       field.Int64  // 用户剩余积分
	Experience     field.Int64  // 用户剩余经验
	SignNum        field.Int64  // 连续签到天数
	Status         field.Int64  // 1为正常，0为禁止
	Level          field.Int64  // 等级
	SpreadUID      field.Int64  // 推广员id
	SpreadTime     field.Int64
	UserType       field.String // 用户类型
	IsPromoter     field.Int64  // 是否为推广员
	PayCount       field.Int64  // 用户购买次数
	SpreadCount    field.Int64  // 下级人数
	Addres         field.String // 详细地址
	Adminid        field.Int64  // 管理员编号
	LoginType      field.String // 用户登陆类型，h5,wechat,routine
	LastLoginTime  field.Int64
	CleanTime      field.Int64
	Path           field.String // 推广等级记录
	Subscribe      field.Int64  // 是否关注公众号
	SubscribeTime  field.Int64
	Sex            field.Int64  // 性别，0未知，1男，2女，3保密
	Country        field.String // 国家，中国CN，其他OTHER
	PromoterTime   field.Int64
	CreatedAt      field.Int64
	UpdatedAt      field.Int64
	DeletedAt      field.Field

	fieldMap map[string]field.Expr
}

func (u user) Table(newTableName string) *user {
	u.userDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user) updateTableName(table string) *user {
	u.ALL = field.NewAsterisk(table)
	u.UID = field.NewInt64(table, "uid")
	u.Account = field.NewString(table, "account")
	u.Pwd = field.NewString(table, "pwd")
	u.RealName = field.NewString(table, "real_name")
	u.Birthday = field.NewString(table, "birthday")
	u.CardID = field.NewString(table, "card_id")
	u.Mark = field.NewString(table, "mark")
	u.PartnerID = field.NewInt64(table, "partner_id")
	u.GroupID = field.NewString(table, "group_id")
	u.TagID = field.NewString(table, "tag_id")
	u.Nickname = field.NewString(table, "nickname")
	u.Avatar = field.NewString(table, "avatar")
	u.Phone = field.NewString(table, "phone")
	u.AddIP = field.NewString(table, "add_ip")
	u.LastIP = field.NewString(table, "last_ip")
	u.NowMoney = field.NewField(table, "now_money")
	u.BrokeragePrice = field.NewField(table, "brokerage_price")
	u.Integral = field.NewInt64(table, "integral")
	u.Experience = field.NewInt64(table, "experience")
	u.SignNum = field.NewInt64(table, "sign_num")
	u.Status = field.NewInt64(table, "status")
	u.Level = field.NewInt64(table, "level")
	u.SpreadUID = field.NewInt64(table, "spread_uid")
	u.SpreadTime = field.NewInt64(table, "spread_time")
	u.UserType = field.NewString(table, "user_type")
	u.IsPromoter = field.NewInt64(table, "is_promoter")
	u.PayCount = field.NewInt64(table, "pay_count")
	u.SpreadCount = field.NewInt64(table, "spread_count")
	u.Addres = field.NewString(table, "addres")
	u.Adminid = field.NewInt64(table, "adminid")
	u.LoginType = field.NewString(table, "login_type")
	u.LastLoginTime = field.NewInt64(table, "last_login_time")
	u.CleanTime = field.NewInt64(table, "clean_time")
	u.Path = field.NewString(table, "path")
	u.Subscribe = field.NewInt64(table, "subscribe")
	u.SubscribeTime = field.NewInt64(table, "subscribe_time")
	u.Sex = field.NewInt64(table, "sex")
	u.Country = field.NewString(table, "country")
	u.PromoterTime = field.NewInt64(table, "promoter_time")
	u.CreatedAt = field.NewInt64(table, "created_at")
	u.UpdatedAt = field.NewInt64(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *user) WithContext(ctx context.Context) IUserDo { return u.userDo.WithContext(ctx) }

func (u user) TableName() string { return u.userDo.TableName() }

func (u user) Alias() string { return u.userDo.Alias() }

func (u user) Columns(cols ...field.Expr) gen.Columns { return u.userDo.Columns(cols...) }

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 42)
	u.fieldMap["uid"] = u.UID
	u.fieldMap["account"] = u.Account
	u.fieldMap["pwd"] = u.Pwd
	u.fieldMap["real_name"] = u.RealName
	u.fieldMap["birthday"] = u.Birthday
	u.fieldMap["card_id"] = u.CardID
	u.fieldMap["mark"] = u.Mark
	u.fieldMap["partner_id"] = u.PartnerID
	u.fieldMap["group_id"] = u.GroupID
	u.fieldMap["tag_id"] = u.TagID
	u.fieldMap["nickname"] = u.Nickname
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["phone"] = u.Phone
	u.fieldMap["add_ip"] = u.AddIP
	u.fieldMap["last_ip"] = u.LastIP
	u.fieldMap["now_money"] = u.NowMoney
	u.fieldMap["brokerage_price"] = u.BrokeragePrice
	u.fieldMap["integral"] = u.Integral
	u.fieldMap["experience"] = u.Experience
	u.fieldMap["sign_num"] = u.SignNum
	u.fieldMap["status"] = u.Status
	u.fieldMap["level"] = u.Level
	u.fieldMap["spread_uid"] = u.SpreadUID
	u.fieldMap["spread_time"] = u.SpreadTime
	u.fieldMap["user_type"] = u.UserType
	u.fieldMap["is_promoter"] = u.IsPromoter
	u.fieldMap["pay_count"] = u.PayCount
	u.fieldMap["spread_count"] = u.SpreadCount
	u.fieldMap["addres"] = u.Addres
	u.fieldMap["adminid"] = u.Adminid
	u.fieldMap["login_type"] = u.LoginType
	u.fieldMap["last_login_time"] = u.LastLoginTime
	u.fieldMap["clean_time"] = u.CleanTime
	u.fieldMap["path"] = u.Path
	u.fieldMap["subscribe"] = u.Subscribe
	u.fieldMap["subscribe_time"] = u.SubscribeTime
	u.fieldMap["sex"] = u.Sex
	u.fieldMap["country"] = u.Country
	u.fieldMap["promoter_time"] = u.PromoterTime
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user) replaceDB(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userDo struct{ gen.DO }

type IUserDo interface {
	gen.SubQuery
	Debug() IUserDo
	WithContext(ctx context.Context) IUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserDo
	WriteDB() IUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserDo
	Not(conds ...gen.Condition) IUserDo
	Or(conds ...gen.Condition) IUserDo
	Select(conds ...field.Expr) IUserDo
	Where(conds ...gen.Condition) IUserDo
	Order(conds ...field.Expr) IUserDo
	Distinct(cols ...field.Expr) IUserDo
	Omit(cols ...field.Expr) IUserDo
	Join(table schema.Tabler, on ...field.Expr) IUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserDo
	Group(cols ...field.Expr) IUserDo
	Having(conds ...gen.Condition) IUserDo
	Limit(limit int) IUserDo
	Offset(offset int) IUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo
	Unscoped() IUserDo
	Create(values ...*model.User) error
	CreateInBatches(values []*model.User, batchSize int) error
	Save(values ...*model.User) error
	First() (*model.User, error)
	Take() (*model.User, error)
	Last() (*model.User, error)
	Find() ([]*model.User, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User, err error)
	FindInBatches(result *[]*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.User) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserDo
	Assign(attrs ...field.AssignExpr) IUserDo
	Joins(fields ...field.RelationField) IUserDo
	Preload(fields ...field.RelationField) IUserDo
	FirstOrInit() (*model.User, error)
	FirstOrCreate() (*model.User, error)
	FindByPage(offset int, limit int) (result []*model.User, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userDo) Debug() IUserDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) IUserDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) ReadDB() IUserDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDo) WriteDB() IUserDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDo) Session(config *gorm.Session) IUserDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDo) Clauses(conds ...clause.Expression) IUserDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Returning(value interface{}, columns ...string) IUserDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDo) Not(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) IUserDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) IUserDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() IUserDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*model.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*model.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Take() (*model.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Last() (*model.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Find() ([]*model.User, error) {
	result, err := u.DO.Find()
	return result.([]*model.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User, err error) {
	buf := make([]*model.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDo) Preload(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDo) FirstOrInit() (*model.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FirstOrCreate() (*model.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*model.User, count int64, err error) {
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

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDo) Delete(models ...*model.User) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}
