package system_admin_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/validation"
)

type Reader interface {
	GetUser(ctx context.Context, userName string) (*model.SystemAdmin, error)
	GetByID(ctx context.Context, id int64) (*model.SystemAdmin, error)
	GetPage(ctx context.Context, condition *validation.SystemAdminSearch) ([]*model.SystemAdmin, int64, error)
	IsExistAccount(ctx context.Context, account string) (bool, error)
	IsSingleAccount(ctx context.Context, id int64, account string) (bool, error)
}

type Writer interface {
	Create(ctx context.Context, entity *model.SystemAdmin) error
	Update(ctx context.Context, entity *model.SystemAdmin) error
	UpdateFields(ctx context.Context, ID int64, umap map[string]interface{}) error
	DeleteByID(ctx context.Context, ID int64) error
}

type Repository interface {
	Reader
	Writer
}
