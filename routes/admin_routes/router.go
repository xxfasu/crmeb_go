package admin_routes

import (
	"crmeb_go/internal/conf"
	"crmeb_go/internal/handler/admin_handler/v1/admin_login_handler"
	"crmeb_go/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var ProviderSet = wire.NewSet(NewRouter)

func NewRouter(
	recoveryM *middleware.Recovery,
	corsM *middleware.Cors,
	logM *middleware.LogM,
	authM *middleware.AuthM,
	casbinM *middleware.CasbinM,
	adminLoginHandler *admin_login_handler.Handler,
) *gin.Engine {
	router := gin.New()
	if true {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Use(recoveryM.Handler())
	router.Use(corsM.Handler())
	router.Use(logM.RequestLogMiddleware())
	router.Use(logM.ResponseLogMiddleware())

	publicGroup := router.Group("/api")
	publicGroup.Use(authM.NoStrictAuth())
	privateGroup := router.Group("/api")
	privateGroup.Use(authM.StrictAuth())

	// 提供静态文件，访问路径为 /crmebimage
	router.Static("/crmebimage", conf.Config.CrmebConfig.ImagePath)

	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		adminLoginRouter(casbinM, publicGroup, privateGroup, adminLoginHandler)
	}
	return router
}
