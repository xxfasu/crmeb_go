package system_admin_repository

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

func (r *repository) GetUser(ctx context.Context, userName string) (*model.SystemAdmin, error) {
	systemAdmin := gen.Q.SystemAdmin
	return gen.SystemAdmin.WithContext(ctx).Where(
		systemAdmin.Account.Eq(userName),
	).First()
}

func (r *repository) Update(ctx context.Context, ID int64, umap map[string]interface{}) error {
	systemAdmin := gen.Q.SystemAdmin
	_, err := gen.SystemAdmin.WithContext(ctx).
		Where(
			systemAdmin.ID.Eq(ID),
		).
		Updates(umap)
	return err
}
