package system_store_service

import (
	"context"
	"crmeb_go/internal/model"
)

type Service interface {
	GetMapInID(ctx context.Context, storeIDList []int64) (map[int64]*model.SystemStore, error)
}
