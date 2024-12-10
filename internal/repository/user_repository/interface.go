package user_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
)

type Reader interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByCondition(ctx context.Context, condition user_data.Condition) (*model.User, error)
}

type Writer interface {
	Create(ctx context.Context, user *model.User) error
	CreateTx(ctx context.Context, query *gen.Query, user *model.User) error
	Update(ctx context.Context, user *model.User) error
}

type UserRepository interface {
	Reader
	Writer
}
