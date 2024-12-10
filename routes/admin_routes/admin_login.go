package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/user_handler"
	"github.com/gin-gonic/gin"
)

func adminLoginRouter(publicRouter *gin.RouterGroup, privateRouter *gin.RouterGroup, userHandler *user_handler.UserHandler) {
	publicRouter = publicRouter.Group("/admin")
	privateRouter = privateRouter.Group("/admin")
	{
		publicRouter.POST("/login", userHandler.Login)
		publicRouter.POST("/logout", userHandler.Register)
		publicRouter.POST("/getAdminInfoByToken", userHandler.FindUser)
		publicRouter.POST("/getLoginPic", userHandler.FindUser)
		publicRouter.POST("/getMenus", userHandler.FindUser)
	}
}
