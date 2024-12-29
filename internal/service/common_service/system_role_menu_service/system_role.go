package system_role_menu_service

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/repository/system_role_menu_repository"
	"github.com/samber/lo"
)

func New(
	tm repository.Transaction,
	systemRoleMenuRepo system_role_menu_repository.Repository,
) Service {
	return &service{
		tm:                 tm,
		systemRoleMenuRepo: systemRoleMenuRepo,
	}
}

type service struct {
	tm                 repository.Transaction
	systemRoleMenuRepo system_role_menu_repository.Repository
}

func (s *service) GetMenuIDList(ctx context.Context, roleID int64) ([]int64, error) {
	menuList, err := s.systemRoleMenuRepo.GetMenuIDListByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}
	menuIDList := lo.Map(menuList, func(item *model.SystemRoleMenu, index int) int64 {
		return item.MenuID
	})
	return menuIDList, nil
}

func (s *service) TxBatchCreateSystemRoleMenu(ctx context.Context, tx *gen.Query, menuIDList []*model.SystemRoleMenu) error {
	return s.systemRoleMenuRepo.TxBatchCreate(ctx, tx, menuIDList)
}

func (s *service) TxDeleteSystemRoleMenuByRoleID(ctx context.Context, tx *gen.Query, roleID int64) error {
	return s.systemRoleMenuRepo.TxDeleteByRoleID(ctx, tx, roleID)
}
