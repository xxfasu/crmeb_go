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
	"crmeb_go/pkg/util"
	"github.com/iancoleman/orderedmap"
	"strconv"
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

func (s *service) ChartUser(ctx context.Context) (*orderedmap.OrderedMap, error) {
	everyDateResp, err := s.userService.GetAddUserCountGroupDate(ctx, constants.SearchDateLately30)
	listDate := util.GetListDate(constants.SearchDateLately30)
	resp := orderedmap.New()
	s.setValue(listDate, resp)
	for _, v := range everyDateResp {
		parse, err := time.Parse(time.RFC3339, v.EveryDate)
		if err != nil {
			return nil, err
		}
		formatDate := parse.Format(constants.SystemTimeMonthDayFormat)
		resp.Set(formatDate, v.ID)
	}
	return resp, err
}

func (s *service) ChartOrder(ctx context.Context) (*response.ChartOrder, error) {
	everyDateResp, err := s.storeOrderService.GetOrderGroupByDate(ctx, constants.SearchDateLately30)
	if err != nil {
		return nil, err
	}
	listDate := util.GetListDate(constants.SearchDateLately30)
	resp := new(response.ChartOrder)
	priceMap, idMap, err := s.getPriceAndIdMap(everyDateResp, listDate, constants.SearchDateLately30)
	if err != nil {
		return nil, err
	}

	resp.Price = priceMap
	resp.Quality = idMap
	return resp, err
}

func (s *service) ChartOrderWeek(ctx context.Context) (*response.ChartOrder, error) {
	listDate := util.GetListDate(constants.SearchDateWeek)
	weekData, err := s.storeOrderService.GetOrderGroupByDate(ctx, constants.SearchDateWeek)

	resp := new(response.ChartOrder)
	priceMap, idMap, err := s.getPriceAndIdMap(weekData, listDate, constants.SearchDateWeek)
	if err != nil {
		return nil, err
	}

	// 上周
	preWeekData, err := s.storeOrderService.GetOrderGroupByDate(ctx, constants.SearchDatePreWeek)
	if err != nil {
		return nil, err
	}

	prePriceMap, preIdMap, err := s.getPriceAndIdMap(preWeekData, listDate, constants.SearchDatePreWeek)
	if err != nil {
		return nil, err
	}

	resp.PrePrice = prePriceMap
	resp.PreQuality = preIdMap
	resp.Price = priceMap
	resp.Quality = idMap

	return resp, err
}

func (s *service) ChartOrderMonth(ctx context.Context) (*response.ChartOrder, error) {
	listDate := util.GetListDate(constants.SearchDateMonth)
	monthData, err := s.storeOrderService.GetOrderGroupByDate(ctx, constants.SearchDateMonth)

	resp := new(response.ChartOrder)
	priceMap, idMap, err := s.getPriceAndIdMap(monthData, listDate, constants.SearchDateMonth)
	if err != nil {
		return nil, err
	}

	// 上月
	preWeekData, err := s.storeOrderService.GetOrderGroupByDate(ctx, constants.SearchDatePreMonth)
	if err != nil {
		return nil, err
	}

	prePriceMap, preIdMap, err := s.getPriceAndIdMap(preWeekData, listDate, constants.SearchDateMonth)
	if err != nil {
		return nil, err
	}

	resp.PrePrice = prePriceMap
	resp.PreQuality = preIdMap
	resp.Price = priceMap
	resp.Quality = idMap

	return resp, err
}

func (s *service) ChartOrderYear(ctx context.Context) (*response.ChartOrder, error) {
	listDate := util.GetListDate(constants.SearchDateYear)
	yearData, err := s.storeOrderService.GetOrderGroupByDate(ctx, constants.SearchDateYear)

	resp := new(response.ChartOrder)
	priceMap, idMap, err := s.getPriceAndIdMap(yearData, listDate, constants.SearchDateYear)
	if err != nil {
		return nil, err
	}

	// 上年
	preWeekData, err := s.storeOrderService.GetOrderGroupByDate(ctx, constants.SearchDatePreYear)
	if err != nil {
		return nil, err
	}

	prePriceMap, preIdMap, err := s.getPriceAndIdMap(preWeekData, listDate, constants.SearchDatePreYear)
	if err != nil {
		return nil, err
	}

	resp.PrePrice = prePriceMap
	resp.PreQuality = preIdMap
	resp.Price = priceMap
	resp.Quality = idMap

	return resp, err
}

func (s *service) getPriceAndIdMap(everyDateList []*data.StoreOrderEveryDate, listDate []string, data string) (*orderedmap.OrderedMap, *orderedmap.OrderedMap, error) {
	priceMap := orderedmap.New()
	idMap := orderedmap.New()
	s.setValue(listDate, priceMap)
	s.setValue(listDate, idMap)
	for _, v := range everyDateList {
		switch data {
		case constants.SearchDateLately30:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, err
			}
			formatDate := parse.Format(constants.SystemTimeMonthDayFormat)
			priceMap.Set(formatDate, v.PayPrice)
			idMap.Set(formatDate, v.ID)
		case constants.SearchDatePreWeek, constants.SearchDateWeek:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, err
			}
			// 获取星期几
			weekday := parse.Weekday()
			weekdayList := []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
			weekdayStr := weekdayList[weekday]
			priceMap.Set(weekdayStr, v.PayPrice)
			idMap.Set(weekdayStr, v.ID)
		case constants.SearchDatePreMonth, constants.SearchDateMonth:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, err
			}
			// 获取是当月几号
			day := parse.Day()
			dayStr := strconv.Itoa(day)
			priceMap.Set(dayStr, v.PayPrice)
			idMap.Set(dayStr, v.ID)
		case constants.SearchDatePreYear, constants.SearchDateYear:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, err
			}
			// 获取是当月几号
			month := parse.Month()
			monthList := []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}
			monthStr := monthList[month-1]
			priceMap.Set(monthStr, v.PayPrice)
			idMap.Set(monthStr, v.ID)
		}
	}

	return priceMap, idMap, nil
}

func (s *service) setValue(listDate []string, priceMap *orderedmap.OrderedMap) {
	for _, v := range listDate {
		if _, ok := priceMap.Get(v); ok {
			continue
		}
		priceMap.Set(v, 0)
	}
}
