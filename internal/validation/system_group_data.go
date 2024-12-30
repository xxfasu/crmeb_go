package validation

type SystemGroupDataSearch struct {
	Keywords string `json:"keywords" form:"keywords"`
	GID      int64  `json:"gid" form:"gid"`
	Status   int64  `json:"status" form:"status"`
	PageParam
}

// SystemFormItemCheck 表单字段明细
type SystemFormItemCheck struct {
	Name  string `json:"name" binding:"required" `
	Value string `json:"value"`
	Title string `json:"title"`
}
