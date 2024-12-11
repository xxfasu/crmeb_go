package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/admin_login_handler"
	"github.com/gin-gonic/gin"
)

func adminLoginRouter(publicRouter *gin.RouterGroup, privateRouter *gin.RouterGroup, handler *admin_login_handler.Handler) {
	publicRouter = publicRouter.Group("/admin")
	privateRouter = privateRouter.Group("/admin")
	{
		publicRouter.POST("/login", handler.Login)
		publicRouter.POST("/logout", handler.Login)
		publicRouter.POST("/getAdminInfoByToken", handler.Login)
		publicRouter.POST("/getLoginPic", handler.Login)
		publicRouter.POST("/getMenus", handler.Login)
	}
}
