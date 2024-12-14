// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameStoreOrder = "eb_store_order"

// StoreOrder 订单表
type StoreOrder struct {
	ID                     int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:订单ID" json:"id"`                            // 订单ID
	OrderID                string                `gorm:"column:order_id;type:varchar(32);not null;comment:订单号" json:"order_id"`                                       // 订单号
	UID                    int64                 `gorm:"column:uid;type:int unsigned;not null;comment:用户id" json:"uid"`                                               // 用户id
	RealName               string                `gorm:"column:real_name;type:varchar(32);not null;comment:用户姓名" json:"real_name"`                                    // 用户姓名
	UserPhone              string                `gorm:"column:user_phone;type:varchar(18);not null;comment:用户电话" json:"user_phone"`                                  // 用户电话
	UserAddress            string                `gorm:"column:user_address;type:varchar(100);not null;comment:详细地址" json:"user_address"`                             // 详细地址
	FreightPrice           decimal.Decimal       `gorm:"column:freight_price;type:decimal(8,2);not null;default:0.00;comment:运费金额" json:"freight_price"`              // 运费金额
	TotalNum               int64                 `gorm:"column:total_num;type:int unsigned;not null;comment:订单商品总数" json:"total_num"`                                 // 订单商品总数
	TotalPrice             decimal.Decimal       `gorm:"column:total_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:订单总价" json:"total_price"`         // 订单总价
	TotalPostage           decimal.Decimal       `gorm:"column:total_postage;type:decimal(8,2) unsigned;not null;default:0.00;comment:邮费" json:"total_postage"`       // 邮费
	PayPrice               decimal.Decimal       `gorm:"column:pay_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:实际支付金额" json:"pay_price"`           // 实际支付金额
	PayPostage             decimal.Decimal       `gorm:"column:pay_postage;type:decimal(8,2) unsigned;not null;default:0.00;comment:支付邮费" json:"pay_postage"`         // 支付邮费
	DeductionPrice         decimal.Decimal       `gorm:"column:deduction_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:抵扣金额" json:"deduction_price"` // 抵扣金额
	CouponID               int64                 `gorm:"column:coupon_id;type:int unsigned;not null;comment:优惠券id" json:"coupon_id"`                                  // 优惠券id
	CouponPrice            decimal.Decimal       `gorm:"column:coupon_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:优惠券金额" json:"coupon_price"`      // 优惠券金额
	Paid                   int32                 `gorm:"column:paid;type:tinyint unsigned;not null;comment:支付状态" json:"paid"`                                         // 支付状态
	PayTime                int64                 `gorm:"column:pay_time;type:bigint" json:"pay_time"`
	PayType                string                `gorm:"column:pay_type;type:varchar(32);not null;comment:支付方式" json:"pay_type"`                                   // 支付方式
	Status                 bool                  `gorm:"column:status;type:tinyint(1);not null;comment:订单状态（0：待发货；1：待收货；2：已收货，待评价；3：已完成；）" json:"status"`          // 订单状态（0：待发货；1：待收货；2：已收货，待评价；3：已完成；）
	RefundStatus           int32                 `gorm:"column:refund_status;type:tinyint unsigned;not null;comment:0 未退款 1 申请中 2 已退款 3 退款中" json:"refund_status"` // 0 未退款 1 申请中 2 已退款 3 退款中
	RefundReasonWapImg     string                `gorm:"column:refund_reason_wap_img;type:varchar(5000);comment:退款图片" json:"refund_reason_wap_img"`                // 退款图片
	RefundReasonWapExplain string                `gorm:"column:refund_reason_wap_explain;type:varchar(255);comment:退款用户说明" json:"refund_reason_wap_explain"`       // 退款用户说明
	RefundReasonWap        string                `gorm:"column:refund_reason_wap;type:varchar(255);comment:前台退款原因" json:"refund_reason_wap"`                       // 前台退款原因
	RefundReason           string                `gorm:"column:refund_reason;type:varchar(255);comment:不退款的理由" json:"refund_reason"`                               // 不退款的理由
	RefundReasonTime       int64                 `gorm:"column:refund_reason_time;type:bigint" json:"refund_reason_time"`
	RefundPrice            decimal.Decimal       `gorm:"column:refund_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:退款金额" json:"refund_price"` // 退款金额
	DeliveryName           string                `gorm:"column:delivery_name;type:varchar(64);comment:快递名称/送货人姓名" json:"delivery_name"`                         // 快递名称/送货人姓名
	DeliveryType           string                `gorm:"column:delivery_type;type:varchar(32);comment:发货类型" json:"delivery_type"`                               // 发货类型
	DeliveryID             string                `gorm:"column:delivery_id;type:varchar(64);comment:快递单号/手机号" json:"delivery_id"`                               // 快递单号/手机号
	GainIntegral           int64                 `gorm:"column:gain_integral;type:int;comment:消费赚取积分" json:"gain_integral"`                                     // 消费赚取积分
	UseIntegral            int64                 `gorm:"column:use_integral;type:int;comment:使用积分" json:"use_integral"`                                         // 使用积分
	BackIntegral           int64                 `gorm:"column:back_integral;type:int;comment:给用户退了多少积分" json:"back_integral"`                                  // 给用户退了多少积分
	Mark                   string                `gorm:"column:mark;type:varchar(512);not null;comment:备注" json:"mark"`                                         // 备注
	IsDel                  int32                 `gorm:"column:is_del;type:tinyint unsigned;not null;comment:是否删除" json:"is_del"`                               // 是否删除
	Remark                 string                `gorm:"column:remark;type:varchar(512);comment:管理员备注" json:"remark"`                                           // 管理员备注
	MerID                  int64                 `gorm:"column:mer_id;type:int unsigned;not null;comment:商户ID" json:"mer_id"`                                   // 商户ID
	IsMerCheck             int32                 `gorm:"column:is_mer_check;type:tinyint unsigned;not null" json:"is_mer_check"`
	CombinationID          int64                 `gorm:"column:combination_id;type:int unsigned;comment:拼团商品id0一般商品" json:"combination_id"`                                // 拼团商品id0一般商品
	PinkID                 int64                 `gorm:"column:pink_id;type:int unsigned;not null;comment:拼团id 0没有拼团" json:"pink_id"`                                      // 拼团id 0没有拼团
	Cost                   decimal.Decimal       `gorm:"column:cost;type:decimal(8,2) unsigned;not null;comment:成本价" json:"cost"`                                          // 成本价
	SeckillID              int64                 `gorm:"column:seckill_id;type:int unsigned;not null;comment:秒杀商品ID" json:"seckill_id"`                                    // 秒杀商品ID
	BargainID              int64                 `gorm:"column:bargain_id;type:int unsigned;comment:砍价id" json:"bargain_id"`                                               // 砍价id
	VerifyCode             string                `gorm:"column:verify_code;type:varchar(12);not null;comment:核销码" json:"verify_code"`                                      // 核销码
	StoreID                int64                 `gorm:"column:store_id;type:int;not null;comment:门店id" json:"store_id"`                                                   // 门店id
	ShippingType           bool                  `gorm:"column:shipping_type;type:tinyint(1);not null;default:1;comment:配送方式 1=快递 ，2=门店自提" json:"shipping_type"`           // 配送方式 1=快递 ，2=门店自提
	ClerkID                int64                 `gorm:"column:clerk_id;type:int;not null;comment:店员id/核销员id" json:"clerk_id"`                                             // 店员id/核销员id
	IsChannel              int32                 `gorm:"column:is_channel;type:tinyint unsigned;comment:支付渠道(0微信公众号1微信小程序2余额)" json:"is_channel"`                          // 支付渠道(0微信公众号1微信小程序2余额)
	IsRemind               int32                 `gorm:"column:is_remind;type:tinyint unsigned;comment:消息提醒" json:"is_remind"`                                             // 消息提醒
	IsSystemDel            bool                  `gorm:"column:is_system_del;type:tinyint(1);comment:后台是否删除" json:"is_system_del"`                                         // 后台是否删除
	DeliveryCode           string                `gorm:"column:delivery_code;type:varchar(50);comment:快递公司简称" json:"delivery_code"`                                        // 快递公司简称
	BargainUserID          int64                 `gorm:"column:bargain_user_id;type:int unsigned;not null;comment:用户拼团活动id 0没有" json:"bargain_user_id"`                    // 用户拼团活动id 0没有
	Type                   int64                 `gorm:"column:type;type:int;not null;comment:订单类型:0-普通订单，1-视频号订单" json:"type"`                                            // 订单类型:0-普通订单，1-视频号订单
	ProTotalPrice          decimal.Decimal       `gorm:"column:pro_total_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:商品总价" json:"pro_total_price"`      // 商品总价
	BeforePayPrice         decimal.Decimal       `gorm:"column:before_pay_price;type:decimal(8,2) unsigned;not null;default:0.00;comment:改价前支付金额" json:"before_pay_price"` // 改价前支付金额
	IsAlterPrice           bool                  `gorm:"column:is_alter_price;type:tinyint(1);not null;comment:是否改价,0-否，1-是" json:"is_alter_price"`                        // 是否改价,0-否，1-是
	OutTradeNo             string                `gorm:"column:out_trade_no;type:varchar(32);comment:商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号" json:"out_trade_no"`             // 商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号
	CreatedAt              int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt              int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt              soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreOrder's table name
func (*StoreOrder) TableName() string {
	return TableNameStoreOrder
}
