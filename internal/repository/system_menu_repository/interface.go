package system_menu_repository

import (
	"context"
	"crmeb_go/internal/model"
)

type Reader interface {
	GetAllPermissions(ctx context.Context) ([]*model.SystemMenu, error)
	GetUserPermission(ctx context.Context, userID int64) ([]*model.SystemMenu, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
