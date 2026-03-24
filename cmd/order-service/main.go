package main

import (
	"github.com/austoin/GolangStore/internal/cart"
	"github.com/austoin/GolangStore/internal/order"
)

func main() {
	repo := order.NewMemoryRepository(nil)
	cartRepo := cart.NewMemoryRepository(nil)
	cartService := cart.NewService(cartRepo)
	service := order.NewService(repo, cartService)
	handler := order.NewHandler(service)
	router := newRouter(handler)
	_ = router.Run(":8082")
}
