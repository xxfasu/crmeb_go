package system_role_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/validation"
)

type Reader interface {
	GetAllList(ctx context.Context) ([]*model.SystemRole, error)
	GetPage(ctx context.Context, condition *validation.SystemRoleSearch) ([]*model.SystemRole, int64, error)
	GetByID(ctx context.Context, id int64) (*model.SystemRole, error)
	ExistRoleName(ctx context.Context, roleName string, id int64) (bool, error)
}

type Writer interface {
	UpdateByID(ctx context.Context, umap map[string]any, id int64) error
	TxCreate(ctx context.Context, tx *gen.Query, entity *model.SystemRole) error
	TxUpdateByID(ctx context.Context, tx *gen.Query, umap map[string]any, id int64) error
	TxDeleteByID(ctx context.Context, tx *gen.Query, id int64) error
}

type Repository interface {
	Reader
	Writer
}
