package casbin

import (
	"crmeb_go/internal/model"
	"github.com/casbin/gorm-adapter/v3"
	"log"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

// InitCasbinEnforcer 初始化Casbin
func InitCasbinEnforcer(db *gorm.DB) *casbin.Enforcer {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("failed to create gorm adapter: %v", err)
	}

	e, err := casbin.NewEnforcer("rbac_model.conf", a)
	if err != nil {
		log.Fatalf("failed to create enforcer: %v", err)
	}

	err = e.LoadPolicy()
	if err != nil {
		log.Fatalf("failed to load policy: %v", err)
	}

	return e
}

// SyncPolicies 同步数据库中的角色-菜单关系到Casbin的policy中
func SyncPolicies(db *gorm.DB, e *casbin.Enforcer) error {
	var roles []model.SystemRole
	if err := db.Find(&roles).Error; err != nil {
		return err
	}

	var menus []model.SystemMenu
	if err := db.Find(&menus).Error; err != nil {
		return err
	}

	menuPermsMap := make(map[int64]string)
	for _, m := range menus {
		menuPermsMap[m.ID] = m.Perms
	}

	// 清空原有策略
	_, err := e.RemoveFilteredPolicy(0)
	if err != nil {
		return err
	}

	for _, role := range roles {
		// 查询角色关联的菜单
		var roleMenus []model.SystemRoleMenu
		if err := db.Where("rid = ?", role.ID).Find(&roleMenus).Error; err != nil {
			return err
		}

		for _, rm := range roleMenus {
			perms, ok := menuPermsMap[rm.MenuID]
			if !ok || perms == "" {
				continue
			}

			// 根据权限字符串判断act
			// 假设菜单权限是"view"，按钮权限是"click"
			act := "view"
			if isButtonPerm(perms) {
				act = "click"
			}

			// 添加策略: p, sub(角色), obj(权限标识), act(动作)
			_, err = e.AddPolicy(role.RoleName, perms, act)
			if err != nil {
				return err
			}
		}
	}

	return e.SavePolicy()
}

func isButtonPerm(perms string) bool {
	// 简单判断，假设以"button:"开头则是按钮权限
	return len(perms) > 7 && perms[:7] == "button:"
}
