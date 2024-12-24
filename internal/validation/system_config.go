package validation

// GetUniq 根据key获取表单配置数据
type GetUniq struct {
	Key string `form:"key" binding:"required"`
}
