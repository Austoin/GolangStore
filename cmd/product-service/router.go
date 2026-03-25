package main

import (
	"net/http"

	"github.com/austoin/GolangStore/internal/product"
	"github.com/gin-gonic/gin"
)

func newRouter(handler product.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"service": "product-service", "status": "ok"})
	})
	router.GET("/products", handler.List)
	router.GET("/products/:id", handler.GetByID)
	router.POST("/products", handler.Create)

	return router
}
