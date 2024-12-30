package system_admin_service

import (
	"context"
	"crmeb_go/internal/common/page"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/validation"
)

type Service interface {
	List(ctx context.Context, req *validation.SystemAdminSearch) (*page.CommonPageResp[response.SystemAdmin], error)
	Save(ctx context.Context, req *validation.SystemAdminAdd) error
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, req *validation.SystemAdminUpdate) error
	Info(ctx context.Context, id int64) (*response.SystemAdmin, error)
	UpdateStatus(ctx context.Context, id, status int64) error
	UpdateIsSms(ctx context.Context, id int64) error
}
