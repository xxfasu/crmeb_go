package home_service

import (
	"context"
	"crmeb_go/internal/common/response"
)

type Service interface {
	// IndexDate 首页数据
	IndexDate(ctx context.Context) (*response.HomeRate, error)

	ChartUser(ctx context.Context) (*map[string]interface{}, error)

	ChartOrder(ctx context.Context) (*map[string]interface{}, error)
}
