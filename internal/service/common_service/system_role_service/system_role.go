package system_role_service

import (
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_role_repository"
)

func New(
	tm repository.Transaction,
	systemRoleRepo system_role_repository.Repository,
) Service {
	return &service{
		tm:             tm,
		systemRoleRepo: systemRoleRepo,
	}
}

type service struct {
	tm             repository.Transaction
	systemRoleRepo system_role_repository.Repository
}
