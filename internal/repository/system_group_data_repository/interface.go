package system_group_data_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/validation"
)

type Reader interface {
	GetGroupDataPageList(ctx context.Context, condition validation.SystemGroupDataSearchReq) ([]*model.SystemGroupData, int64, error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
