package system_role_handler

import (
	"crmeb_go/internal/service/common_service/system_role_service"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
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
	req := new(validation.SystemRole)
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
	req := new(validation.SystemRole)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	err := h.service.Update(ctx, req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
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
	idStr := ctx.Param("id")
	statusStr := ctx.Param("status")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	err = h.service.UpdateStatus(ctx, int64(id), int64(status))
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	response.Ok(ctx)
}
