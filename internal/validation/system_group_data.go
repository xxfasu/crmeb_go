package validation

type SystemGroupDataSearchReq struct {
	Keywords string `json:"keywords"`
	GID      int64  `json:"gid"`
	Status   int64  `json:"status"`
	PageParam
}

// SystemFormItemCheckReq 表单字段明细
type SystemFormItemCheckReq struct {
	Name  string `json:"name" binding:"required" `
	Value string `json:"value"`
	Title string `json:"title"`
}
