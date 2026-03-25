package cart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

type updateCheckedRequest struct {
	Checked bool `json:"checked"`
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

func (h Handler) AddItem(ctx *gin.Context) {
	var item Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	saved := h.service.AddItem(item)
	ctx.JSON(http.StatusCreated, saved)
}

func (h Handler) DeleteItem(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	productID, err := strconv.ParseUint(ctx.Param("productId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	
	h.service.DeleteItem(userID, productID)
	ctx.Status(http.StatusNoContent)
}

func (h Handler) SetChecked(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	productID, err := strconv.ParseUint(ctx.Param("productId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var req updateCheckedRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	h.service.SetItemChecked(userID, productID, req.Checked)
	ctx.Status(http.StatusNoContent)
}
