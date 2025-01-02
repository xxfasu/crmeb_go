package handler

import (
	"crmeb_go/internal/conf"
	"crmeb_go/internal/middleware"
	redis2 "crmeb_go/internal/redis"
	"crmeb_go/pkg/cache"
	"crmeb_go/pkg/jwt"
	"crmeb_go/pkg/logs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"os"
	"testing"
)

var (
	userId = "xxx"
)

var router *gin.Engine
var jwtJWT *jwt.JWT
var rClient redis.UniversalClient

func TestMain(m *testing.M) {
	fmt.Println("begin")
	var err error
	conf.InitConfig("/Users/xxfs/GolandProjects/crmeb_go/config")
	if err != nil {
		panic(err)
	}
	logs.InitLog()
	rClient, err = redis2.InitRedis()
	if err != nil {
		panic(err)
	}
	cacheCache := cache.InitLocalCache(rClient)
	jwtJWT = jwt.NewJwt(cacheCache)
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	logM := middleware.NewLogM()
	recoveryM := middleware.NewRecoveryM()
	corsM := middleware.NewCorsM()
	router.Use(logM.Handler())
	router.Use(recoveryM.Handler())
	router.Use(corsM.Handler())

	code := m.Run()
	fmt.Println("test end")
	os.Exit(code)

}
