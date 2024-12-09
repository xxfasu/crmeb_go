package admin_routes

import (
	"crmeb_go/internal/handler/admin_handler/v1/user_handler"
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
	userHandler *user_handler.UserHandler,
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
	router.Use(authM.NoStrictAuth())

	publicGroup := router.Group("/api")
	publicGroup.Use(authM.NoStrictAuth())
	privateGroup := router.Group("/api")
	privateGroup.Use(authM.StrictAuth())

	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		adminLoginRouter(publicGroup, privateGroup, userHandler)
	}
	return router
}
