package system_menu_repository

import (
	"gorm.io/gen"
)

type Querier interface {
	//		SELECT DISTINCT m.*
	//		FROM eb_system_menu m
	//		RIGHT JOIN eb_system_role_menu rm on rm.menu_id = m.id
	//		RIGHT JOIN eb_system_role r on rm.rid = r.id
	//		RIGHT JOIN eb_system_admin a on FIND_IN_SET(r.id, a.roles)
	// 		{{where}}
	// 			m.deleted_at = 0 AND r.status = 1 AND a.id = @userID
	// 		{{end}}
	GetUserPermission(userID int64) ([]*gen.T, error)
}
