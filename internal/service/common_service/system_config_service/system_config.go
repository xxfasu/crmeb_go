package system_config_service

import (
	"context"
	"crmeb_go/constants"
	"crmeb_go/internal/conf"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_config_repository"
	"crmeb_go/pkg/logs"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func New(
	redisClient redis.UniversalClient,
	tm repository.Transaction,
	systemConfigRepo system_config_repository.Repository,
) Service {
	return &service{
		redisClient:      redisClient,
		tm:               tm,
		systemConfigRepo: systemConfigRepo,
	}
}

type service struct {
	redisClient      redis.UniversalClient
	tm               repository.Transaction
	systemConfigRepo system_config_repository.Repository
}

// GetValueByKey 根据menu name 获取 value
func (s *service) GetValueByKey(ctx context.Context, key string) string {
	return s.get(ctx, key)
}

// BatchGetValueByKey 根据menu name 批量获取 value
func (s *service) BatchGetValueByKey(ctx context.Context, keyList []string) []string {
	valueList := make([]string, 0, len(keyList))
	for _, key := range keyList {
		valueList = append(valueList, s.GetValueByKey(ctx, key))
	}
	return valueList
}

func (s *service) get(ctx context.Context, name string) string {
	if !conf.Config.CrmebConfig.AsyncConfig {
		systemConfig, err := s.systemConfigRepo.GetConfigByName(ctx, name)
		if err != nil {
			logs.Log.Error("GetConfigByName err: ", zap.Error(err))
			return ""
		}
		return systemConfig.Value
	}

	s.setRedisByVoList(ctx)

	result, _ := s.redisClient.HGet(ctx, constants.RedisConfigListKey, name).Result()
	if lo.IsEmpty(result) {
		return ""
	}
	return result
}

func (s *service) setRedisByVoList(ctx context.Context) {
	size, _ := s.redisClient.HLen(ctx, constants.RedisConfigListKey).Result()
	if size > 0 || !conf.Config.CrmebConfig.AsyncConfig {
		return
	}
	systemConfigList, err := s.systemConfigRepo.GetConfigALL(ctx)
	if err != nil {
		logs.Log.Error("GetConfigALL err: ", zap.Error(err))
		return
	}
	s.async(ctx, systemConfigList)
}

func (s *service) async(ctx context.Context, systemConfigList []*model.SystemConfig) {
	if !conf.Config.CrmebConfig.AsyncConfig {
		return
	}
	for _, config := range systemConfigList {
		s.redisClient.HSet(ctx, constants.RedisConfigListKey, config.Name, config.Value)
	}
}
