package system_role_service

import (
	"context"
	"crmeb_go/internal/common/page"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/validation"
)

type Service interface {
	List(ctx context.Context, req *validation.SystemRoleSearch) (*page.CommonPageResp[response.SystemRole], error)
	Save(ctx context.Context, req *validation.SystemRole) error
	Info(ctx context.Context, id int64) (*response.RoleInfo, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, req *validation.SystemRole) error
	UpdateStatus(ctx context.Context, id, status int64) error
}
