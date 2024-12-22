package response

import (
	"crmeb_go/internal/model"
	"crmeb_go/pkg/util"
	"github.com/jinzhu/copier"
)

type SystemStore struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`            // 门店名称
	Introduction    string `json:"introduction"`    // 简介
	Phone           string `json:"phone"`           // 手机号码
	Address         string `json:"address"`         // 省市区
	DetailedAddress string `json:"detailedAddress"` // 详细地址
	Image           string `json:"image"`           // 门店logo
	Latitude        string `json:"latitude"`        // 纬度
	Longitude       string `json:"longitude"`       // 经度
	ValidTime       string `json:"validTime"`       // 核销有效日期
	DayTime         string `json:"dayTime"`         // 每日营业开关时间
	IsShow          bool   `json:"isShow"`          // 是否显示
	CreateTime      string `json:"createTime"`      // 创建时间
	UpdateTime      string `json:"updateTime"`      // 修改时间
}

func (s *SystemStore) ConvertFromModel(m *model.SystemStore) error {
	return copier.CopyWithOption(s, m, util.ConvertCopyOption(*s, *m))
}
