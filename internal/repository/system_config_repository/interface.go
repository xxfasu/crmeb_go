package system_config_repository

import (
	"context"
	"crmeb_go/internal/model"
)

type Reader interface {
	GetConfigByName(ctx context.Context, name string) (*model.SystemConfig, error)
	GetConfigALL(ctx context.Context) ([]*model.SystemConfig, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
