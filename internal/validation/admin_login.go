package validation

type SystemAdminLogin struct {
	Account string `form:"account" json:"account" binding:"required,max=32"`
	Pwd     string `form:"pwd" json:"pwd" binding:"required,min=6,max=30"`
	Key     string `form:"key" json:"key" binding:"required"`
	Code    string `form:"code" json:"code" binding:"required"`
}
