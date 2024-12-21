package user_service

import (
	"context"
	"crmeb_go/internal/model"
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

func (s *service) GetMapInID(ctx context.Context, uidList []int64) (map[int64]*model.User, error) {
	userMap := make(map[int64]*model.User)
	list, err := s.userRepo.GetUserListInID(ctx, uidList)
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		userMap[item.ID] = item
	}
	return userMap, nil
}
