package system_store_staff_handler

import (
	"crmeb_go/internal/service/common_service/system_store_staff_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service system_store_staff_service.Service
}

func New(service system_store_staff_service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetList(ctx *gin.Context) {
	req := new(validation.GetSystemStoreStaffList)
	if err := ctx.ShouldBindQuery(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	resp, err := h.service.GetStaffList(ctx, req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}
