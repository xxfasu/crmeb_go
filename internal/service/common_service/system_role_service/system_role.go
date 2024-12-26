package system_role_service

import (
	"context"
	"crmeb_go/internal/casbin"
	"crmeb_go/internal/common/page"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_role_repository"
	"crmeb_go/internal/validation"
)

func New(
	tm repository.Transaction,
	systemRoleRepo system_role_repository.Repository,
	casbinService casbin.Service,
) Service {
	return &service{
		tm:             tm,
		systemRoleRepo: systemRoleRepo,
		casbinService:  casbinService,
	}
}

type service struct {
	tm             repository.Transaction
	systemRoleRepo system_role_repository.Repository
	casbinService  casbin.Service
}

func (s *service) List(ctx context.Context, req *validation.SystemRoleSearch) (*page.CommonPageResp[response.SystemRole], error) {
	systemRoleList, total, err := s.systemRoleRepo.GetList(ctx, req)
	if err != nil {
		return nil, err
	}
	systemRoleRespList := make([]response.SystemRole, 0, len(systemRoleList))
	for _, systemRole := range systemRoleList {
		var systemRoleResp response.SystemRole
		err = systemRoleResp.ConvertFromModel(systemRole)
		if err != nil {
			return nil, err
		}
		systemRoleRespList = append(systemRoleRespList, systemRoleResp)
	}

	resp := page.RestPage(req.PageParam, systemRoleRespList, total)
	return resp, nil
}
