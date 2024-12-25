package home_handler

import (
	"crmeb_go/internal/service/common_service/home_service"
	"crmeb_go/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service home_service.Service
}

func New(service home_service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) IndexDate(ctx *gin.Context) {
	resp, err := h.service.IndexDate(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (h *Handler) ChartUser(ctx *gin.Context) {
	resp, err := h.service.ChartUser(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (h *Handler) ChartOrder(ctx *gin.Context) {
	resp, err := h.service.ChartOrder(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (h *Handler) ChartOrderInWeek(ctx *gin.Context) {
	resp, err := h.service.ChartOrderWeek(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (h *Handler) ChartOrderInMonth(ctx *gin.Context) {
	resp, err := h.service.ChartOrderMonth(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (h *Handler) ChartOrderInYear(ctx *gin.Context) {
	resp, err := h.service.ChartOrderYear(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}
