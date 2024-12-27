package system_menu_service

import (
	"context"
	"crmeb_go/constants"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_menu_repository"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

func New(
	tm repository.Transaction,
	systemMenuRepo system_menu_repository.Repository,
	redisClient redis.UniversalClient,
) Service {
	return &service{
		tm:             tm,
		systemMenuRepo: systemMenuRepo,
		redisClient:    redisClient,
	}
}

type service struct {
	tm             repository.Transaction
	systemMenuRepo system_menu_repository.Repository
	redisClient    redis.UniversalClient
}

func (s *service) GetAllPermissions(ctx context.Context) ([]*model.SystemMenu, error) {
	return s.systemMenuRepo.GetAllPermissions(ctx)
}

func (s *service) GetUserPermissions(ctx context.Context, userID int64) ([]*model.SystemMenu, error) {
	return s.systemMenuRepo.GetPermissionsByUserID(ctx, userID)
}

func (s *service) GetAllMenus(ctx context.Context) ([]*model.SystemMenu, error) {
	return s.systemMenuRepo.GetAllMenus(ctx)
}

func (s *service) GetUserMenus(ctx context.Context, userID int64) ([]*model.SystemMenu, error) {
	return s.systemMenuRepo.GetMenusByUserID(ctx, userID)
}

func (s *service) GetCacheList(ctx context.Context) ([]*model.SystemMenu, error) {
	result, err := s.redisClient.Exists(ctx, constants.RedisMenuListKey).Result()
	if err != nil {
		return nil, err
	}
	systemMenuList := make([]*model.SystemMenu, 0)
	if result > 0 {
		menuListStr, err := s.redisClient.Get(ctx, constants.RedisMenuListKey).Result()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(menuListStr), &systemMenuList)
		if err != nil {
			return nil, err
		}
		return systemMenuList, nil
	}
	systemMenuList, err = s.systemMenuRepo.GetAllMenus(ctx)
	if err != nil {
		return nil, err
	}
	s.redisClient.Set(ctx, constants.RedisMenuListKey, systemMenuList, 0)
	return systemMenuList, nil
}

func (s *service) GetCacheTree(ctx context.Context) ([]*response.MenuCheck, error) {
	systemMenuList, err := s.GetCacheList(ctx)
	if err != nil {
		return nil, err
	}
	menuCheckList := lo.Map(systemMenuList, func(item *model.SystemMenu, index int) *response.MenuCheck {
		menuCheck := new(response.MenuCheck)
		menuCheck.ConvertFromModel(item)
		return menuCheck
	})
	resp := s.BuildTree(menuCheckList)
	return resp, nil
}
