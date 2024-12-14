// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameUserBill = "eb_user_bill"

// UserBill 用户账单表
type UserBill struct {
	ID        int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:用户账单id" json:"id"`            // 用户账单id
	UID       int64                 `gorm:"column:uid;type:int unsigned;not null;comment:用户uid" json:"uid"`                                // 用户uid
	LinkID    string                `gorm:"column:link_id;type:varchar(32);not null;default:0;comment:关联id" json:"link_id"`                // 关联id
	Pm        int32                 `gorm:"column:pm;type:tinyint unsigned;not null;comment:0 = 支出 1 = 获得" json:"pm"`                      // 0 = 支出 1 = 获得
	Title     string                `gorm:"column:title;type:varchar(64);not null;comment:账单标题" json:"title"`                              // 账单标题
	Category  string                `gorm:"column:category;type:varchar(64);not null;comment:明细种类" json:"category"`                        // 明细种类
	Type      string                `gorm:"column:type;type:varchar(64);not null;comment:明细类型" json:"type"`                                // 明细类型
	Number    decimal.Decimal       `gorm:"column:number;type:decimal(8,2) unsigned;not null;default:0.00;comment:明细数字" json:"number"`     // 明细数字
	Balance   decimal.Decimal       `gorm:"column:balance;type:decimal(16,2) unsigned;not null;default:0.00;comment:剩余" json:"balance"`    // 剩余
	Mark      string                `gorm:"column:mark;type:varchar(512);not null;comment:备注" json:"mark"`                                 // 备注
	Status    bool                  `gorm:"column:status;type:tinyint(1);not null;default:1;comment:0 = 带确定 1 = 有效 -1 = 无效" json:"status"` // 0 = 带确定 1 = 有效 -1 = 无效
	CreatedAt int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName UserBill's table name
func (*UserBill) TableName() string {
	return TableNameUserBill
}
