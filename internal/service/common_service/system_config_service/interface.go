package system_config_service

import "context"

type Service interface {
	// GetValueByKey 根据menu name 获取 value
	GetValueByKey(ctx context.Context, name string) string
	// BatchGetValueByKey 根据menu name 批量获取 value
	BatchGetValueByKey(ctx context.Context, keyList []string) []string
}
