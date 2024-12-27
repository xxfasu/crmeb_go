package system_menu_service

import (
	"context"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/model"
)

type Service interface {
	GetAllPermissions(ctx context.Context) ([]*model.SystemMenu, error)
	GetUserPermissions(ctx context.Context, userID int64) ([]*model.SystemMenu, error)
	GetAllMenus(ctx context.Context) ([]*model.SystemMenu, error)
	GetUserMenus(ctx context.Context, userID int64) ([]*model.SystemMenu, error)
	GetCacheTree(ctx context.Context) ([]*response.MenuCheck, error)
	BuildTree(menuList []*response.MenuCheck) []*response.MenuCheck
}
