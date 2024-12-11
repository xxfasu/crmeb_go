package admin_login_service

import (
	"context"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/validation"
)

type Service interface {
	SystemAdminLogin(ctx context.Context, req *validation.Login, ip string) (response.SystemLoginResp, error)
}
