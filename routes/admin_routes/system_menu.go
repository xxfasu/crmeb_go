package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/system_menu_handler"
	"crmeb_go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func systemMenuRouter(casbinM *middleware.CasbinM, privateRouter *gin.RouterGroup, handler *system_menu_handler.Handler) {
	privateRouter = privateRouter.Group("/admin/system/menu")

	{
		privateRouter.GET("/cache/tree", casbinM.CasbinMiddleware("admin:system:menu:cache:tree"), handler.CacheTree)
	}
}
