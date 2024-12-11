package admin_login_handler

import (
	"crmeb_go/internal/service/admin_service/admin_login_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service admin_login_service.Service
}

func New(service admin_login_service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(ctx *gin.Context) {
	var req validation.Login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	token, err := h.service.SystemAdminLogin(ctx, &req, ctx.ClientIP())
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, token)
}
