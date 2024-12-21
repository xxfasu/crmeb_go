package user_repository

import (
	"context"
	"crmeb_go/internal/model"
)

type Reader interface {
	GetUserListInID(ctx context.Context, idList []int64) ([]*model.User, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
