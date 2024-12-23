package response

import (
	"crmeb_go/internal/model"
	"crmeb_go/pkg/util"
	"github.com/jinzhu/copier"
)

type ValidateCode struct {
	Key  string `json:"key"`
	Code string `json:"code"`
}

type SystemLogin struct {
	ID       int64  `json:"id"`
	Account  string `json:"account"`
	RealName string `json:"realName"`
	Token    string `json:"token"`
	IsSMS    bool   `json:"isSms"`
}

type SystemLoginPic struct {
	BackgroundImage string                            `json:"backgroundImage"`
	Logo            string                            `json:"logo"`
	LoginLogo       string                            `json:"loginLogo"`
	Banner          []SystemGroupDataAdminLoginBanner `json:"banner"`
}

type SystemMenu struct {
	ID        int64         `json:"id"`        // ID
	Pid       int64         `json:"pid"`       // 父级ID
	Name      string        `json:"name"`      // 名称
	Icon      string        `json:"icon"`      // 图标
	Perms     string        `json:"perms"`     // 权限标识
	Component string        `json:"component"` // 组件路径
	MenuType  string        `json:"menuType"`  // 类型，M-目录，C-菜单，A-按钮
	Sort      int64         `json:"sort"`      // 排序
	ChildList []*SystemMenu `json:"childList"` // 子对象列表
}

func (s *SystemMenu) ConvertFromModel(m *model.SystemMenu) error {
	return copier.CopyWithOption(s, m, util.ConvertCopyOption(*s, *m))
}

type SystemAdmin struct {
	ID              int64    `json:"id"`
	Account         string   `json:"account"`
	RealName        string   `json:"realName"`
	Roles           string   `json:"roles"`
	RoleNames       string   `json:"roleNames"`
	LastIP          string   `json:"lastIp"`
	LastTime        string   `json:"lastTime"`
	AddTime         string   `json:"addTime"`
	LoginCount      int64    `json:"loginCount"`
	Level           int64    `json:"level"`
	Status          bool     `json:"status"`
	Token           string   `json:"token"`
	Phone           string   `json:"phone"`
	IsSMS           bool     `json:"isSms"`
	PermissionsList []string `json:"permissionsList"`
}

func (s *SystemAdmin) ConvertFromModel(m *model.SystemAdmin) error {
	return copier.CopyWithOption(s, m, util.ConvertCopyOption(*s, *m))
}

type SystemGroupDataAdminLoginBanner struct {
	PIC string `json:"pic"`
}
