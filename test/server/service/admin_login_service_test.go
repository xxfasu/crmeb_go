package service

import (
	"context"
	"crmeb_go/internal/common/data"
	"crmeb_go/internal/conf"
	"crmeb_go/internal/model"
	redis2 "crmeb_go/internal/redis"
	"crmeb_go/internal/service/admin_service/admin_login_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/cache"
	"crmeb_go/pkg/jwt"
	"crmeb_go/pkg/logs"
	"crmeb_go/test/mocks/pkg/mocks_captcha"
	"crmeb_go/test/mocks/repository/mocks_system_admin_repository"
	"crmeb_go/test/mocks/repository/mocks_transaction"
	"crmeb_go/test/mocks/service/common_service/mocks_system_config_service"
	"crmeb_go/test/mocks/service/common_service/mocks_system_group_data_service"
	"crmeb_go/test/mocks/service/common_service/mocks_system_menu_service"
	"fmt"
	"github.com/redis/go-redis/v9"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

var jwtJWT *jwt.JWT
var mockCaptcha *mocks_captcha.MockCaptcha
var rClient redis.UniversalClient
var mockSystemAdminRepo *mocks_system_admin_repository.MockRepository
var mockTransaction *mocks_transaction.MockTransaction
var systemMenuService *mocks_system_menu_service.MockService
var systemConfigService *mocks_system_config_service.MockService
var systemGroupDataService *mocks_system_group_data_service.MockService

func TestMain(m *testing.M) {
	fmt.Println("begin")
	var err error
	conf.InitConfig("D:\\goproject\\xxfasu\\crmeb_go\\config")
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

	code := m.Run()
	fmt.Println("test end")
	os.Exit(code)
}

func NewService(t *testing.T) admin_login_service.Service {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCaptcha = mocks_captcha.NewMockCaptcha(ctrl)
	mockSystemAdminRepo = mocks_system_admin_repository.NewMockRepository(ctrl)
	mockTransaction = mocks_transaction.NewMockTransaction(ctrl)
	systemMenuService = mocks_system_menu_service.NewMockService(ctrl)
	systemConfigService = mocks_system_config_service.NewMockService(ctrl)
	systemGroupDataService = mocks_system_group_data_service.NewMockService(ctrl)
	adminLoginService := admin_login_service.New(mockTransaction,
		mockCaptcha, jwtJWT,
		systemMenuService,
		systemConfigService,
		systemGroupDataService,
		mockSystemAdminRepo)
	return adminLoginService
}

func TestAdminLoginService_SystemAdminLogin(t *testing.T) {
	Convey("AdminLoginService SystemAdminLogin方法测试", t, func() {
		adminLoginService := NewService(t)
		ctx := context.Background()

		// 模拟一个登录请求
		req := &validation.SystemAdminLogin{
			Account: "123456",
			Pwd:     "123456",
			Key:     "",
			Code:    "",
		}
		ip := "127.0.0.1"

		// mockCaptcha 验证成功
		mockCaptcha.EXPECT().Verify(gomock.Any(), gomock.Any()).Return(true).AnyTimes()

		// mockSystemAdminRepo 模拟查询到用户信息
		mockSystemAdminRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&model.SystemAdmin{
			ID:    1,
			Pwd:   "$2a$10$XoY938kDItzKPVQRF9PWWufqRjz289xu7jAOZKqrgAJPGs6tbE1YC",
			Roles: "1",
		}, nil).AnyTimes()

		// mockSystemAdminRepo 模拟更新
		mockSystemAdminRepo.EXPECT().UpdateFields(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

		// systemMenuService 模拟权限查询
		systemMenuService.EXPECT().GetAllPermissions(gomock.Any()).Return([]*model.SystemMenu{}, nil).AnyTimes()
		login, err := adminLoginService.SystemAdminLogin(ctx, req, ip)
		Printf("结果为:%#v", *login)
		So(err, ShouldBeNil)
		So(login, ShouldNotBeNil)
	})
}

func TestAdminLoginService_GetCode(t *testing.T) {
	Convey("AdminLoginService GetCode方法测试", t, func() {
		adminLoginService := NewService(t)
		ctx := context.Background()

		// mockCaptcha 验证成功
		mockCaptcha.EXPECT().Gen().Return("key", "code", nil).AnyTimes()

		resp, err := adminLoginService.GetCode(ctx)
		Printf("结果为:%#v", *resp)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
	})
}

func TestAdminLoginService_GetAdminInfo(t *testing.T) {
	Convey("AdminLoginServiceGetAdminInfo方法测试", t, func() {
		adminLoginService := NewService(t)
		ctx := context.Background()

		req := data.LoginUser{
			Token:      "7cfcdd35-138d-4752-8e92-885dc837a758",
			LoginTime:  1735818500,
			ExpireTime: 1735836500,
			User:       &model.SystemAdmin{ID: 1, Account: "admin", Pwd: "123456", RealName: "超管", Roles: "1", LastIP: "127.0.0.1", LoginCount: 516, Level: 1, Status: 1, Phone: "11111111111", IsSms: 0, CreatedAt: 1734768311, UpdatedAt: 1735809086, DeletedAt: 0},
		}
		resp, err := adminLoginService.GetAdminInfo(ctx, req)
		Printf("结果为:%#v", *resp)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
	})
}
