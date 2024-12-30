package system_store_staff_repository

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

func (r *repository) GetStoreStaffPageList(ctx context.Context, condition *validation.GetSystemStoreStaffList) ([]*model.SystemStoreStaff, int64, error) {
	systemStoreStaff := gen.Q.SystemStoreStaff
	tx := systemStoreStaff.WithContext(ctx)
	if condition.StoreID > 0 {
		tx = tx.Where(
			systemStoreStaff.StoreID.Eq(condition.StoreID),
		)
	}
	return tx.FindByPage(page.PageParam(condition.PageParam))
}
