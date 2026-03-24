package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h Handler) GetByOrderNo(ctx *gin.Context) {
	orderNo := ctx.Param("orderNo")
	if orderNo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order no"})
		return
	}

	entity, err := h.service.GetOrder(orderNo)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entity)
}
