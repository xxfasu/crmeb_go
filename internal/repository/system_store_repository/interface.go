package system_store_repository

import (
	"context"
	"crmeb_go/internal/model"
)

type Reader interface {
	GetStoreList(ctx context.Context, storeIDList []int64) ([]*model.SystemStore, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
