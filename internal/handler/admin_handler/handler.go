package admin_handler

import (
	"crmeb_go/internal/handler/admin_handler/v1/admin_login_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_store_staff_handler"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	admin_login_handler.New,
	system_store_staff_handler.New,
)
