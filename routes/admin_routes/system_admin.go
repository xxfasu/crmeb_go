package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/system_admin_handler"
	"crmeb_go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func systemAdminRouter(casbinM *middleware.CasbinM, privateRouter *gin.RouterGroup, handler *system_admin_handler.Handler) {
	privateRouter = privateRouter.Group("/admin")

	{
		privateRouter.GET("/list", casbinM.CasbinMiddleware("admin:system:admin:list"), handler.List)
		privateRouter.POST("/save", casbinM.CasbinMiddleware("admin:system:admin:save"), handler.Save)
		privateRouter.GET("/delete", casbinM.CasbinMiddleware("admin:system:admin:delete"), handler.Delete)
		privateRouter.POST("/update", casbinM.CasbinMiddleware("admin:system:admin:update"), handler.Update)
		privateRouter.GET("/info", casbinM.CasbinMiddleware("admin:system:admin:info"), handler.Info)
		privateRouter.GET("/updateStatus", casbinM.CasbinMiddleware("admin:system:admin:update:status"), handler.UpdateStatus)
		privateRouter.GET("/update/isSms", casbinM.CasbinMiddleware("admin:system:admin:update:sms"), handler.UpdateSms)
	}
}
