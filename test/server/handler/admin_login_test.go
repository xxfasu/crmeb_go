package handler

import (
	"bytes"
	"crmeb_go/internal/common/response"
	"crmeb_go/internal/handler/admin_handler/v1/admin_login_handler"
	"crmeb_go/internal/validation"
	res "crmeb_go/pkg/response"
	"crmeb_go/test/mocks/service/admin_service/mocks_admin_login_service"
	"encoding/json"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdminLogin_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// authM := middleware.NewAuthM(jwtJWT)
	// casbinService := mocks_casbin.NewMockService(ctrl)
	// casbinM := middleware.NewCasbinM(casbinService, jwtJWT)
	adminLoginService := mocks_admin_login_service.NewMockService(ctrl)
	handler := admin_login_handler.New(adminLoginService)
	router.POST("/login", handler.Login)

	adminLoginService.EXPECT().SystemAdminLogin(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&response.SystemLogin{
			ID:       1,
			Account:  "1",
			RealName: "1",
			Token:    "1",
			IsSMS:    false,
		}, nil)

	// 4. 准备请求体（JSON）
	requestBody := &validation.SystemAdminLogin{
		Account: "1111",
		Pwd:     "123456",
		Key:     "1",
		Code:    "1",
	}
	jsonBody, err := json.Marshal(requestBody)

	// 5. 构造请求
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("构造请求失败: %v", err)
	}
	// 注意设置 Content-Type，才能被 ctx.ShouldBindJSON 解析
	req.Header.Set("Content-Type", "application/json")

	// 6. 创建响应记录器
	w := httptest.NewRecorder()

	// 7. 让路由执行这个请求
	router.ServeHTTP(w, req)

	resp := new(res.Response)
	json.Unmarshal(w.Body.Bytes(), resp)
	respData := new(response.SystemLogin)
	dataJson, err := json.Marshal(resp.Data)
	json.Unmarshal(dataJson, respData)
	t.Logf("%#v", respData)
}
