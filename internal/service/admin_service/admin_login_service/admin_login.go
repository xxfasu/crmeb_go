package admin_login_service

import (
	"context"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/captcha"
	"crmeb_go/pkg/jwt"
)

func New(
	tm repository.Transaction,
	jwt *jwt.JWT,
	captcha captcha.Captcha,
) Service {
	return &service{
		tx:      tm,
		jwt:     jwt,
		captcha: captcha,
	}
}

type service struct {
	tx      repository.Transaction
	jwt     *jwt.JWT
	captcha captcha.Captcha
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
	return response.SystemLoginResp{}, nil
}
