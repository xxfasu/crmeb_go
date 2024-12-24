package home_service

import (
	"context"
	"crmeb_go/constants"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/service/common_service/store_order_service"
	"crmeb_go/internal/service/common_service/user_service"
	"crmeb_go/internal/service/common_service/user_visit_record_service"
	"time"
)

func New(
	tm repository.Transaction,
	storeOrderService store_order_service.Service,
	userService user_service.Service,
	userVisitRecordService user_visit_record_service.Service,
) Service {
	return &service{
		tm:                     tm,
		storeOrderService:      storeOrderService,
		userService:            userService,
		userVisitRecordService: userVisitRecordService,
	}
}

type service struct {
	tm                     repository.Transaction
	storeOrderService      store_order_service.Service
	userService            user_service.Service
	userVisitRecordService user_visit_record_service.Service
}

// IndexDate 首页数据
func (s *service) IndexDate(ctx context.Context) (*response.HomeRate, error) {
	resp := new(response.HomeRate)
	now := time.Now()
	// 今天0点时间
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayUnix := today.Unix()
	// 计算昨天0点的时间
	yesterdayZero := today.Add(-24 * time.Hour).Unix()
	// 计算昨天24点的时间
	yesterdayEnd := today.Add(-1 * time.Second).Unix()
	// 现在的时间
	nowUnix := time.Now().Unix() - 1
	dateParams := &data.DateParams{
		Start:          todayUnix,
		End:            nowUnix,
		YesterdayStart: yesterdayZero,
		YesterdayEnd:   yesterdayEnd,
	}
	// 取今天，昨日访问量
	pageViews, err := s.userVisitRecordService.FindPageViewsByDate(ctx, dateParams)
	if err != nil {
		return nil, err
	}
	// 今日昨日注册用户数量
	registerNum, err := s.userService.GetRegisterNumByDate(ctx, dateParams)
	if err != nil {
		return nil, err
	}
	// 今日昨日注册用户数量
	orderNum, err := s.storeOrderService.GetOrderNumByDate(ctx, dateParams)
	if err != nil {
		return nil, err
	}
	// 今日昨日销售额
	payOrderAmount, err := s.storeOrderService.GetPayOrderAmountByDate(ctx, dateParams)
	if err != nil {
		return nil, err
	}
	resp.PageViews = pageViews.NowData
	resp.YesterdayPageViews = pageViews.YesterdayData
	resp.NewUserNum = registerNum.NowData
	resp.YesterdayNewUserNum = registerNum.YesterdayData
	resp.OrderNum = orderNum.NowData
	resp.YesterdayOrderNum = orderNum.YesterdayData
	resp.Sale = payOrderAmount.NowSale
	resp.YesterdaySale = payOrderAmount.YesterdaySale
	return resp, nil
}

func (s *service) ChartUser(ctx context.Context) (*map[string]interface{}, error) {
	resp, err := s.userService.GetAddUserCountGroupDate(ctx, constants.SearchDateLately30)
	return resp, err
}

func (s *service) ChartOrder(ctx context.Context) (*map[string]interface{}, error) {
	resp, err := s.storeOrderService.GetOrderGroupByDate(ctx, constants.SearchDateLately30)
	return resp, err
}
