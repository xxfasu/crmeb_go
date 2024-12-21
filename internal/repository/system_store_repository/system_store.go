package system_store_repository

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

func (r *repository) GetStoreList(ctx context.Context, storeIDList []int64) ([]*model.SystemStore, error) {
	systemStore := gen.Q.SystemStore

	return gen.SystemStore.WithContext(ctx).
		Where(systemStore.ID.In(storeIDList...)).
		Find()
}
