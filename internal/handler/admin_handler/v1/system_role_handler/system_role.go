package system_role_handler

import (
	"crmeb_go/internal/service/common_service/system_role_service"
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

}

func (h *Handler) Save(ctx *gin.Context) {

}

func (h *Handler) Delete(ctx *gin.Context) {

}

func (h *Handler) Update(ctx *gin.Context) {

}

func (h *Handler) Info(ctx *gin.Context) {

}

func (h *Handler) UpdateStatus(ctx *gin.Context) {

}
