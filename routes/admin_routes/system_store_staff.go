package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/system_store_staff_handler"
	"crmeb_go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func systemStoreStaffRouter(casbinM *middleware.CasbinM, privateRouter *gin.RouterGroup, handler *system_store_staff_handler.Handler) {
	privateRouter = privateRouter.Group("/admin/system/store/staff")

	{
		privateRouter.GET("/list", casbinM.CasbinMiddleware("admin:system:staff:list"), handler.GetList)
		privateRouter.POST("/save", casbinM.CasbinMiddleware("admin:system:staff:save"), handler.GetList)
		privateRouter.GET("/delete", casbinM.CasbinMiddleware("admin:system:staff:delete"), handler.GetList)
		privateRouter.POST("/update", casbinM.CasbinMiddleware("admin:system:staff:update"), handler.GetList)
		privateRouter.GET("/update/status", casbinM.CasbinMiddleware("admin:system:staff:update:status"), handler.GetList)
		privateRouter.GET("/info", casbinM.CasbinMiddleware("admin:system:staff:info"), handler.GetList)
	}
}
