package admin_login_service

import (
	"context"
	"crmeb_go/constants"
	"crmeb_go/internal/common/data/login_user"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/system_admin_repository"
	"crmeb_go/internal/service/common_service/system_config_service"
	"crmeb_go/internal/service/common_service/system_menu_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/captcha"
	"crmeb_go/pkg/jwt"
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
	systemAdminRepo system_admin_repository.Repository,
) Service {
	return &service{
		tm:                  tm,
		captcha:             captcha,
		jwt:                 jwt,
		systemMenuService:   systemMenuService,
		systemConfigService: systemConfigService,
		systemAdminRepo:     systemAdminRepo,
	}
}

type service struct {
	tm                  repository.Transaction
	captcha             captcha.Captcha
	jwt                 *jwt.JWT
	systemMenuService   system_menu_service.Service
	systemConfigService system_config_service.Service
	systemAdminRepo     system_admin_repository.Repository
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
	systemAdmin, err := s.systemAdminRepo.GetUser(ctx, req.Account)
	if err != nil {
		return resp, errors.New("用户不存在")
	}
	token, err := getToken(ctx, s, systemAdmin)
	if err != nil {
		return resp, err
	}
	resp.Token = token
	err = copier.Copy(&resp, systemAdmin)
	if err != nil {
		return resp, err
	}
	umap := make(map[string]interface{}, 2)
	umap["login_count"] = systemAdmin.LoginCount + 1
	umap["last_ip"] = ip
	err = s.systemAdminRepo.Update(ctx, systemAdmin.ID, umap)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) SystemAdminLogout(ctx context.Context, token string) error {
	return s.jwt.DeleteToken(token)
}

func (s *service) GetAdminInfo(ctx context.Context, loginUserData login_user.LoginUserData) (response.SystemAdminResp, error) {
	var resp response.SystemAdminResp
	systemAdmin := loginUserData.User
	err := copier.Copy(&resp, systemAdmin)
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
func (s *service) GetLoginPic(ctx context.Context) (response.SystemLoginPicResp, error) {
	var resp response.SystemLoginPicResp
	// 背景图
	resp.BackgroundImage = s.systemConfigService.GetValueByKey(ctx, constants.ConfigKeyAdminLoginBackgroundImage)
	// logo
	resp.Logo = s.systemConfigService.GetValueByKey(ctx, constants.ConfigKeyAdminLoginLogoLeftTop)
	resp.LoginLogo = s.systemConfigService.GetValueByKey(ctx, constants.ConfigKeyAdminLoginLogoLogin)
	// 轮播图
	resp.BackgroundImage = s.systemConfigService.GetValueByKey(ctx, "login_background_image")
	return resp, nil
}

func (s *service) GetMenus(ctx context.Context, loginUserData login_user.LoginUserData) ([]response.SystemMenusResp, error) {
	var resp []response.SystemMenusResp
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
		return resp, err
	}
	var flag error
	resp = lo.Map(menuList, func(item *model.SystemMenu, index int) response.SystemMenusResp {
		var temp response.SystemMenusResp
		if err := copier.Copy(&temp, item); err != nil {
			flag = err
			return temp
		}
		return temp
	})
	if flag != nil {
		return resp, flag
	}
	return resp, nil
}
