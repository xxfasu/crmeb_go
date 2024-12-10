package admin_service

import (
	"crmeb_go/internal/service/common_service/user_service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(user_service.NewUserService)
