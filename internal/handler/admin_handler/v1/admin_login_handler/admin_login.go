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

func (h *Handler) GetCode(ctx *gin.Context) {
	resp, err := h.service.GetCode(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (h *Handler) Login(ctx *gin.Context) {
	var req validation.SystemAdminLogin
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	resp, err := h.service.SystemAdminLogin(ctx, &req, ctx.ClientIP())
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}
