package system_admin_repository

import (
	"context"
	"crmeb_go/internal/model"
)

type Reader interface {
	GetUser(ctx context.Context, userName string) (*model.SystemAdmin, error)
}

type Writer interface {
	Update(ctx context.Context, ID int64, umap map[string]interface{}) error
}

type Repository interface {
	Reader
	Writer
}
