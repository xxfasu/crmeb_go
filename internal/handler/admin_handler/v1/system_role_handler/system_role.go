package system_role_handler

import (
	"crmeb_go/internal/service/common_service/system_role_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service system_role_service.Service
}

func New(service system_role_service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(ctx *gin.Context) {
	req := new(validation.SystemRoleSearch)
	if err := ctx.ShouldBindQuery(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	resp, err := h.service.List(ctx, req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (h *Handler) Save(ctx *gin.Context) {

}

func (h *Handler) Delete(ctx *gin.Context) {

}

func (h *Handler) Update(ctx *gin.Context) {

}

func (h *Handler) Info(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := h.service.Info(ctx, id)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (h *Handler) UpdateStatus(ctx *gin.Context) {

}
