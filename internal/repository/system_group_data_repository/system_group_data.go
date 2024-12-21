package system_group_data_repository

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

func (r *repository) GetGroupDataPageList(ctx context.Context, condition validation.SystemGroupDataSearch) ([]*model.SystemGroupData, int64, error) {
	systemGroupData := gen.Q.SystemGroupData
	return gen.SystemGroupData.WithContext(ctx).
		Where(
			systemGroupData.Gid.Eq(condition.GID),
			systemGroupData.Status.Eq(condition.Status),
		).
		Order(systemGroupData.Sort.Asc(), systemGroupData.ID.Asc()).
		FindByPage((condition.Page-1)*condition.Limit, condition.Limit)
}
