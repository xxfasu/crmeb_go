package user_service

import (
	"context"
	"crmeb_go/internal/model"
)

type Service interface {
	GetMapInID(ctx context.Context, uidList []int64) (map[int64]*model.User, error)
}
