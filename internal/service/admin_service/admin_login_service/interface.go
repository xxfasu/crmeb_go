package admin_login_service

import (
	"context"
	"crmeb_go/internal/common/data/login_user"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/validation"
)

type Service interface {
	GetCode(ctx context.Context) (*response.ValidateCodeResp, error)
	SystemAdminLogin(ctx context.Context, req *validation.SystemAdminLogin, ip string) (*response.SystemLoginResp, error)
	SystemAdminLogout(ctx context.Context, token string) error
	GetAdminInfo(ctx context.Context, loginUserData login_user.LoginUserData) (*response.SystemAdminResp, error)
	GetLoginPic(ctx context.Context) (*response.SystemLoginPicResp, error)
	GetMenus(ctx context.Context, loginUserData login_user.LoginUserData) ([]*response.SystemMenusResp, error)
}
