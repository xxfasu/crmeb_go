package user_repository

import (
	"context"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/model"
)

type Reader interface {
	GetUserListInID(ctx context.Context, idList []int64) ([]*model.User, error)
	GetRegisterNumByDate(ctx context.Context, start, end int64) (int64, error)
	GetAddUserCountGroupDate(ctx context.Context, start, end int64) ([]*data.UserEveryDate, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
