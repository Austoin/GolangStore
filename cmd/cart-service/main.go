package main

import (
	"github.com/austoin/GolangStore/internal/cart"
	"github.com/austoin/GolangStore/pkg/config"
	projectmysql "github.com/austoin/GolangStore/pkg/mysql"
)

func main() {
	conf := config.Load()
	db, err := projectmysql.Open(conf.MySQL)
	if err != nil {
		panic(err)
	}
	repo := cart.NewMySQLRepository(db)
	service := cart.NewService(repo)
	handler := cart.NewHandler(service)
	router := newRouter(handler)
	_ = router.Run(":8083")
}
