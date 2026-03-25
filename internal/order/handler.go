package order

import (
	"net/http"

	"github.com/austoin/GolangStore/internal/cart"
	"github.com/gin-gonic/gin"
)

type CreateFromCartRequest struct {
	UserID uint64      `json:"user_id"`
	Items  []cart.Item `json:"items"`
}

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

func (h Handler) Create(ctx *gin.Context) {
	var req CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if len(req.Items) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "items cannot be empty"})
		return
	}

	entity, err := h.service.CreateOrder(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, entity)
}

func (h Handler) CreateFromCheckedCartItems(ctx *gin.Context) {
	var req CreateFromCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	entity, err := h.service.CreateOrderFromCheckedCartItems(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, entity)
}

func (h Handler) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.service.ListOrders())
}
