package system_store_service

import (
	"context"
	"crmeb_go/internal/common/response"
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

func (s *service) GetMapInID(ctx context.Context, storeIDList []int64) (map[int64]response.SystemStore, error) {
	systemStoreMap := make(map[int64]response.SystemStore)
	if len(storeIDList) == 0 {
		return systemStoreMap, nil
	}
	list, err := s.systemStoreRepo.GetStoreList(ctx, storeIDList)
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		var systemStore response.SystemStore
		err = systemStore.ConvertFromModel(item)
		if err != nil {
			return nil, err
		}
		systemStoreMap[item.ID] = systemStore
	}
	return systemStoreMap, nil
}
