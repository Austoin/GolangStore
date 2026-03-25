package product

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandlerGetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := NewMemoryRepository([]Product{{ID: 3, Name: "mouse", Price: 99, Status: 1, Stock: 8}})
	handler := NewHandler(NewService(repo))

	router := gin.New()
	router.GET("/products/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/products/3", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}

func TestHandlerGetByIDBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewService(NewMemoryRepository(nil)))

	router := gin.New()
	router.GET("/products/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/products/abc", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", resp.Code)
	}
}


func TestHandlerCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := NewMemoryRepository(nil)
	handler := NewHandler(NewService(repo))

	router := gin.New()
	router.POST("/products", handler.Create)

	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(`{"name":"phone","description":"smart phone","price":199900,"status":1,"stock":8}`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.Code)
	}
}
