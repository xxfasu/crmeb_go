package admin_handler

import (
	"crmeb_go/internal/handler/admin_handler/v1/admin_login_handler"
	"crmeb_go/internal/handler/admin_handler/v1/home_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_admin_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_config_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_menu_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_role_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_store_staff_handler"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	admin_login_handler.New,
	home_handler.New,
	system_config_handler.New,
	system_store_staff_handler.New,
	system_role_handler.New,
	system_menu_handler.New,
	system_admin_handler.New,
)
