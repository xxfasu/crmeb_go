// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameUserGroup = "eb_user_group"

// UserGroup 用户分组表
type UserGroup struct {
	ID        int64                 `gorm:"column:id;type:smallint unsigned;primaryKey;autoIncrement:true" json:"id"`
	GroupName string                `gorm:"column:group_name;type:varchar(64);comment:用户分组名称" json:"group_name"` // 用户分组名称
	CreatedAt int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName UserGroup's table name
func (*UserGroup) TableName() string {
	return TableNameUserGroup
}
