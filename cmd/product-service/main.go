package main

import (
	"github.com/austoin/GolangStore/internal/product"
	"github.com/austoin/GolangStore/pkg/config"
	projectmysql "github.com/austoin/GolangStore/pkg/mysql"
)

func main() {
	conf := config.Load()
	db, err := projectmysql.Open(conf.MySQL)
	if err != nil {
		panic(err)
	}
	repo := product.NewMySQLRepository(db)
	service := product.NewService(repo)
	handler := product.NewHandler(service)
	router := newRouter(handler)
	_ = router.Run(":8081")
}
