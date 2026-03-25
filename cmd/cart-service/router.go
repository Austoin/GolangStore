package main

import (
	"net/http"

	"github.com/austoin/GolangStore/internal/cart"
	"github.com/gin-gonic/gin"
)

func newRouter(handler cart.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"service": "cart-service", "status": "ok"})
	})
	router.GET("/carts/:userId", handler.ListByUserID)
	router.POST("/carts", handler.AddItem)
	router.DELETE("/carts/:userId/:productId", handler.DeleteItem)
	router.PATCH("/carts/:userId/:productId/checked", handler.SetChecked)

	return router
}
