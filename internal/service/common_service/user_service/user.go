package user_service

import (
	"context"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/user_repository"
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
