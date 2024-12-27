package system_role_service

import (
	"context"
	"crmeb_go/internal/casbin"
	"crmeb_go/internal/common/page"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_role_repository"
	"crmeb_go/internal/service/common_service/system_menu_service"
	"crmeb_go/internal/service/common_service/system_role_menu_service"
	"crmeb_go/internal/validation"
	"errors"
	"github.com/samber/lo"
	"slices"
	"strconv"
)

func New(
	tm repository.Transaction,
	systemRoleRepo system_role_repository.Repository,
	casbinService casbin.Service,
	systemMenuService system_menu_service.Service,
	systemRoleMenuService system_role_menu_service.Service,
) Service {
	return &service{
		tm:                    tm,
		systemRoleRepo:        systemRoleRepo,
		casbinService:         casbinService,
		systemMenuService:     systemMenuService,
		systemRoleMenuService: systemRoleMenuService,
	}
}

type service struct {
	tm                    repository.Transaction
	systemRoleRepo        system_role_repository.Repository
	casbinService         casbin.Service
	systemMenuService     system_menu_service.Service
	systemRoleMenuService system_role_menu_service.Service
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

func (s *service) Info(ctx context.Context, id string) (*response.RoleInfo, error) {
	roleID, err := strconv.Atoi(id)
	systemRole, err := s.systemRoleRepo.GetByID(ctx, int64(roleID))
	if err != nil {
		return nil, err
	}
	if systemRole == nil {
		return nil, errors.New("角色不存在")
	}
	// 查询角色对应的菜单(权限)
	menuList, err := s.systemMenuService.GetCacheList(ctx)
	menuIDList, err := s.systemRoleMenuService.GetMenuIDList(ctx, int64(roleID))
	menuCheckList := lo.Map(menuList, func(item *model.SystemMenu, index int) *response.MenuCheck {
		menuCheck := new(response.MenuCheck)
		menuCheck.ConvertFromModel(item)
		if slices.Contains(menuIDList, item.ID) {
			menuCheck.Checked = 1
		}
		return menuCheck
	})
	resp := new(response.RoleInfo)
	err = resp.ConvertFromModel(systemRole)
	if err != nil {
		return nil, err
	}
	resp.MenuList = s.systemMenuService.BuildTree(menuCheckList)
	return resp, nil
}
