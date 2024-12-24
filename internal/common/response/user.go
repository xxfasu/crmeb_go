package response

import (
	"crmeb_go/internal/model"
	"crmeb_go/pkg/util"
	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
)

type User struct {
	ID             int64           `json:"uid"`            // 用户id
	Account        string          `json:"account"`        // 用户账号
	Pwd            string          `json:"pwd"`            // 用户密码
	RealName       string          `json:"realName"`       // 真实姓名
	Birthday       string          `json:"birthday"`       // 生日
	CardID         string          `json:"cardId"`         // 身份证号码
	Mark           string          `json:"mark"`           // 用户备注
	PartnerID      int64           `json:"partnerId"`      // 合伙人id
	GroupID        string          `json:"groupId"`        // 用户分组id
	TagID          string          `json:"tagId"`          // 标签id
	Nickname       string          `json:"nickname"`       // 用户昵称
	Avatar         string          `json:"avatar"`         // 用户头像
	Phone          string          `json:"phone"`          // 手机号码
	Sex            int64           `json:"sex"`            // 性别，0未知，1男，2女，3保密
	Country        string          `json:"country"`        // 国家，中国CN，其他OTHER
	AddIP          string          `json:"addIp"`          // 添加ip
	LastIP         string          `json:"lastIp"`         // 最后一次登录ip
	NowMoney       decimal.Decimal `json:"nowMoney"`       // 用户余额
	BrokeragePrice decimal.Decimal `json:"brokeragePrice"` // 佣金金额
	Integral       int64           `json:"integral"`       // 用户剩余积分
	Experience     int64           `json:"experience"`     // 用户剩余经验
	SignNum        int64           `json:"signNum"`        // 连续签到天数
	Status         bool            `json:"status"`         // 1为正常，0为禁止
	Level          int64           `json:"level"`          // 等级
	SpreadUID      int64           `json:"spreadUid"`      // 推广员id
	SpreadTime     string          `json:"spreadTime"`     // 推广员关联时间
	UserType       string          `json:"userType"`       // 用户类型
	IsPromoter     bool            `json:"isPromoter"`     // 是否为推广员
	PayCount       int64           `json:"payCount"`       // 用户购买次数
	SpreadCount    int64           `json:"spreadCount"`    // 下级人数
	Addres         string          `json:"addres"`         // 详细地址
	Adminid        int64           `json:"adminid"`        // 管理员编号
	LoginType      string          `json:"loginType"`      // 用户登陆类型，h5,wechat,routine
	CreateTime     string          `json:"createTime"`     // 创建时间
	UpdateTime     string          `json:"updateTime"`     // 更新时间
	LastLoginTime  string          `json:"lastLoginTime"`  // 最后一次登录时间
	CleanTime      string          `json:"cleanTime"`      // 最后一次登录时间
	Path           string          `json:"path"`           // 推广等级记录
	Subscribe      int64           `json:"subscribe"`      // 是否关注公众号
	PromoterTime   string          `json:"promoterTime"`   // 成为分销员时间
}

func (u *User) ConvertFromModel(m *model.User) error {
	return copier.CopyWithOption(u, m, util.ConvertCopyOption(*u, *m))
}
