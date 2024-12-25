package home_service

import (
	"context"
	"crmeb_go/internal/common/response"
	"github.com/iancoleman/orderedmap"
)

type Service interface {
	// IndexDate 首页数据
	IndexDate(ctx context.Context) (*response.HomeRate, error)

	ChartUser(ctx context.Context) (*orderedmap.OrderedMap, error)

	ChartOrder(ctx context.Context) (*response.ChartOrder, error)

	ChartOrderWeek(ctx context.Context) (*response.ChartOrder, error)

	ChartOrderMonth(ctx context.Context) (*response.ChartOrder, error)

	ChartOrderYear(ctx context.Context) (*response.ChartOrder, error)
}
