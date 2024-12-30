package system_admin_handler

import (
	"crmeb_go/internal/service/common_service/system_admin_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	service system_admin_service.Service
}

func New(service system_admin_service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(ctx *gin.Context) {
	req := new(validation.SystemAdminSearch)
	if err := ctx.ShouldBindJSON(req); err != nil {
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
	req := new(validation.SystemAdminAdd)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := h.service.Save(ctx, req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (h *Handler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	err = h.service.Delete(ctx, int64(id))
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (h *Handler) Update(ctx *gin.Context) {
	req := new(validation.SystemAdminSearch)
	if err := ctx.ShouldBindJSON(req); err != nil {
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

func (h *Handler) Info(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	resp, err := h.service.Info(ctx, int64(id))
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (h *Handler) UpdateStatus(ctx *gin.Context) {
	req := new(validation.SystemAdminSearch)
	if err := ctx.ShouldBindJSON(req); err != nil {
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

func (h *Handler) UpdateSms(ctx *gin.Context) {
	req := new(validation.SystemAdminSearch)
	if err := ctx.ShouldBindJSON(req); err != nil {
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
