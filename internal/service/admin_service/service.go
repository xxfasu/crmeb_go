package admin_service

import (
	"crmeb_go/internal/service/admin_service/admin_login_service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	admin_login_service.New,
)
