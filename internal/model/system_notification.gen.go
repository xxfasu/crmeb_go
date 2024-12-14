// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSystemNotification = "eb_system_notification"

// SystemNotification 通知设置表
type SystemNotification struct {
	ID          int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:id" json:"id"`             // id
	Mark        string                `gorm:"column:mark;type:varchar(50);not null;comment:标识" json:"mark"`                               // 标识
	Type        string                `gorm:"column:type;type:varchar(50);not null;comment:通知类型" json:"type"`                             // 通知类型
	Description string                `gorm:"column:description;type:varchar(100);not null;comment:通知场景说明" json:"description"`            // 通知场景说明
	IsWechat    int32                 `gorm:"column:is_wechat;type:tinyint;not null;comment:公众号模板消息（0：不存在，1：开启，2：关闭）" json:"is_wechat"`   // 公众号模板消息（0：不存在，1：开启，2：关闭）
	WechatID    int64                 `gorm:"column:wechat_id;type:int;not null;comment:模板消息id" json:"wechat_id"`                         // 模板消息id
	IsRoutine   int32                 `gorm:"column:is_routine;type:tinyint;not null;comment:小程序订阅消息（0：不存在，1：开启，2：关闭）" json:"is_routine"` // 小程序订阅消息（0：不存在，1：开启，2：关闭）
	RoutineID   int64                 `gorm:"column:routine_id;type:int;not null;comment:订阅消息id" json:"routine_id"`                       // 订阅消息id
	IsSms       int32                 `gorm:"column:is_sms;type:tinyint;not null;comment:发送短信（0：不存在，1：开启，2：关闭）" json:"is_sms"`            // 发送短信（0：不存在，1：开启，2：关闭）
	SmsID       int64                 `gorm:"column:sms_id;type:int;not null;comment:短信id" json:"sms_id"`                                 // 短信id
	SendType    int32                 `gorm:"column:send_type;type:tinyint;not null;default:1;comment:发送类型（1：用户，2：管理员）" json:"send_type"` // 发送类型（1：用户，2：管理员）
	CreatedAt   int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt   int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName SystemNotification's table name
func (*SystemNotification) TableName() string {
	return TableNameSystemNotification
}
