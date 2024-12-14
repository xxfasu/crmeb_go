package system_menu_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

func New(
	db *gorm.DB,
) Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetAllPermissions(ctx context.Context) ([]*model.SystemMenu, error) {
	systemMenu := gen.Q.SystemMenu
	return gen.SystemMenu.WithContext(ctx).Where(
		systemMenu.MenuType.Neq("M"),
	).Find()
}

func (r *repository) GetPermissionsByUserID(ctx context.Context, userID int64) ([]*model.SystemMenu, error) {
	return gen.SystemMenu.WithContext(ctx).GetUserPermission(userID)
}

func (r *repository) GetAllMenus(ctx context.Context) ([]*model.SystemMenu, error) {
	systemMenu := gen.Q.SystemMenu
	return gen.SystemMenu.WithContext(ctx).Where(
		systemMenu.IsShow.Is(true),
		systemMenu.MenuType.Eq("A"),
	).Find()
}

func (r *repository) GetMenusByUserID(ctx context.Context, userID int64) ([]*model.SystemMenu, error) {
	return gen.SystemMenu.WithContext(ctx).GetUserMenus(userID)
}
