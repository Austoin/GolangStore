package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/austoin/GolangStore/internal/product"
)

func TestNewRouterHealth(t *testing.T) {
	handler := product.NewHandler(product.NewService(product.NewMemoryRepository(nil)))
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}

func TestNewRouterProductRoutes(t *testing.T) {
	handler := product.NewHandler(product.NewService(product.NewMemoryRepository(nil)))
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.Code)
	}
}


func TestNewRouterProductListRoute(t *testing.T) {
	handler := product.NewHandler(product.NewService(product.NewMemoryRepository(nil)))
	router := newRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}
