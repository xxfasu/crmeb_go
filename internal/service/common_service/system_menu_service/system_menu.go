package system_menu_service

import (
	"context"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_menu_repository"
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

func (s *service) GetCacheTree(ctx context.Context) ([]*response.MenuCheck, error) {
	systemMenuList, err := s.getCacheList(ctx)
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
