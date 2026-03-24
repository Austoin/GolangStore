package main

import (
	"github.com/austoin/GolangStore/internal/cart"
	"github.com/austoin/GolangStore/internal/order"
	"github.com/austoin/GolangStore/internal/product"
	"github.com/austoin/GolangStore/pkg/config"
	projectmysql "github.com/austoin/GolangStore/pkg/mysql"
	"gorm.io/gorm"
)

func main() {
	conf := config.Load()
	db, err := projectmysql.Open(conf.MySQL)
	if err != nil {
		panic(err)
	}
	handler := buildHandler(db)
	router := newRouter(handler)
	_ = router.Run(":8082")
}

func buildHandler(db *gorm.DB) order.Handler {
	orderRepo := order.NewMySQLRepository(db)
	cartRepo := cart.NewMySQLRepository(db)
	productRepo := product.NewMySQLRepository(db)
	cartService := cart.NewService(cartRepo)
	service := order.NewService(orderRepo, cartService, productRepo)
	return order.NewHandler(service)
}

func buildRuntimeDependencies(conf config.Config) string {
	return projectmysql.BuildDSN(conf.MySQL)
}
