package system_menu_repository

import (
	"context"
	"crmeb_go/internal/model"
)

type Reader interface {
	GetAllPermissions(ctx context.Context) ([]*model.SystemMenu, error)
	GetPermissionsByUserID(ctx context.Context, userID int64) ([]*model.SystemMenu, error)
	GetMenusByIDList(ctx context.Context, menuIDList []int64) ([]*model.SystemMenu, error)
	GetAllMenus(ctx context.Context) ([]*model.SystemMenu, error)
	GetMenusByUserID(ctx context.Context, userID int64) ([]*model.SystemMenu, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
