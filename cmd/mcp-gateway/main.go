package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"service": "mcp-gateway", "status": "ok"})
	})
	_ = router.Run(":8085")
}
