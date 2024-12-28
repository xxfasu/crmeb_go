package system_role_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/validation"
)

type Reader interface {
	GetList(ctx context.Context, condition *validation.SystemRoleSearch) ([]*model.SystemRole, int64, error)
	GetByID(ctx context.Context, id int64) (*model.SystemRole, error)
	ExistRoleName(ctx context.Context, roleName string, id int64) (bool, error)
}

type Writer interface {
	TxCreate(ctx context.Context, tx *gen.Query, entity *model.SystemRole) error
}

type Repository interface {
	Reader
	Writer
}
