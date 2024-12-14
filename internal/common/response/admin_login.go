package response

import "time"

type ValidateCodeResp struct {
	Key  string `json:"key"`
	Code string `json:"code"`
}

type SystemLoginResp struct {
	ID       int64  `json:"id"`
	Account  string `json:"account"`
	RealName string `json:"real_name"`
	Token    string `json:"token"`
	IsSMS    bool   `json:"is_sms"`
}

type SystemLoginPicResp struct {
	BackgroundImage string `json:"backgroundImage"`
	Logo            string `json:"logo"`
	LoginLogo       string `json:"loginLogo"`
	Banner          string `json:"banner"`
}

type SystemMenusResp struct {
	ID        int               `json:"id"`         // ID
	Pid       int               `json:"pid"`        // 父级ID
	Name      string            `json:"name"`       // 名称
	Icon      string            `json:"icon"`       // 图标
	Perms     string            `json:"perms"`      // 权限标识
	Component string            `json:"component"`  // 组件路径
	MenuType  string            `json:"menu_type"`  // 类型，M-目录，C-菜单，A-按钮
	Sort      int               `json:"sort"`       // 排序
	ChildList []SystemMenusResp `json:"child_list"` // 子对象列表
}

type SystemAdminResp struct {
	ID              int       `json:"id"`
	Account         string    `json:"account"`
	RealName        string    `json:"real_name"`
	Roles           string    `json:"roles"`
	RoleNames       string    `json:"role_names"`
	LastIP          string    `json:"last_ip"`
	LastTime        time.Time `json:"last_time"`
	AddTime         int       `json:"add_time"`
	LoginCount      int       `json:"login_count"`
	Level           int       `json:"level"`
	Status          bool      `json:"status"`
	Token           string    `json:"token"`
	Phone           string    `json:"phone"`
	IsSMS           bool      `json:"is_sms"`
	PermissionsList []string  `json:"permissions_list"`
}
