package validation

type SystemRoleSearch struct {
	RoleName string `form:"roleName"`
	Status   *int64 `form:"status"`
	PageParam
}

type SystemRole struct {
	ID       int64  `json:"id"`
	RoleName string `json:"roleName" binding:"max=32"`
	Rules    string `json:"rules" binding:"required"`
	Status   int64  `json:"status" binding:"required"`
}
