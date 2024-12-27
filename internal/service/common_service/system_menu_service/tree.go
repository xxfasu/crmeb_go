package system_menu_service

import (
	"crmeb_go/internal/common/response"
	"github.com/samber/lo"
	"sort"
)

func (s *service) BuildTree(menuList []*response.MenuCheck) []*response.MenuCheck {
	menuMap := lo.SliceToMap(menuList, func(item *response.MenuCheck) (int64, *response.MenuCheck) {
		return item.ID, item
	})
	menuTree := make([]*response.MenuCheck, 0)
	// 第二次遍历，建立父子关系
	for _, menu := range menuList {
		if menu.Pid == 0 { // 或者其他表示顶级菜单的条件
			menuCheck := menuMap[menu.ID]
			menuTree = append(menuTree, menuCheck)
		} else if parentMenu, exists := menuMap[menu.Pid]; exists {
			if parentMenu.ChildList == nil {
				parentMenu.ChildList = make([]*response.MenuCheck, 0)
			}
			parentMenu.ChildList = append(parentMenu.ChildList, menuMap[menu.ID])
		}
	}
	s.sortList(menuTree)
	return menuTree
}

func (s *service) sortList(menuTree []*response.MenuCheck) {
	sort.Slice(menuTree, func(i, j int) bool {
		// 如果需要升序，改成 <；降序，则使用 >
		return menuTree[i].Sort > menuTree[j].Sort
	})
	for _, menu := range menuTree {
		if menu.ChildList != nil {
			s.sortList(menu.ChildList)
		}
	}
}
