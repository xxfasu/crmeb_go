package system_role_menu_service

import (
	"context"
)

type Service interface {
	GetMenuIDList(ctx context.Context, roleID int64) ([]int64, error)
}
