package validation

type SystemRoleSearch struct {
	RoleName string `form:"roleName"`
	Status   *int64 `form:"status"`
	PageParam
}

type SystemRole struct {
	ID       int64  `json:"id"`
	RoleName string `json:"roleName"`
	Rules    string `json:"rules"`
	Status   int64  `json:"status"`
}
