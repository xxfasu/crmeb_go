package system_role_menu_repository

import (
	"context"
	"crmeb_go/internal/model"
)

type Reader interface {
	GetMenuIDListByRoleID(ctx context.Context, roleID int64) ([]*model.SystemRoleMenu, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
