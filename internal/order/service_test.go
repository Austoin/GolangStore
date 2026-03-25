package order

import (
	"testing"

	"github.com/austoin/GolangStore/internal/cart"
)

type fakeCartQuery struct {
	items []cart.Item
}

func (f fakeCartQuery) ListCheckedItems(userID uint64) []cart.Item {
	return f.items
}

type fakeStockStore struct {
	stock map[uint64]int
}

func (f *fakeStockStore) HasEnough(productID uint64, quantity int) bool {
	return f.stock[productID] >= quantity
}

func (f *fakeStockStore) Deduct(productID uint64, quantity int) {
	f.stock[productID] -= quantity
}

func TestServiceGetOrder(t *testing.T) {
	repo := NewMemoryRepository([]Order{{OrderNo: "O2026002", UserID: 2, Status: StatusPaid}})
	service := NewService(repo, fakeCartQuery{}, nil)

	entity, err := service.GetOrder("O2026002")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if entity.Status != StatusPaid {
		t.Fatalf("expected status paid, got %d", entity.Status)
	}
}

func TestConvertCartItemsToOrderItems(t *testing.T) {
	service := NewService(NewMemoryRepository(nil), fakeCartQuery{}, nil)

	items := service.ConvertCartItems([]cart.Item{{
		ProductID:   9,
		ProductName: "laptop",
		Price:       599900,
		Quantity:    1,
	}})

	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	if items[0].ProductName != "laptop" {
		t.Fatalf("expected laptop, got %s", items[0].ProductName)
	}
}

func TestServiceCreateOrder(t *testing.T) {
	service := NewService(NewMemoryRepository(nil), fakeCartQuery{}, nil)

	entity, err := service.CreateOrder(CreateRequest{
		UserID: 3,
		Items: []CreateItemRequest{{
			ProductID:   1,
			ProductName: "phone",
			Price:       100,
			Quantity:    2,
		}},
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if entity.Status != StatusPending {
		t.Fatalf("expected pending status, got %d", entity.Status)
	}

	if entity.TotalAmount != 200 {
		t.Fatalf("expected total amount 200, got %d", entity.TotalAmount)
	}

	if entity.OrderNo == "" {
		t.Fatal("expected generated order no")
	}
}

func TestServiceCreateOrderFromCheckedCartItems(t *testing.T) {
	service := NewService(NewMemoryRepository(nil), fakeCartQuery{items: []cart.Item{
		{UserID: 3, ProductID: 1, ProductName: "phone", Price: 100, Quantity: 2, Checked: true},
		{UserID: 3, ProductID: 2, ProductName: "cable", Price: 50, Quantity: 1, Checked: true},
	}}, nil)

	entity, err := service.CreateOrderFromCheckedCartItems(3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if entity.TotalAmount != 250 {
		t.Fatalf("expected total amount 250, got %d", entity.TotalAmount)
	}

	if entity.Status != StatusPending {
		t.Fatalf("expected pending status, got %d", entity.Status)
	}
}

func TestServiceCreateOrderFromCheckedCartItemsEmpty(t *testing.T) {
	service := NewService(NewMemoryRepository(nil), fakeCartQuery{}, nil)

	_, err := service.CreateOrderFromCheckedCartItems(3)
	if err == nil {
		t.Fatal("expected error for empty checked cart items")
	}
}

func TestServiceCreateOrderFromCheckedCartItemsDeductsStock(t *testing.T) {
	stocks := &fakeStockStore{stock: map[uint64]int{1: 5}}
	service := NewService(NewMemoryRepository(nil), fakeCartQuery{items: []cart.Item{{
		UserID: 3, ProductID: 1, ProductName: "phone", Price: 100, Quantity: 2, Checked: true,
	}}}, stocks)

	_, err := service.CreateOrderFromCheckedCartItems(3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if stocks.stock[1] != 3 {
		t.Fatalf("expected stock 3, got %d", stocks.stock[1])
	}
}

func TestServiceCreateOrderFromCheckedCartItemsRejectsInsufficientStock(t *testing.T) {
	stocks := &fakeStockStore{stock: map[uint64]int{1: 1}}
	service := NewService(NewMemoryRepository(nil), fakeCartQuery{items: []cart.Item{{
		UserID: 3, ProductID: 1, ProductName: "phone", Price: 100, Quantity: 2, Checked: true,
	}}}, stocks)

	_, err := service.CreateOrderFromCheckedCartItems(3)
	if err == nil {
		t.Fatal("expected insufficient stock error")
	}
}


func TestServiceListOrders(t *testing.T) {
	service := NewService(NewMemoryRepository([]Order{{OrderNo: "O1", UserID: 1, Status: StatusPending}}), fakeCartQuery{}, nil)
	items := service.ListOrders()
	if len(items) != 1 {
		t.Fatalf("expected 1 order, got %d", len(items))
	}
}
