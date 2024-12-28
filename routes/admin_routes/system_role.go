package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/system_role_handler"
	"crmeb_go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func systemRoleRouter(casbinM *middleware.CasbinM, privateRouter *gin.RouterGroup, handler *system_role_handler.Handler) {
	privateRouter = privateRouter.Group("/admin/system/role")

	{
		privateRouter.GET("/list", casbinM.CasbinMiddleware("admin:system:role:list"), handler.List)
		privateRouter.POST("/save", casbinM.CasbinMiddleware("admin:system:role:save"), handler.Save)
		privateRouter.GET("/delete", casbinM.CasbinMiddleware("admin:system:role:delete"), handler.Delete)
		privateRouter.POST("/update", casbinM.CasbinMiddleware("admin:system:role:update"), handler.Update)
		privateRouter.GET("/info/{id}", casbinM.CasbinMiddleware("admin:system:role:info"), handler.Info)
		privateRouter.GET("/updateStatus", casbinM.CasbinMiddleware("admin:system:role:update:status"), handler.UpdateStatus)
	}
}
