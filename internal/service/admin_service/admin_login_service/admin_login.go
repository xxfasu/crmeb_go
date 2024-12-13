package admin_login_service

import (
	"context"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/service/admin_service/system_menu_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/captcha"
	"errors"
)

func New(
	tm repository.Transaction,
	captcha captcha.Captcha,
	systemMenuService system_menu_service.Service,
) Service {
	return &service{
		tm:                tm,
		captcha:           captcha,
		systemMenuService: systemMenuService,
	}
}

type service struct {
	tm                repository.Transaction
	captcha           captcha.Captcha
	systemMenuService system_menu_service.Service
}

func (s *service) GetCode(ctx context.Context) (response.ValidateCodeResp, error) {
	var resp response.ValidateCodeResp
	key, code, err := s.captcha.Gen()
	if err != nil {
		return resp, err
	}
	resp.Key = key
	resp.Code = code
	return resp, nil
}

func (s *service) SystemAdminLogin(ctx context.Context, req *validation.SystemAdminLogin, ip string) (response.SystemLoginResp, error) {
	var resp response.SystemLoginResp
	if !s.captcha.Verify(req.Key, req.Code) {
		return resp, errors.New("验证码错误")
	}

	return resp, nil
}
