package system_config_handler

import (
	"crmeb_go/constants"
	"crmeb_go/internal/service/common_service/system_config_service"
	"crmeb_go/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service system_config_service.Service
}

func New(service system_config_service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetJsConfig(ctx *gin.Context) {
	resp := h.service.GetValueByKey(ctx, constants.JSConfigCrmebChatStatistics)
	response.OkWithData(ctx, resp)
}
