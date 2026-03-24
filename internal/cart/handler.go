package cart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h Handler) ListByUserID(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	items := h.service.ListItems(userID)
	ctx.JSON(http.StatusOK, items)
}
