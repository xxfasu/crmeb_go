package system_store_service

import (
	"context"
	"crmeb_go/internal/common/response"
)

type Service interface {
	GetMapInID(ctx context.Context, storeIDList []int64) (map[int64]response.SystemStore, error)
}
