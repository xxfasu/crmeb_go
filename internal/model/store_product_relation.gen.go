// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameStoreProductRelation = "eb_store_product_relation"

// StoreProductRelation 商品点赞和收藏表
type StoreProductRelation struct {
	ID        int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:id" json:"id"`       // id
	UID       int64                 `gorm:"column:uid;type:int unsigned;not null;comment:用户ID" json:"uid"`                        // 用户ID
	ProductID int64                 `gorm:"column:product_id;type:int unsigned;not null;comment:商品ID" json:"product_id"`          // 商品ID
	Type      string                `gorm:"column:type;type:varchar(32);not null;comment:类型(收藏(collect）、点赞(like))" json:"type"`   // 类型(收藏(collect）、点赞(like))
	Category  string                `gorm:"column:category;type:varchar(32);not null;comment:某种类型的商品(普通商品、秒杀商品)" json:"category"` // 某种类型的商品(普通商品、秒杀商品)
	CreatedAt int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreProductRelation's table name
func (*StoreProductRelation) TableName() string {
	return TableNameStoreProductRelation
}
