package validation

type SystemGroupDataSearch struct {
	Keywords string `json:"keywords"`
	GID      int64  `json:"gid"`
	Status   int64  `json:"status"`
	PageParam
}

// SystemFormItemCheck 表单字段明细
type SystemFormItemCheck struct {
	Name  string `json:"name" binding:"required" `
	Value string `json:"value"`
	Title string `json:"title"`
}
