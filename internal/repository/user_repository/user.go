package user_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
	"errors"
	"gorm.io/gorm"
)

func NewUserRepository(
	db *gorm.DB,
) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	err := gen.User.WithContext(ctx).Create(user)
	return err
}

func (r *userRepository) CreateTx(ctx context.Context, query *gen.Query, user *model.User) error {
	err := query.User.WithContext(ctx).Create(user)
	return err
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	err := gen.User.WithContext(ctx).Save(user)
	return err
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*model.User, error) {
	userList, err := gen.User.WithContext(ctx).Where(gen.User.Account.Eq(userId)).Find()
	return userList[0], err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	userList, err := gen.User.WithContext(ctx).Where(gen.User.Account.Eq(email)).Find()
	return userList[0], err
}

func (r *userRepository) GetUserByCondition(ctx context.Context, condition user_data.Condition) (*model.User, error) {
	user, err := gen.User.WithContext(ctx).GetUserByCondition(condition)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, err
}
