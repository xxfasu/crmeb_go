package user_service

import (
	"context"
	"crmeb_go/constants"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/user_repository"
	"crmeb_go/pkg/util"
	"time"
)

func New(
	tm repository.Transaction,
	userRepo user_repository.Repository,
) Service {
	return &service{
		tm:       tm,
		userRepo: userRepo,
	}
}

type service struct {
	tm       repository.Transaction
	userRepo user_repository.Repository
}

func (s *service) GetMapInID(ctx context.Context, uidList []int64) (map[int64]response.User, error) {
	userMap := make(map[int64]response.User)
	list, err := s.userRepo.GetUserListInID(ctx, uidList)
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		var user response.User
		err = user.ConvertFromModel(item)
		if err != nil {
			return nil, err
		}
		userMap[item.ID] = user
	}
	return userMap, nil
}

func (s *service) GetRegisterNumByDate(ctx context.Context, params *data.DateParams) (*data.DateResult, error) {
	resp := new(data.DateResult)
	// 今日用户注册量
	nowPageViews, err := s.userRepo.GetRegisterNumByDate(ctx, params.Start, params.End)
	if err != nil {
		return resp, err
	}

	// 昨天用户注册量
	yesterdayPageViews, err := s.userRepo.GetRegisterNumByDate(ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil {
		return nil, err
	}

	resp.NowData = nowPageViews
	resp.YesterdayData = yesterdayPageViews
	return resp, nil
}

func (s *service) GetAddUserCountGroupDate(ctx context.Context, date string) (*map[string]interface{}, error) {
	start, end := util.CalculateDateRange(date)
	// 计算时间范围
	data, err := s.userRepo.GetAddUserCountGroupDate(ctx, start, end)
	if err != nil {
		return nil, err
	}

	resp := make(map[string]interface{}, len(data))
	for _, v := range data {
		parse, err := time.Parse(time.RFC3339, v.EveryDate)
		if err != nil {
			return nil, err
		}
		formatDate := parse.Format(constants.SystemTimeMonthDayFormat)
		resp[formatDate] = v.ID
	}

	return &resp, nil
}
