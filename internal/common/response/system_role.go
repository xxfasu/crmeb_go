package response

import (
	"crmeb_go/internal/model"
	"crmeb_go/pkg/util"
	"github.com/jinzhu/copier"
)

type SystemRole struct {
	ID         int64  `json:"id"`
	RoleName   string `json:"roleName"`
	Rules      string `json:"rules"`
	Status     int64  `json:"status"`
	Level      int64  `json:"level"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func (s *SystemRole) ConvertFromModel(m *model.SystemRole) error {
	return copier.CopyWithOption(s, m, util.ConvertCopyOption(*s, *m))
}
