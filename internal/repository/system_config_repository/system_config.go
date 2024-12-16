package system_config_repository

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

func (r *repository) GetConfigByName(ctx context.Context, name string) (*model.SystemConfig, error) {
	systemConfig := gen.Q.SystemConfig
	return gen.SystemConfig.WithContext(ctx).Where(
		systemConfig.Status.Eq(0),
		systemConfig.Name.Eq(name),
	).First()
}

func (r *repository) GetConfigALL(ctx context.Context) ([]*model.SystemConfig, error) {
	systemConfig := gen.Q.SystemConfig
	return gen.SystemConfig.WithContext(ctx).Where(
		systemConfig.Status.Eq(0),
	).Find()
}
