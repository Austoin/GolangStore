package cart

import (
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
