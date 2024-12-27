package system_menu_handler

import (
	"crmeb_go/internal/service/common_service/system_menu_service"
	"crmeb_go/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service system_menu_service.Service
}

func New(service system_menu_service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CacheTree(ctx *gin.Context) {
	resp, err := h.service.GetCacheTree(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}
