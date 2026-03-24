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
