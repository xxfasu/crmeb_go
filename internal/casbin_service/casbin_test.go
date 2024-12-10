package casbin_service

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

const dsn = "root:password@tcp(127.0.0.1:3306)/crmeb?charset=utf8mb4&parseTime=True&loc=Local"

func TestSyncPolicies(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		t.Fatalf("failed to create gorm adapter: %v", err)
	}
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	m, err := casbinmodel.NewModelFromFile(text)
	if err != nil {
		t.Fatalf("failed to create casbin model", err)
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		t.Fatalf("failed to create enforcer: %v", err)
	}

	err = e.LoadPolicy()
	if err != nil {
		t.Fatalf("failed to load policy: %v", err)
	}
	gen.SetDefault(db)
	roles, err := gen.SystemRole.WithContext(context.Background()).Find()

	if err != nil {
		t.Fatalf("SystemRole Find : %v", err)
	}
	menus, err := gen.SystemMenu.WithContext(context.Background()).Find()
	if err != nil {
		t.Fatalf("SystemMenu Find: %v", err)
	}
	menuPermsMap := make(map[int64]string)
	for _, m := range menus {
		menuPermsMap[m.ID] = m.Perms
	}

	for _, role := range roles {
		// 查询角色关联的菜单
		var roleMenus []model.SystemRoleMenu
		if err := db.Where("rid = ?", role.ID).Find(&roleMenus).Error; err != nil {
			t.Fatalf("failed to query role menus: %v", err)
		}

		for _, rm := range roleMenus {
			perms, ok := menuPermsMap[rm.MenuID]
			if !ok || perms == "" {
				continue
			}

			// 添加策略: p, sub(角色), obj(权限标识), act(动作)
			_, err = e.AddPolicy(strconv.FormatInt(role.ID, 10), perms, "ALL")
			if err != nil {
				t.Fatalf("failed to add policy: %v", err)
			}
		}
	}

	err = e.SavePolicy()
	if err != nil {
		t.Fatalf("failed to save policy: %v", err)
	}
}

func TestInitCasbinEnforcer(t *testing.T) {
	t.Log("TestInitCasbinEnforcer")
}
