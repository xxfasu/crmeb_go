package casbin_service

type Service interface {
	Enforce(roleID, obj, act string) (bool, error)
	UpdateCasbinApi(oldPerms, newPerms string) error
	RemoveFilteredPolicy(roleID string) error
	AddPolicies(rules [][]string) error
	FreshCasbin() error
}
