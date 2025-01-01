package service

import (
	"context"
	redis2 "crmeb_go/internal/redis"
	"crmeb_go/internal/service/admin_service/admin_login_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/cache"
	"crmeb_go/pkg/captcha"
	"crmeb_go/pkg/jwt"
	"crmeb_go/test/mocks/repository/mocks_system_admin_repository"
	"crmeb_go/test/mocks/repository/mocks_transaction"
	"crmeb_go/test/mocks/service/common_service/mocks_system_config_service"
	"crmeb_go/test/mocks/service/common_service/mocks_system_group_data_service"
	"crmeb_go/test/mocks/service/common_service/mocks_system_menu_service"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

var jwtJWT *jwt.JWT
var captchaCaptcha captcha.Captcha
var rClient redis.UniversalClient
var mockSystemAdminRepo *mocks_system_admin_repository.MockRepository
var mockTransaction *mocks_transaction.MockTransaction
var systemMenuService *mocks_system_menu_service.MockService
var systemConfigService *mocks_system_config_service.MockService
var systemGroupDataService *mocks_system_group_data_service.MockService

func TestMain(m *testing.M) {
	fmt.Println("begin")
	var err error
	rClient, err = redis2.InitRedis()
	if err != nil {
		panic(err)
	}
	cacheCache := cache.InitLocalCache(rClient)
	jwtJWT = jwt.NewJwt(cacheCache)
	captchaCaptcha = captcha.New(cacheCache)

	code := m.Run()
	fmt.Println("test end")
	os.Exit(code)
}

func NewService(t *testing.T) admin_login_service.Service {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSystemAdminRepo = mocks_system_admin_repository.NewMockRepository(ctrl)
	mockTransaction = mocks_transaction.NewMockTransaction(ctrl)
	systemMenuService = mocks_system_menu_service.NewMockService(ctrl)
	systemConfigService = mocks_system_config_service.NewMockService(ctrl)
	systemGroupDataService = mocks_system_group_data_service.NewMockService(ctrl)
	adminLoginService := admin_login_service.New(mockTransaction,
		captchaCaptcha, jwtJWT,
		systemMenuService,
		systemConfigService,
		systemGroupDataService,
		mockSystemAdminRepo)
	return adminLoginService
}

func TestAdminLoginService_SystemAdminLogin(t *testing.T) {
	adminLoginService := NewService(t)
	ctx := context.Background()
	req := &validation.SystemAdminLogin{
		Account: "12345678",
		Pwd:     "12345678",
		Key:     "12345678",
		Code:    "12345678",
	}
	ip := "127.0.0.1"

	login, err := adminLoginService.SystemAdminLogin(ctx, req, ip)

	if err != nil {
		t.Error(err)
	}
	t.Log(login)
}
