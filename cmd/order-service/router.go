package main

import (
	"net/http"

	"github.com/austoin/GolangStore/internal/order"
	"github.com/gin-gonic/gin"
)

func newRouter(handler order.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"service": "order-service", "status": "ok"})
	})
	router.GET("/orders", handler.List)
	router.GET("/orders/:orderNo", handler.GetByOrderNo)
	router.POST("/orders", handler.Create)
	router.POST("/orders/from-cart", handler.CreateFromCheckedCartItems)

	return router
}
