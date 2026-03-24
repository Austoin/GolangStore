package mysql

import (
	"testing"

	"github.com/austoin/GolangStore/pkg/config"
)

func TestBuildDSN(t *testing.T) {
	conf := config.MySQL{
		Host:     "mysql",
		Port:     "3306",
		User:     "root",
		Password: "root",
		Database: "golang_store",
	}

	dsn := BuildDSN(conf)
	expected := "root:root@tcp(mysql:3306)/golang_store?charset=utf8mb4&parseTime=True&loc=Local"

	if dsn != expected {
		t.Fatalf("expected %s, got %s", expected, dsn)
	}
}
