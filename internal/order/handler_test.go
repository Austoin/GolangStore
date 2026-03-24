package order

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandlerGetByOrderNo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := NewMemoryRepository([]Order{{OrderNo: "O2026003", UserID: 3, Status: StatusPending}})
	handler := NewHandler(NewService(repo))

	router := gin.New()
	router.GET("/orders/:orderNo", handler.GetByOrderNo)

	req := httptest.NewRequest(http.MethodGet, "/orders/O2026003", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}

func TestHandlerGetByOrderNoNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewService(NewMemoryRepository(nil)))

	router := gin.New()
	router.GET("/orders/:orderNo", handler.GetByOrderNo)

	req := httptest.NewRequest(http.MethodGet, "/orders/O404", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.Code)
	}
}
