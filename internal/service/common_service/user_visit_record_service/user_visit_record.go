package user_visit_record_service

import (
	"context"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/user_visit_record_repository"
	"crmeb_go/internal/service/common_service/user_service"
)

func New(
	tm repository.Transaction,
	userService user_service.Service,
	userVisitRecordRepo user_visit_record_repository.Repository,
) Service {
	return &service{
		tm:                  tm,
		userService:         userService,
		userVisitRecordRepo: userVisitRecordRepo,
	}
}

type service struct {
	tm                  repository.Transaction
	userService         user_service.Service
	userVisitRecordRepo user_visit_record_repository.Repository
}

// FindPageViewsByDate 首页数据
func (s *service) FindPageViewsByDate(ctx context.Context, params *data.DateParams) (*data.DateResult, error) {
	resp := new(data.DateResult)
	// 今日用户访问量
	nowPageViews, err := s.userVisitRecordRepo.FindPageViewsByDate(ctx, params.Start, params.End)
	if err != nil {
		return nil, err
	}

	// 昨天用户访问量
	yesterdayPageViews, err := s.userVisitRecordRepo.FindPageViewsByDate(ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil {
		return nil, err
	}

	resp.NowData = nowPageViews
	resp.YesterdayData = yesterdayPageViews

	return resp, nil
}
