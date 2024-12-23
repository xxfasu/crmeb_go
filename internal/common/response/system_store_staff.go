package response

import (
	"crmeb_go/internal/model"
	"crmeb_go/pkg/util"
	"github.com/jinzhu/copier"
)

type SystemStoreStaff struct {
	ID           int64       `json:"id"`           // id
	UID          int64       `json:"uid"`          // 微信用户id
	Avatar       string      `json:"avatar"`       // 店员头像
	User         User        `json:"user"`         // 用户信息
	StoreID      int64       `json:"storeId"`      // 门店id
	SystemStore  SystemStore `json:"systemStore"`  // 门店信息
	StaffName    string      `json:"staffName"`    // 店员名称
	Phone        string      `json:"phone"`        // 手机号码
	VerifyStatus int64       `json:"verifyStatus"` // 核销开关
	Status       int64       `json:"status"`       // 状态
	CreateTime   string      `json:"createTime"`   // 创建时间
	UpdateTime   string      `json:"updateTime"`   // 更新时间
}

func (s *SystemStoreStaff) ConvertFromModel(m *model.SystemStoreStaff) error {
	return copier.CopyWithOption(s, m, util.ConvertCopyOption(*s, *m))
}
