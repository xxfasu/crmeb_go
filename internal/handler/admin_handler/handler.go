package admin_handler

import (
	"crmeb_go/internal/handler/admin_handler/v1/admin_login_handler"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(admin_login_handler.New)
