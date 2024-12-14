package admin_login_service

import (
	"context"
	"crmeb_go/internal/common/data/login_user"
	"crmeb_go/internal/model"
	"github.com/samber/lo"
	"slices"
	"strings"
)

func getToken(ctx context.Context, s *service, systemAdmin *model.SystemAdmin) (string, error) {
	roles := strings.Split(systemAdmin.Roles, ",")
	var menuList []*model.SystemMenu
	var err error
	if slices.Contains(roles, "1") { // 超级管理员
		menuList, err = s.systemMenuService.GetAllPermissions(ctx)
		if err != nil {
			return "", err
		}
	} else {
		menuList, err = s.systemMenuService.GetUserPermissions(ctx, systemAdmin.ID)
		if err != nil {
			return "", err
		}
	}
	menuList = lo.Filter(menuList, func(item *model.SystemMenu, index int) bool {
		return lo.IsNotEmpty(item.Perms)
	})
	permissionsList := lo.Map(menuList, func(item *model.SystemMenu, index int) *model.SystemPermissions {
		return &model.SystemPermissions{
			ID:   item.ID,
			Pid:  item.Pid,
			Path: item.Perms,
			Name: item.Name,
			Sort: item.Sort,
		}
	})
	loginUser := login_user.LoginUserData{User: systemAdmin, Permissions: permissionsList}
	token, err := s.jwt.CreateToken(loginUser)
	return token, nil
}
