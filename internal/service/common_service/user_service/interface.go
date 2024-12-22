package user_service

import (
	"context"
	"crmeb_go/internal/common/response"
)

type Service interface {
	GetMapInID(ctx context.Context, uidList []int64) (map[int64]response.User, error)
}
