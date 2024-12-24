package system_menu_service

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_menu_repository"
)

func New(
	tm repository.Transaction,
	systemMenuRepo system_menu_repository.Repository,
) Service {
	return &service{
		tm:             tm,
		systemMenuRepo: systemMenuRepo,
	}
}

type service struct {
	tm             repository.Transaction
	systemMenuRepo system_menu_repository.Repository
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
