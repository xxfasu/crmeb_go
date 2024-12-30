package system_admin_service

import (
	"context"
	"crmeb_go/internal/common/page"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_admin_repository"
	"crmeb_go/internal/service/common_service/system_role_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/util"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"strings"
)

func New(
	tm repository.Transaction,
	systemRoleService system_role_service.Service,
	systemAdminRepo system_admin_repository.Repository,
) Service {
	return &service{
		tm:                tm,
		systemAdminRepo:   systemAdminRepo,
		systemRoleService: systemRoleService,
	}
}

type service struct {
	tm                repository.Transaction
	systemRoleService system_role_service.Service
	systemAdminRepo   system_admin_repository.Repository
}

func (s *service) List(ctx context.Context, req *validation.SystemAdminSearch) (*page.CommonPageResp[response.SystemAdmin], error) {
	resp := new(page.CommonPageResp[response.SystemAdmin])
	systemAdminList, count, err := s.systemAdminRepo.GetPage(ctx, req)
	if err != nil {
		return resp, err
	}
	systemAdminRespList := make([]response.SystemAdmin, 0, len(systemAdminList))
	roleList, err := s.systemRoleService.GetAllSystemRoleList(ctx)
	if err != nil {
		return resp, err
	}
	for _, systemAdmin := range systemAdminList {
		var systemAdminResp response.SystemAdmin
		systemAdminResp.ConvertFromModel(systemAdmin)
		systemAdminResp.LastTime = util.TimeFormat(systemAdmin.UpdatedAt)
		if len(systemAdmin.Roles) == 0 {
			continue
		}
		roleIDList := util.StrToArrInt64(systemAdmin.Roles)
		roleNameList := make([]string, 0, len(roleIDList))
		for _, roleID := range roleIDList {
			hasRoleList := lo.Filter(roleList, func(item *model.SystemRole, index int) bool {
				return item.ID == roleID
			})
			if len(hasRoleList) > 0 {
				roleNames := lo.Map(hasRoleList, func(item *model.SystemRole, index int) string {
					return item.RoleName
				})
				roleNameList = append(roleNameList, strings.Join(roleNames, ","))
			}
		}
		systemAdminResp.RoleNames = strings.Join(roleNameList, ",")
		systemAdminRespList = append(systemAdminRespList, systemAdminResp)
	}
	page.RestPage(req.PageParam, systemAdminList, count)
	return resp, nil
}

func (s *service) Save(ctx context.Context, req *validation.SystemAdminAdd) error {
	exist, err := s.systemAdminRepo.IsExistAccount(ctx, req.Account)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("管理员已存在")
	}
	systemAdmin := new(model.SystemAdmin)
	copier.Copy(systemAdmin, req)
	hashPassword, err := util.HashPassword(req.Pwd)
	if err != nil {
		return err
	}
	systemAdmin.Pwd = hashPassword
	err = s.systemAdminRepo.Create(ctx, systemAdmin)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.systemAdminRepo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Update(ctx context.Context, req *validation.SystemAdminUpdate) error {
	single, err := s.systemAdminRepo.IsSingleAccount(ctx, req.ID, req.Account)
	if err != nil {
		return err
	}
	if !single {
		return errors.New("管理员已存在")
	}
	systemAdmin := new(model.SystemAdmin)
	copier.Copy(systemAdmin, req)
	if len(req.Pwd) > 0 {
		hashPassword, err := util.HashPassword(req.Pwd)
		if err != nil {
			return err
		}
		systemAdmin.Pwd = hashPassword
	}
	err = s.systemAdminRepo.Update(ctx, systemAdmin)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Info(ctx context.Context, id int64) (*response.SystemAdmin, error) {
	systemAdmin, err := s.systemAdminRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := new(response.SystemAdmin)
	resp.ConvertFromModel(systemAdmin)
	return resp, nil
}

func (s *service) UpdateStatus(ctx context.Context, id, status int64) error {
	systemAdmin, err := s.systemAdminRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if systemAdmin == nil {
		return errors.New("管理员不存在")
	}
	if systemAdmin.Status == status {
		return nil
	}
	umap := make(map[string]any)
	umap["status"] = status
	return s.systemAdminRepo.UpdateFields(ctx, id, umap)
}

func (s *service) UpdateIsSms(ctx context.Context, id int64) error {
	systemAdmin, err := s.systemAdminRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if systemAdmin == nil {
		return errors.New("管理员不存在")
	}
	if len(systemAdmin.Phone) == 0 {
		return errors.New("请先为管理员添加手机号")
	}
	umap := make(map[string]any)
	umap["is_sms"] = systemAdmin.IsSms ^ 1
	return s.systemAdminRepo.UpdateFields(ctx, id, umap)
}
