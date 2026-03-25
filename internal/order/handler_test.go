package order

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/austoin/GolangStore/internal/cart"
	"github.com/gin-gonic/gin"
)

func TestHandlerGetByOrderNo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := NewMemoryRepository([]Order{{OrderNo: "O2026003", UserID: 3, Status: StatusPending}})
	handler := NewHandler(NewService(repo, fakeCartQuery{}, nil))

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
	handler := NewHandler(NewService(NewMemoryRepository(nil), fakeCartQuery{}, nil))

	router := gin.New()
	router.GET("/orders/:orderNo", handler.GetByOrderNo)

	req := httptest.NewRequest(http.MethodGet, "/orders/O404", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", resp.Code)
	}
}

func TestHandlerCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewService(NewMemoryRepository(nil), fakeCartQuery{}, nil))

	router := gin.New()
	router.POST("/orders", handler.Create)

	body := []byte(`{"user_id":1,"items":[{"product_id":1,"product_name":"phone","price":100,"quantity":2}]}`)
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.Code)
	}
}

func TestHandlerCreateBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewService(NewMemoryRepository(nil), fakeCartQuery{}, nil))

	router := gin.New()
	router.POST("/orders", handler.Create)

	body := []byte(`{"user_id":1,"items":[]}`)
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", resp.Code)
	}
}

func TestHandlerCreateFromCheckedCartItems(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewService(NewMemoryRepository(nil), fakeCartQuery{items: []cart.Item{{UserID: 1, ProductID: 1, ProductName: "phone", Price: 100, Quantity: 2, Checked: true}}}, nil))

	router := gin.New()
	router.POST("/orders/from-cart", handler.CreateFromCheckedCartItems)

	body := []byte(`{"user_id":1}`)
	req := httptest.NewRequest(http.MethodPost, "/orders/from-cart", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.Code)
	}
}

func TestHandlerCreateFromCheckedCartItemsBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewService(NewMemoryRepository(nil), fakeCartQuery{}, nil))

	router := gin.New()
	router.POST("/orders/from-cart", handler.CreateFromCheckedCartItems)

	body := []byte(`{"user_id":1}`)
	req := httptest.NewRequest(http.MethodPost, "/orders/from-cart", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", resp.Code)
	}
}


func TestHandlerList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewService(NewMemoryRepository([]Order{{OrderNo: "O1", UserID: 1, Status: StatusPending}}), fakeCartQuery{}, nil))

	router := gin.New()
	router.GET("/orders", handler.List)

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.Code)
	}
}
