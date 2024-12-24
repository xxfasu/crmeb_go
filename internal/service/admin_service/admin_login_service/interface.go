package admin_login_service

import (
	"context"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/validation"
)

type Service interface {
	GetCode(ctx context.Context) (*response.ValidateCode, error)
	SystemAdminLogin(ctx context.Context, req *validation.SystemAdminLogin, ip string) (*response.SystemLogin, error)
	SystemAdminLogout(ctx context.Context, token string) error
	GetAdminInfo(ctx context.Context, loginUserData data.LoginUser) (*response.SystemAdmin, error)
	GetLoginPic(ctx context.Context) (*response.SystemLoginPic, error)
	GetMenus(ctx context.Context, loginUserData data.LoginUser) ([]*response.SystemMenu, error)
}
