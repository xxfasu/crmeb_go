package validation

type SystemAdminSearch struct {
	RealName string `form:"realName"` // 后台管理员姓名
	Roles    string `form:"roles"`    // 后台管理员角色(menus_id)
	Status   *int64 `form:"status"`   // 后台管理员状态 1有效0无效
	PageParam
}

type SystemAdminAdd struct {
	Account  string `json:"account" binding:"required,max=32"`     // 后台管理员账号
	Pwd      string `json:"pwd" binding:"required,max=32"`         // 后台管理员密码
	RealName string `json:"real_name" binding:"required,max=16"`   // 后台管理员姓名
	Roles    string `json:"roles" binding:"required,max=128"`      // 后台管理员角色(menus_id)
	Status   int64  `json:"status" binding:"required,min=0,max=1"` // 后台管理员状态 1有效0无效
	Phone    string `json:"phone" binding:"required,phone"`        // 手机号
}

type SystemAdminUpdate struct {
	ID       int64  `json:"id" binding:"required"`
	Account  string `json:"account" binding:"required,max=32"`     // 后台管理员账号
	Pwd      string `json:"pwd" binding:"max=32"`                  // 后台管理员密码
	RealName string `json:"real_name" binding:"required,max=16"`   // 后台管理员姓名
	Roles    string `json:"roles" binding:"required,max=128"`      // 后台管理员角色(menus_id)
	Status   int64  `json:"status" binding:"required,min=0,max=1"` // 后台管理员状态 1有效0无效
	Phone    string `json:"phone" binding:"required,phone"`        // 手机号
}
