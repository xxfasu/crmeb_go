package user_repository

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

func (r *repository) GetUserListInID(ctx context.Context, idList []int64) ([]*model.User, error) {
	user := gen.Q.User

	return gen.User.WithContext(ctx).
		Where(user.ID.In(idList...)).
		Find()
}
