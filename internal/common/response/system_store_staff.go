package response

import "crmeb_go/internal/model"

type SystemStoreStaffResp struct {
	ID           int64              `json:"id"`            // id
	UID          int                `json:"uid"`           // 微信用户id
	Avatar       string             `json:"avatar"`        // 店员头像
	User         *model.User        `json:"user"`          // 用户信息
	StoreID      int64              `json:"store_id"`      // 门店id
	SystemStore  *model.SystemStore `json:"system_store"`  // 门店信息
	StaffName    string             `json:"staff_name"`    // 店员名称
	Phone        string             `json:"phone"`         // 手机号码
	VerifyStatus int64              `json:"verify_status"` // 核销开关
	Status       int64              `json:"status"`        // 状态
	CreateTime   string             `json:"create_time"`   // 创建时间
	UpdateTime   string             `json:"update_time"`   // 更新时间
}
