//go:build wireinject
// +build wireinject

package admin

import (
	"crmeb_go/internal/casbin"
	"crmeb_go/internal/handler/admin_handler/v1"
	"crmeb_go/internal/middleware"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/service/common_service"
	"crmeb_go/pkg/cache"
	"crmeb_go/pkg/jwt"
	"crmeb_go/routes/admin_routes"
	"github.com/gin-gonic/gin"
	"github.com/go-redsync/redsync/v4"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func newWire(client *redis.Client, rLock *redsync.Redsync) (*gin.Engine, func(), error) {
	panic(wire.Build(
		middleware.ProviderSet,
		repository.ProviderSet,
		common_service.ProviderSet,
		v1.ProviderSet,
		admin_routes.ProviderSet,
		jwt.NewJwt,
		casbin.InitCasbinEnforcer,
		cache.InitLocalCache,
	))
}
