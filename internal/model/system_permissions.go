package model

type SystemPermissions struct {
	ID       int64  `json:"id" `        // ID
	Pid      int64  `json:"pid" `       // 父级ID
	Name     string `json:"name" `      // 权限名称
	Path     string `json:"path" `      // 权限地址
	Sort     int64  `json:"sort" `      // 排序
	IsDelete bool   `json:"is_delete" ` // 是否删除"
}
