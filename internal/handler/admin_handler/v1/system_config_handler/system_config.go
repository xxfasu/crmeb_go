package system_config_handler

import (
	"crmeb_go/constants"
	"crmeb_go/internal/service/common_service/system_config_service"
	"crmeb_go/internal/validation"
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

func (h *Handler) GetUniq(ctx *gin.Context) {
	req := new(validation.GetUniq)
	if err := ctx.ShouldBindQuery(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	resp := h.service.GetValueByKey(ctx, req.Key)
	response.OkWithData(ctx, resp)
}
