package system_role_service

import (
	"context"
	"crmeb_go/internal/common/page"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/validation"
)

type Service interface {
	List(ctx context.Context, req *validation.SystemRoleSearch) (*page.CommonPageResp[response.SystemRole], error)
}
