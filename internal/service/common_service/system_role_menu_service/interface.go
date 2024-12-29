package system_role_menu_service

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
)

type Service interface {
	GetMenuIDList(ctx context.Context, roleID int64) ([]int64, error)
	TxBatchCreateSystemRoleMenu(ctx context.Context, tx *gen.Query, menuIDList []*model.SystemRoleMenu) error
	TxDeleteSystemRoleMenuByRoleID(ctx context.Context, tx *gen.Query, roleID int64) error
}
