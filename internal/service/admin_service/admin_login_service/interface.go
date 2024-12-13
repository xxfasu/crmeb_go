package admin_login_service

import (
	"context"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/validation"
)

type Service interface {
	GetCode(ctx context.Context) (response.ValidateCodeResp, error)
	SystemAdminLogin(ctx context.Context, req *validation.SystemAdminLogin, ip string) (response.SystemLoginResp, error)
}
