package order

import "testing"

func TestOrderIsPending(t *testing.T) {
	entity := Order{Status: StatusPending}
	if !entity.IsPending() {
		t.Fatal("expected order to be pending")
	}
}

func TestOrderIsPaid(t *testing.T) {
	entity := Order{Status: StatusPaid}
	if !entity.IsPaid() {
		t.Fatal("expected order to be paid")
	}
}

func TestOrderHasItems(t *testing.T) {
	entity := Order{Items: []Item{{ProductID: 1, Quantity: 1}}}
	if !entity.HasItems() {
		t.Fatal("expected order to have items")
	}
}

func TestOrderHasNoItems(t *testing.T) {
	entity := Order{}
	if entity.HasItems() {
		t.Fatal("expected order to have no items")
	}
}

func TestCreateRequestCarriesUserAndItems(t *testing.T) {
	req := CreateRequest{
		UserID: 12,
		Items: []CreateItemRequest{{
			ProductID:   101,
			ProductName: "phone",
			Price:       199900,
			Quantity:    1,
		}},
	}

	if req.UserID != 12 {
		t.Fatalf("expected user id 12, got %d", req.UserID)
	}

	if len(req.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(req.Items))
	}

	if req.Items[0].ProductID != 101 {
		t.Fatalf("expected product id 101, got %d", req.Items[0].ProductID)
	}
}

func TestOrderCalculateTotalAmount(t *testing.T) {
	entity := Order{
		Items: []Item{
			{ProductID: 1, Price: 100, Quantity: 2},
			{ProductID: 2, Price: 350, Quantity: 1},
		},
	}

	if entity.CalculateTotalAmount() != 550 {
		t.Fatalf("expected total amount 550, got %d", entity.CalculateTotalAmount())
	}
}
