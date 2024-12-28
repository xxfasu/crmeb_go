package system_role_menu_repository

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

func (r *repository) GetMenuIDListByRoleID(ctx context.Context, roleID int64) ([]*model.SystemRoleMenu, error) {
	systemRoleMenu := gen.Q.SystemRoleMenu
	return systemRoleMenu.WithContext(ctx).
		Select(systemRoleMenu.MenuID).
		Where(systemRoleMenu.Rid.Eq(roleID)).
		Find()
}

func (r *repository) TxBatchCreate(ctx context.Context, tx *gen.Query, entityList []*model.SystemRoleMenu) error {
	return tx.WithContext(ctx).SystemRoleMenu.Create(entityList...)
}
