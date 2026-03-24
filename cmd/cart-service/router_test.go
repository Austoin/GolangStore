package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/austoin/GolangStore/internal/cart"
)

func TestNewRouterHealth(t *testing.T) {
	handler := cart.NewHandler(cart.NewService(cart.NewMemoryRepository(nil)))
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}

func TestNewRouterCartRoutes(t *testing.T) {
	handler := cart.NewHandler(cart.NewService(cart.NewMemoryRepository(nil)))
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/carts/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}
