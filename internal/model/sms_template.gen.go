// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSmsTemplate = "eb_sms_template"

// SmsTemplate 短信模板表
type SmsTemplate struct {
	ID        int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:id" json:"id"`   // id
	TempID    string                `gorm:"column:temp_id;type:varchar(20);not null;default:0;comment:短信模板id" json:"temp_id"` // 短信模板id
	TempType  int64                 `gorm:"column:temp_type;type:tinyint;not null;default:1;comment:模板类型" json:"temp_type"`   // 模板类型
	Title     string                `gorm:"column:title;type:varchar(100);not null;comment:模板说明" json:"title"`                // 模板说明
	Type      string                `gorm:"column:type;type:varchar(20);not null;comment:类型" json:"type"`                     // 类型
	TempKey   string                `gorm:"column:temp_key;type:varchar(50);not null;comment:模板编号" json:"temp_key"`           // 模板编号
	Status    int64                 `gorm:"column:status;type:tinyint;not null;default:1;comment:状态" json:"status"`           // 状态
	Content   string                `gorm:"column:content;type:varchar(500);not null;comment:短息内容" json:"content"`            // 短息内容
	CreatedAt int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName SmsTemplate's table name
func (*SmsTemplate) TableName() string {
	return TableNameSmsTemplate
}
