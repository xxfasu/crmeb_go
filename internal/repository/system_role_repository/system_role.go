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
	return tx.FindByPage((condition.Page-1)*condition.Limit, condition.Limit)
}

func (r *repository) GetByID(ctx context.Context, id int64) (*model.SystemRole, error) {
	systemRole := gen.Q.SystemRole
	return systemRole.WithContext(ctx).Where(systemRole.ID.Eq(id)).First()
}

func (r *repository) ExistRoleName(ctx context.Context, roleName string, id int64) (bool, error) {
	systemRole := gen.Q.SystemRole
	tx := systemRole.WithContext(ctx)
	tx = tx.Where(systemRole.RoleName.Eq(roleName))
	if id > 0 {
		tx = tx.Where(systemRole.ID.Neq(id))
	}
	tx.Limit(1)
	count, err := tx.Count()
	return count > 0, err
}

func (r *repository) UpdateByID(ctx context.Context, umap map[string]any, id int64) error {
	systemRole := gen.Q.SystemRole
	_, err := systemRole.WithContext(ctx).Where(systemRole.ID.Eq(id)).Updates(umap)
	return err
}

func (r *repository) TxCreate(ctx context.Context, tx *gen.Query, entity *model.SystemRole) error {
	return tx.SystemRole.WithContext(ctx).Create(entity)
}

func (r *repository) TxUpdateByID(ctx context.Context, tx *gen.Query, umap map[string]any, id int64) error {
	systemRole := tx.SystemRole
	_, err := systemRole.WithContext(ctx).Where(systemRole.ID.Eq(id)).Updates(umap)
	return err
}

func (r *repository) TxDeleteByID(ctx context.Context, tx *gen.Query, id int64) error {
	systemRole := tx.SystemRole
	_, err := systemRole.WithContext(ctx).Where(systemRole.ID.Eq(id)).Delete()
	return err
}
