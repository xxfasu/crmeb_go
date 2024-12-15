// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameStoreBargainUser = "eb_store_bargain_user"

// StoreBargainUser 用户参与砍价表
type StoreBargainUser struct {
	ID              int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:用户参与砍价表ID" json:"id"`           // 用户参与砍价表ID
	UID             int64                 `gorm:"column:uid;type:int unsigned;comment:用户ID" json:"uid"`                                            // 用户ID
	BargainID       int64                 `gorm:"column:bargain_id;type:int unsigned;comment:砍价商品id" json:"bargain_id"`                            // 砍价商品id
	BargainPriceMin decimal.Decimal       `gorm:"column:bargain_price_min;type:decimal(8,2) unsigned;comment:砍价的最低价" json:"bargain_price_min"`     // 砍价的最低价
	BargainPrice    decimal.Decimal       `gorm:"column:bargain_price;type:decimal(8,2);comment:砍价金额" json:"bargain_price"`                        // 砍价金额
	Price           decimal.Decimal       `gorm:"column:price;type:decimal(8,2) unsigned;comment:砍掉的价格" json:"price"`                              // 砍掉的价格
	Status          int64                 `gorm:"column:status;type:tinyint unsigned;not null;comment:状态 1参与中 2 活动结束参与失败 3活动结束参与成功" json:"status"` // 状态 1参与中 2 活动结束参与失败 3活动结束参与成功
	AddTime         int64                 `gorm:"column:add_time;type:bigint unsigned;comment:参与时间" json:"add_time"`                               // 参与时间
	CreatedAt       int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt       int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreBargainUser's table name
func (*StoreBargainUser) TableName() string {
	return TableNameStoreBargainUser
}
