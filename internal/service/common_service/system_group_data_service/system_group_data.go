package system_group_data_service

import (
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_group_data_repository"
	"github.com/redis/go-redis/v9"
)

func New(
	redisClient redis.UniversalClient,
	tm repository.Transaction,
	systemGroupDataRepo system_group_data_repository.Repository,
) Service {
	return &service{
		redisClient:         redisClient,
		tm:                  tm,
		systemGroupDataRepo: systemGroupDataRepo,
	}
}

type service struct {
	redisClient         redis.UniversalClient
	tm                  repository.Transaction
	systemGroupDataRepo system_group_data_repository.Repository
}

func (s *service) GetListByGid(gid int) ([]any, error) {
	return nil, nil
}

// func GetListByGid[T any](gid int) ([]T, error) {
// 	return nil, nil
// }
