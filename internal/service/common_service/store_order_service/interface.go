package store_order_service

import (
	"context"
	"crmeb_go/internal/common/data"
)

type Service interface {
	GetOrderNumByDate(ctx context.Context, params *data.DateParams) (*data.DateResult, error)
	GetPayOrderAmountByDate(ctx context.Context, params *data.DateParams) (*data.DateResult, error)
	GetOrderGroupByDate(ctx context.Context, date string) (*map[string]interface{}, error)
}
