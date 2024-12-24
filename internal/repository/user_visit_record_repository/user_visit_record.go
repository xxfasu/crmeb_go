package user_visit_record_repository

import (
	"context"
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

func (r *repository) FindPageViewsByDate(ctx context.Context, start int64, end int64) (data int64, err error) {
	userVisitRecord := gen.Q.UserVisitRecord

	return userVisitRecord.WithContext(ctx).Select(userVisitRecord.ID).Where(userVisitRecord.Date.Between(start, end)).Count()
}
