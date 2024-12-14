// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSystemStoreStaff = "eb_system_store_staff"

// SystemStoreStaff 门店店员表
type SystemStoreStaff struct {
	ID           int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	UID          int64                 `gorm:"column:uid;type:int unsigned;not null;comment:管理员id" json:"uid"`               // 管理员id
	Avatar       string                `gorm:"column:avatar;type:varchar(255);not null;comment:店员头像" json:"avatar"`          // 店员头像
	StoreID      int64                 `gorm:"column:store_id;type:int;not null;comment:门店id" json:"store_id"`               // 门店id
	StaffName    string                `gorm:"column:staff_name;type:varchar(64);comment:店员名称" json:"staff_name"`            // 店员名称
	Phone        string                `gorm:"column:phone;type:char(15);comment:手机号码" json:"phone"`                         // 手机号码
	VerifyStatus int32                 `gorm:"column:verify_status;type:tinyint;not null;comment:核销开关" json:"verify_status"` // 核销开关
	Status       int32                 `gorm:"column:status;type:tinyint;default:1;comment:状态" json:"status"`                // 状态
	CreatedAt    int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt    int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt    soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName SystemStoreStaff's table name
func (*SystemStoreStaff) TableName() string {
	return TableNameSystemStoreStaff
}
