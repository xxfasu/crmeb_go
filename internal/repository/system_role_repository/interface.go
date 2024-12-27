package system_role_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/validation"
)

type Reader interface {
	GetList(ctx context.Context, condition *validation.SystemRoleSearch) ([]*model.SystemRole, int64, error)
	GetByID(ctx context.Context, id int64) (*model.SystemRole, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
