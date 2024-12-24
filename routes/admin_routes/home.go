package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/home_handler"
	"crmeb_go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func homeRouter(casbinM *middleware.CasbinM, privateRouter *gin.RouterGroup, handler *home_handler.Handler) {
	privateRouter = privateRouter.Group("/admin/statistics/home")

	{
		privateRouter.GET("/index", casbinM.CasbinMiddleware("admin:statistics:home:index"), handler.IndexDate)
		privateRouter.GET("/chart/user", casbinM.CasbinMiddleware("admin:statistics:home:chart:user"), handler.ChartUser)
		privateRouter.GET("/chart/order", casbinM.CasbinMiddleware("admin:statistics:home:chart:order"), handler.ChartOrder)
		privateRouter.GET("/chart/order/week", casbinM.CasbinMiddleware("admin:statistics:home:chart:order:week"), handler.ChartOrderInWeek)
		privateRouter.GET("/chart/order/month", casbinM.CasbinMiddleware("admin:statistics:home:chart:order:month"), handler.ChartOrderInMonth)
		privateRouter.GET("/chart/order/year", casbinM.CasbinMiddleware("admin:statistics:home:chart:order:year"), handler.ChartOrderInYear)
	}
}
