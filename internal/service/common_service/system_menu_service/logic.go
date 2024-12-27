package system_menu_service

import (
	"context"
	"crmeb_go/constants"
	"crmeb_go/internal/model"
	"encoding/json"
)

func (s *service) getCacheList(ctx context.Context) ([]*model.SystemMenu, error) {
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
