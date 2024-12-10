package v1

import (
	"crmeb_go/internal/handler/admin_handler/v1/user_handler"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(user_handler.NewUserHandler)
