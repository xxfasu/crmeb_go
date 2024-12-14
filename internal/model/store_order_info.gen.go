// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameStoreOrderInfo = "eb_store_order_info"

// StoreOrderInfo 订单购物详情表
type StoreOrderInfo struct {
	ID           int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键" json:"id"`                            // 主键
	OrderID      int64                 `gorm:"column:order_id;type:int unsigned;not null;comment:订单id" json:"order_id"`                          // 订单id
	ProductID    int64                 `gorm:"column:product_id;type:int unsigned;not null;comment:商品ID" json:"product_id"`                      // 商品ID
	Info         string                `gorm:"column:info;type:text;not null;comment:购买东西的详细信息" json:"info"`                                     // 购买东西的详细信息
	Unique       string                `gorm:"column:unique;type:char(32);not null;comment:唯一id" json:"unique"`                                  // 唯一id
	OrderNo      string                `gorm:"column:order_no;type:varchar(32);not null;comment:订单号" json:"order_no"`                            // 订单号
	ProductName  string                `gorm:"column:product_name;type:varchar(128);not null;comment:商品名称" json:"product_name"`                  // 商品名称
	AttrValueID  int64                 `gorm:"column:attr_value_id;type:int unsigned;comment:规格属性值id" json:"attr_value_id"`                      // 规格属性值id
	Image        string                `gorm:"column:image;type:varchar(256);not null;comment:商品图片" json:"image"`                                // 商品图片
	Sku          string                `gorm:"column:sku;type:varchar(128);not null;comment:商品sku" json:"sku"`                                   // 商品sku
	Price        decimal.Decimal       `gorm:"column:price;type:decimal(8,2) unsigned;not null;comment:商品价格" json:"price"`                       // 商品价格
	PayNum       int64                 `gorm:"column:pay_num;type:int unsigned;not null;comment:购买数量" json:"pay_num"`                            // 购买数量
	Weight       decimal.Decimal       `gorm:"column:weight;type:decimal(8,2) unsigned;not null;comment:重量" json:"weight"`                       // 重量
	Volume       decimal.Decimal       `gorm:"column:volume;type:decimal(8,2) unsigned;not null;comment:体积" json:"volume"`                       // 体积
	GiveIntegral int64                 `gorm:"column:give_integral;type:int unsigned;not null;comment:赠送积分" json:"give_integral"`                // 赠送积分
	IsReply      bool                  `gorm:"column:is_reply;type:tinyint(1);not null;comment:是否评价，0-未评价，1-已评价" json:"is_reply"`                // 是否评价，0-未评价，1-已评价
	IsSub        bool                  `gorm:"column:is_sub;type:tinyint(1);not null;comment:是否单独分佣,0-否，1-是" json:"is_sub"`                      // 是否单独分佣,0-否，1-是
	VipPrice     decimal.Decimal       `gorm:"column:vip_price;type:decimal(8,2) unsigned;not null;comment:会员价" json:"vip_price"`                // 会员价
	ProductType  int64                 `gorm:"column:product_type;type:int;not null;comment:商品类型:0-普通，1-秒杀，2-砍价，3-拼团，4-视频号" json:"product_type"` // 商品类型:0-普通，1-秒杀，2-砍价，3-拼团，4-视频号
	CreatedAt    int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt    int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt    soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreOrderInfo's table name
func (*StoreOrderInfo) TableName() string {
	return TableNameStoreOrderInfo
}
