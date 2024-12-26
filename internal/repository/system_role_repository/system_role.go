package system_role_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/validation"
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

func (r *repository) GetList(ctx context.Context, condition *validation.SystemRoleSearch) ([]*model.SystemRole, int64, error) {
	systemRole := gen.Q.SystemRole
	tx := systemRole.WithContext(ctx)
	tx = tx.Select(systemRole.ID, systemRole.RoleName, systemRole.Status, systemRole.CreatedAt, systemRole.UpdatedAt)
	if len(condition.RoleName) != 0 {
		tx = tx.Where(systemRole.RoleName.Like(condition.RoleName))
	}
	if condition.Status != nil {
		tx = tx.Where(systemRole.Status.Eq(*condition.Status))
	}
	tx = tx.Order(systemRole.ID.Asc())
	return tx.FindByPage(condition.Page, condition.Limit)
}
