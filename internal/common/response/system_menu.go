package response

import (
	"crmeb_go/internal/model"
	"crmeb_go/pkg/util"
	"github.com/jinzhu/copier"
)

type MenuCheck struct {
	ID        int64        `json:"id"`
	Pid       int64        `json:"pid"`
	Name      string       `json:"name"`
	Icon      string       `json:"icon"`
	Checked   int64        `json:"checked"`
	Sort      int64        `json:"sort"`
	ChildList []*MenuCheck `json:"childList"`
}

func (c *MenuCheck) ConvertFromModel(m *model.SystemMenu) error {
	return copier.CopyWithOption(c, m, util.ConvertCopyOption(*c, *m))
}
