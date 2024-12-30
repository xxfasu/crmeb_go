package system_admin_repository

import (
	"context"
	"crmeb_go/internal/common/page"
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

func (r *repository) GetUser(ctx context.Context, userName string) (*model.SystemAdmin, error) {
	systemAdmin := gen.Q.SystemAdmin
	return systemAdmin.WithContext(ctx).Where(
		systemAdmin.Account.Eq(userName),
	).First()
}

func (r *repository) GetByID(ctx context.Context, id int64) (*model.SystemAdmin, error) {
	systemAdmin := gen.Q.SystemAdmin
	return systemAdmin.WithContext(ctx).Where(
		systemAdmin.ID.Eq(id),
	).First()
}

func (r *repository) GetPage(ctx context.Context, condition *validation.SystemAdminSearch) ([]*model.SystemAdmin, int64, error) {
	systemAdmin := gen.Q.SystemAdmin
	tx := systemAdmin.WithContext(ctx)
	if len(condition.Roles) > 0 {
		tx = tx.Where(systemAdmin.Roles.Eq(condition.Roles))
	}
	if condition.Status != nil {
		tx = tx.Where(systemAdmin.Status.Eq(*condition.Status))
	}
	if len(condition.RealName) > 0 {
		tx = tx.Where(
			tx.Where(systemAdmin.RealName.Like(condition.RealName)).
				Or(systemAdmin.Account.Like(condition.RealName)))
	}
	return tx.FindByPage(page.PageParam(condition.PageParam))
}

func (r *repository) IsSingleAccount(ctx context.Context, id int64, account string) (bool, error) {
	systemAdmin := gen.Q.SystemAdmin
	count, err := systemAdmin.WithContext(ctx).Where(
		systemAdmin.ID.Neq(id),
		systemAdmin.Account.Eq(account),
	).Count()
	return count == 0, err
}

func (r *repository) IsExistAccount(ctx context.Context, account string) (bool, error) {
	systemAdmin := gen.Q.SystemAdmin
	count, err := systemAdmin.WithContext(ctx).Where(
		systemAdmin.Account.Eq(account),
	).Count()
	return count != 0, err
}

func (r *repository) Create(ctx context.Context, entity *model.SystemAdmin) error {
	systemAdmin := gen.Q.SystemAdmin
	return systemAdmin.WithContext(ctx).Create(entity)
}

func (r *repository) Update(ctx context.Context, entity *model.SystemAdmin) error {
	systemAdmin := gen.Q.SystemAdmin
	_, err := systemAdmin.WithContext(ctx).Updates(entity)
	return err
}

func (r *repository) UpdateFields(ctx context.Context, ID int64, umap map[string]interface{}) error {
	systemAdmin := gen.Q.SystemAdmin
	_, err := systemAdmin.WithContext(ctx).
		Where(
			systemAdmin.ID.Eq(ID),
		).
		Updates(umap)
	return err
}

func (r *repository) DeleteByID(ctx context.Context, ID int64) error {
	systemAdmin := gen.Q.SystemAdmin
	_, err := systemAdmin.WithContext(ctx).
		Where(
			systemAdmin.ID.Eq(ID),
		).
		Delete()
	return err
}
