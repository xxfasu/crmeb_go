package user_visit_record_service

import (
	"context"
	"crmeb_go/internal/common/data"
)

type Service interface {
	// FindPageViewsByDate 首页数据
	FindPageViewsByDate(ctx context.Context, params *data.DateParams) (*data.DateResult, error)
}
