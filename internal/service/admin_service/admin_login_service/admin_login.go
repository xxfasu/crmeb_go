package admin_login_service

import (
	"context"
	"crmeb_go/constants"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_admin_repository"
	"crmeb_go/internal/service/common_service/system_config_service"
	"crmeb_go/internal/service/common_service/system_group_data_service"
	"crmeb_go/internal/service/common_service/system_menu_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/captcha"
	"crmeb_go/pkg/jwt"
	"crmeb_go/pkg/util"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"slices"
	"strings"
)

func New(
	tm repository.Transaction,
	captcha captcha.Captcha,
	jwt *jwt.JWT,
	systemMenuService system_menu_service.Service,
	systemConfigService system_config_service.Service,
	systemGroupDataService system_group_data_service.Service,
	systemAdminRepo system_admin_repository.Repository,
) Service {
	return &service{
		tm:                     tm,
		captcha:                captcha,
		jwt:                    jwt,
		systemMenuService:      systemMenuService,
		systemConfigService:    systemConfigService,
		systemGroupDataService: systemGroupDataService,
		systemAdminRepo:        systemAdminRepo,
	}
}

type service struct {
	tm                     repository.Transaction
	captcha                captcha.Captcha
	jwt                    *jwt.JWT
	systemMenuService      system_menu_service.Service
	systemConfigService    system_config_service.Service
	systemGroupDataService system_group_data_service.Service
	systemAdminRepo        system_admin_repository.Repository
}

func (s *service) GetCode(ctx context.Context) (*response.ValidateCode, error) {
	resp := new(response.ValidateCode)
	key, code, err := s.captcha.Gen()
	if err != nil {
		return resp, err
	}
	resp.Key = key
	resp.Code = code
	return resp, nil
}

func (s *service) SystemAdminLogin(ctx context.Context, req *validation.SystemAdminLogin, ip string) (*response.SystemLogin, error) {
	resp := new(response.SystemLogin)
	if !s.captcha.Verify(req.Key, req.Code) {
		return resp, errors.New("验证码错误")
	}
	systemAdmin, err := s.systemAdminRepo.GetUser(ctx, req.Account)
	if err != nil {
		return resp, errors.New("用户不存在")
	}
	if !util.ComparePasswords(systemAdmin.Pwd, req.Pwd) {
		return resp, errors.New("密码错误")
	}
	token, err := s.getToken(ctx, systemAdmin)
	if err != nil {
		return resp, err
	}
	resp.Token = token
	err = copier.Copy(resp, systemAdmin)
	if err != nil {
		return resp, err
	}
	umap := make(map[string]interface{}, 2)
	umap["login_count"] = systemAdmin.LoginCount + 1
	umap["last_ip"] = ip
	err = s.systemAdminRepo.UpdateFields(ctx, systemAdmin.ID, umap)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) SystemAdminLogout(ctx context.Context, token string) error {
	return s.jwt.DeleteToken(token)
}

func (s *service) GetAdminInfo(ctx context.Context, loginUserData data.LoginUser) (*response.SystemAdmin, error) {
	resp := new(response.SystemAdmin)
	systemAdmin := loginUserData.User
	err := resp.ConvertFromModel(systemAdmin)
	if err != nil {
		return resp, err
	}
	roleList := strings.Split(systemAdmin.Roles, ",")
	permList := make([]string, 0)
	if slices.Contains(roleList, "1") {
		permList = append(permList, "*:*:*")
	} else {
		permList = lo.Map(loginUserData.Permissions, func(item *model.SystemPermissions, index int) string {
			return item.Path
		})
	}
	resp.PermissionsList = permList
	return resp, nil
}

func (s *service) GetLoginPic(ctx context.Context) (*response.SystemLoginPic, error) {
	resp := new(response.SystemLoginPic)
	// 背景图
	resp.BackgroundImage = s.systemConfigService.GetValueByKey(ctx, constants.ConfigKeyAdminLoginBackgroundImage)
	// logo
	resp.Logo = s.systemConfigService.GetValueByKey(ctx, constants.ConfigKeyAdminLoginLogoLeftTop)
	resp.LoginLogo = s.systemConfigService.GetValueByKey(ctx, constants.ConfigKeyAdminLoginLogoLogin)
	// 轮播图
	list, err := s.systemGroupDataService.GetListByGID(ctx, constants.GroupDataIDAdminLoginBannerImageList)
	if err != nil {
		return resp, err
	}
	bannerList, err := util.ConvertSlice[response.SystemGroupDataAdminLoginBanner](list)
	if err != nil {
		return resp, err
	}
	resp.Banner = bannerList
	return resp, nil
}

func (s *service) GetMenus(ctx context.Context, loginUserData data.LoginUser) ([]*response.SystemMenu, error) {
	systemAdmin := loginUserData.User
	roleList := strings.Split(systemAdmin.Roles, ",")
	menuList := make([]*model.SystemMenu, 0)
	var err error
	if slices.Contains(roleList, "1") {
		menuList, err = s.systemMenuService.GetAllMenus(ctx)
	} else {
		menuList, err = s.systemMenuService.GetUserMenus(ctx, systemAdmin.ID)
	}
	if err != nil {
		return nil, err
	}

	return buildTree(menuList), nil
}

func buildTree(menuList []*model.SystemMenu) []*response.SystemMenu {
	menuMap := lo.SliceToMap(menuList, func(item *model.SystemMenu) (int64, *response.SystemMenu) {
		temp := new(response.SystemMenu)
		temp.ConvertFromModel(item)
		temp.ChildList = make([]*response.SystemMenu, 0)
		return temp.ID, temp
	})
	menuTree := make([]*response.SystemMenu, 0)
	// 第二次遍历，建立父子关系
	for _, menu := range menuList {
		if menu.Pid == 0 { // 或者其他表示顶级菜单的条件
			systemMenu := menuMap[menu.ID]
			menuTree = append(menuTree, systemMenu)
		} else if parentMenu, exists := menuMap[menu.Pid]; exists {
			if parentMenu.ChildList == nil {
				parentMenu.ChildList = make([]*response.SystemMenu, 0)
			}
			parentMenu.ChildList = append(parentMenu.ChildList, menuMap[menu.ID])
		}
	}
	return menuTree
}
