package system_group_data_service

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/validation"
)

type Service interface {
	GetList(ctx context.Context, req validation.SystemGroupDataSearch) ([]*model.SystemGroupData, error)
	GetListByGID(ctx context.Context, gid int64) ([]any, error)
}
