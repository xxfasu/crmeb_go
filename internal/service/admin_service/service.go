package admin_service

import (
	"crmeb_go/internal/service/admin_service/admin_login_service"
	"crmeb_go/internal/service/common_service/system_config_service"
	"crmeb_go/internal/service/common_service/system_group_data_service"
	"crmeb_go/internal/service/common_service/system_menu_service"
	"crmeb_go/internal/service/common_service/system_store_service"
	"crmeb_go/internal/service/common_service/system_store_staff_service"
	"crmeb_go/internal/service/common_service/user_service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	admin_login_service.New,

	system_config_service.New,
	system_group_data_service.New,
	system_menu_service.New,
	system_store_service.New,
	system_store_staff_service.New,
	user_service.New,
)
