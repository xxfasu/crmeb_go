package admin_routes

import (
	"crmeb_go/internal/conf"
	"crmeb_go/internal/handler/admin_handler/v1/admin_login_handler"
	"crmeb_go/internal/handler/admin_handler/v1/home_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_admin_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_config_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_menu_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_role_handler"
	"crmeb_go/internal/handler/admin_handler/v1/system_store_staff_handler"
	"crmeb_go/internal/middleware"
	"crmeb_go/pkg/validators"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	systemStoreStaffHandler *system_store_staff_handler.Handler,
	systemConfigHandler *system_config_handler.Handler,
	systemRoleHandler *system_role_handler.Handler,
	systemMenuHandler *system_menu_handler.Handler,
	systemAdminHandler *system_admin_handler.Handler,
	homeHandler *home_handler.Handler,
) *gin.Engine {
	router := gin.New()
	if conf.Env.Environment == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", validators.Phone)
	}

	router.Use(recoveryM.Handler())
	router.Use(corsM.Handler())
	router.Use(logM.Handler())

	publicGroup := router.Group("/api")
	publicGroup.Use(authM.NoStrictAuth())
	privateGroup := router.Group("/api")
	privateGroup.Use(authM.StrictAuth())

	// 提供静态文件，访问路径为 /crmebimage
	router.Static("/crmebimage", "./crmebimage")

	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})

	}

	{
		jsConfigRouter(casbinM, privateGroup, systemConfigHandler)

		adminLoginRouter(casbinM, publicGroup, privateGroup, adminLoginHandler)

		systemConfigRouter(casbinM, privateGroup, systemConfigHandler)

		systemStoreStaffRouter(casbinM, privateGroup, systemStoreStaffHandler)

		systemRoleRouter(casbinM, privateGroup, systemRoleHandler)

		systemMenuRouter(casbinM, privateGroup, systemMenuHandler)

		systemAdminRouter(casbinM, privateGroup, systemAdminHandler)

		homeRouter(casbinM, privateGroup, homeHandler)
	}
	return router
}
