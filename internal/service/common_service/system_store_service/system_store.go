package system_store_service

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_store_repository"
)

func New(
	tm repository.Transaction,
	systemStoreRepo system_store_repository.Repository,
) Service {
	return &service{
		tm:              tm,
		systemStoreRepo: systemStoreRepo,
	}
}

type service struct {
	tm              repository.Transaction
	systemStoreRepo system_store_repository.Repository
}

func (s *service) GetMapInID(ctx context.Context, storeIDList []int64) (map[int64]*model.SystemStore, error) {
	systemStoreMap := make(map[int64]*model.SystemStore)
	if len(storeIDList) == 0 {
		return systemStoreMap, nil
	}
	list, err := s.systemStoreRepo.GetStoreList(ctx, storeIDList)
	if err != nil {
		return nil, err
	}
	for _, systemStore := range list {
		systemStoreMap[systemStore.ID] = systemStore
	}
	return systemStoreMap, nil
}
