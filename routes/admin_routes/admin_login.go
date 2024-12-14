package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/admin_login_handler"
	"crmeb_go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func adminLoginRouter(casbinM *middleware.CasbinM, publicRouter *gin.RouterGroup, privateRouter *gin.RouterGroup, handler *admin_login_handler.Handler) {
	publicRouter = publicRouter.Group("/admin")
	privateRouter = privateRouter.Group("/admin")
	{
		publicRouter.GET("/validate/code/get", handler.GetCode) // 获取验证码
		publicRouter.POST("/login", handler.Login)
		publicRouter.GET("/getLoginPic", handler.GetLoginPic)
	}

	{
		privateRouter.POST("/logout", casbinM.CasbinMiddleware("admin:logout"), handler.Logout)
		privateRouter.POST("/getAdminInfoByToken", casbinM.CasbinMiddleware("admin:info"), handler.GetAdminInfo)
		privateRouter.GET("/getMenus", casbinM.CasbinMiddleware("admin:login:menus"), handler.GetMenus)
	}
}
