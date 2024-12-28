package system_role_menu_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
)

type Reader interface {
	GetMenuIDListByRoleID(ctx context.Context, roleID int64) ([]*model.SystemRoleMenu, error)
}

type Writer interface {
	TxBatchCreate(ctx context.Context, tx *gen.Query, entityList []*model.SystemRoleMenu) error
}

type Repository interface {
	Reader
	Writer
}
