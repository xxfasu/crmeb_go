package user_repository

import (
	"context"
	"crmeb_go/internal/common/data"
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

func (r *repository) GetUserListInID(ctx context.Context, idList []int64) ([]*model.User, error) {
	user := gen.Q.User

	return user.WithContext(ctx).
		Where(user.ID.In(idList...)).
		Find()
}

func (r *repository) GetRegisterNumByDate(ctx context.Context, start, end int64) (int64, error) {
	user := gen.Q.User

	return user.WithContext(ctx).
		Where(user.CreatedAt.Between(start, end)).
		Count()
}

func (r *repository) GetAddUserCountGroupDate(ctx context.Context, start, end int64) ([]*data.UserEveryDate, error) {
	user := gen.Q.User
	return user.WithContext(ctx).GetAddUserCountGroupDate(&data.DateCondition{Start: start, End: end})
}
