package system_store_staff_repository

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

func (r *repository) GetStoreStaffPageList(ctx context.Context, condition *validation.GetSystemStoreStaffList) ([]*model.SystemStoreStaff, int64, error) {
	systemStoreStaff := gen.Q.SystemStoreStaff
	tx := gen.SystemStoreStaff.WithContext(ctx)
	if condition.StoreID > 0 {
		tx = tx.Where(
			systemStoreStaff.StoreID.Eq(condition.StoreID),
		)
	}
	return tx.FindByPage((condition.Page-1)*condition.Limit, condition.Limit)
}
