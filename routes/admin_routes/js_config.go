package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/admin_login_handler"
	"crmeb_go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func jsConfigRouter(casbinM *middleware.CasbinM, privateRouter *gin.RouterGroup, handler *admin_login_handler.Handler) {
	privateRouter = privateRouter.Group("/public/jsconfig")

	{
		privateRouter.GET("/getcrmebchatconfig", casbinM.CasbinMiddleware("public:jsconfig:getcrmebchatconfig"), handler.GetMenus)
	}
}
