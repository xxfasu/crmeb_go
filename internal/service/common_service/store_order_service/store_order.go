package store_order_service

import (
	"context"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/store_order_repository"
	"crmeb_go/pkg/util"
)

func New(
	tm repository.Transaction,
	storeOrderRepo store_order_repository.Repository,
) Service {
	return &service{
		tm:             tm,
		storeOrderRepo: storeOrderRepo,
	}
}

type service struct {
	tm             repository.Transaction
	storeOrderRepo store_order_repository.Repository
}

func (s *service) GetOrderNumByDate(ctx context.Context, params *data.DateParams) (*data.DateResult, error) {
	resp := new(data.DateResult)
	// 今日订单数量
	nowPageViews, err := s.storeOrderRepo.FindOrderNumByDate(ctx, params.Start, params.End)
	if err != nil {
		return nil, err
	}

	// 昨天订单数量
	yesterdayPageViews, err := s.storeOrderRepo.FindOrderNumByDate(ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil {
		return nil, err
	}

	resp.NowData = nowPageViews
	resp.YesterdayData = yesterdayPageViews
	return resp, nil
}

func (s *service) GetPayOrderAmountByDate(ctx context.Context, params *data.DateParams) (*data.DateResult, error) {
	resp := new(data.DateResult)
	// 今日销售额
	nowSale, err := s.storeOrderRepo.FindPayOrderAmountByDate(ctx, params.Start, params.End)
	if err != nil {
		return nil, err
	}

	// 昨天销售额
	yesterdaySale, err := s.storeOrderRepo.FindPayOrderAmountByDate(ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil {
		return nil, err
	}

	resp.NowSale = nowSale
	resp.YesterdaySale = yesterdaySale
	return resp, nil
}

func (s *service) GetOrderGroupByDate(ctx context.Context, date string) ([]*data.StoreOrderEveryDate, error) {
	start, end := util.CalculateDateRange(date)
	// 计算时间范围
	list, err := s.storeOrderRepo.FindOrderGroupByDate(ctx, start, end)
	if err != nil {
		return nil, err
	}
	return list, nil
}
