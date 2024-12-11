package admin_login_service

import (
	"context"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/jwt"
)

func New(
	tm repository.Transaction,
	jwt *jwt.JWT,
) Service {
	return &service{
		tx:  tm,
		jwt: jwt,
	}
}

type service struct {
	tx  repository.Transaction
	jwt *jwt.JWT
}

func (s *service) SystemAdminLogin(ctx context.Context, req *validation.Login, ip string) (response.SystemLoginResp, error) {
	return response.SystemLoginResp{}, nil
}
