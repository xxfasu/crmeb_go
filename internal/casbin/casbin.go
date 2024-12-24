package casbin

import (
	"crmeb_go/pkg/logs"
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
)

type service struct {
	e  *casbin.SyncedCachedEnforcer
	db *gorm.DB
}

// InitCasbinEnforcer 初始化Casbin
func InitCasbinEnforcer(db *gorm.DB) (Service, error) {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("failed to create gorm adapter: %v", err)
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
	m, err := casbinmodel.NewModelFromString(text)
	if err != nil {
		logs.Log.Error("failed to create casbin model", zap.Error(err))
		return nil, err
	}
	e, err := casbin.NewSyncedCachedEnforcer(m, a)
	if err != nil {
		logs.Log.Error("failed to create enforcer:", zap.Error(err))
	}
	e.SetExpireTime(60 * 60)

	err = e.LoadPolicy()
	if err != nil {
		log.Fatalf("failed to load policy: %v", err)
	}

	return &service{e, db}, nil
}

func (s *service) Enforce(role, obj, act string) (bool, error) {
	return s.e.Enforce(role, obj, act)
}

func (s *service) UpdateCasbinApi(oldPerms, newPerms string) error {
	err := s.db.Model(&gormadapter.CasbinRule{}).Where("v1 = ?", oldPerms).Update("v1", newPerms).Error
	if err != nil {
		return err
	}
	return s.e.LoadPolicy()
}

func (s *service) RemoveFilteredPolicy(roleID string) error {
	err := s.db.Delete(&gormadapter.CasbinRule{}, "v0 = ?", roleID).Error
	if err != nil {
		return err
	}
	return s.e.LoadPolicy()
}

func (s *service) AddPolicies(rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    "ALL",
		})
	}
	return s.db.Create(&casbinRules).Error
}

func (s *service) FreshCasbin() error {
	return s.e.LoadPolicy()
}
