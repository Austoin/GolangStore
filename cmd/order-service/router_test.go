package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/austoin/GolangStore/internal/cart"
	"github.com/austoin/GolangStore/internal/order"
	"github.com/austoin/GolangStore/pkg/config"
	projectmysql "github.com/austoin/GolangStore/pkg/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewRouterHealth(t *testing.T) {
	handler := order.NewHandler(order.NewService(order.NewMemoryRepository(nil), cart.NewService(cart.NewMemoryRepository(nil)), nil))
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}

func TestNewRouterOrderRoutes(t *testing.T) {
	handler := order.NewHandler(order.NewService(order.NewMemoryRepository(nil), cart.NewService(cart.NewMemoryRepository(nil)), nil))
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/orders/O404", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.Code)
	}
}

func TestBuildRuntimeDependencies(t *testing.T) {
	conf := config.Config{
		MySQL: config.MySQL{
			Host:     "127.0.0.1",
			Port:     "3306",
			User:     "root",
			Password: "root",
			Database: "golang_store",
		},
	}

	build := func() string {
		return buildRuntimeDependencies(conf)
	}

	if build() != projectmysql.BuildDSN(conf.MySQL) {
		t.Fatalf("expected mysql dsn %s, got %s", projectmysql.BuildDSN(conf.MySQL), build())
	}
}

func TestBuildHandler(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("expected sqlite db, got %v", err)
	}

	handler := buildHandler(db)
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}
