package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/system_config_handler"
	"crmeb_go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func systemConfigRouter(casbinM *middleware.CasbinM, privateRouter *gin.RouterGroup, handler *system_config_handler.Handler) {
	privateRouter = privateRouter.Group("/admin/system/config")

	{
		privateRouter.GET("/getuniq", casbinM.CasbinMiddleware("admin:system:config:getuniq"), handler.GetUniq)
	}
}
