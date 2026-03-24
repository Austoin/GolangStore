package cart

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandlerListByUserID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := NewMemoryRepository([]Item{{UserID: 7, ProductID: 1, Quantity: 2, Checked: true}})
	handler := NewHandler(NewService(repo))

	router := gin.New()
	router.GET("/carts/:userId", handler.ListByUserID)

	req := httptest.NewRequest(http.MethodGet, "/carts/7", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}

func TestHandlerListByUserIDBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewService(NewMemoryRepository(nil)))

	router := gin.New()
	router.GET("/carts/:userId", handler.ListByUserID)

	req := httptest.NewRequest(http.MethodGet, "/carts/abc", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", resp.Code)
	}
}

func TestHandlerAddItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewService(NewMemoryRepository(nil)))

	router := gin.New()
	router.POST("/carts", handler.AddItem)

	body := []byte(`{"user_id":1,"product_id":11,"product_name":"phone","price":100,"quantity":2,"checked":true}`)
	req := httptest.NewRequest(http.MethodPost, "/carts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.Code)
	}
}
