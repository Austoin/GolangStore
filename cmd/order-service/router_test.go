package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/austoin/GolangStore/internal/cart"
	"github.com/austoin/GolangStore/internal/order"
)

func TestNewRouterHealth(t *testing.T) {
	handler := order.NewHandler(order.NewService(order.NewMemoryRepository(nil), cart.NewService(cart.NewMemoryRepository(nil))))
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}

func TestNewRouterOrderRoutes(t *testing.T) {
	handler := order.NewHandler(order.NewService(order.NewMemoryRepository(nil), cart.NewService(cart.NewMemoryRepository(nil))))
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/orders/O404", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.Code)
	}
}
