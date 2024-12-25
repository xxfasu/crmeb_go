package user_service

import (
	"context"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/common/response"
)

type Service interface {
	GetMapInID(ctx context.Context, uidList []int64) (map[int64]response.User, error)
	GetRegisterNumByDate(ctx context.Context, params *data.DateParams) (*data.DateResult, error)
	GetAddUserCountGroupDate(ctx context.Context, date string) ([]*data.UserEveryDate, error)
}
