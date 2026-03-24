package product

import "testing"

func TestProductIsOnSale(t *testing.T) {
	item := Product{Status: 1}
	if !item.IsOnSale() {
		t.Fatal("expected product to be on sale")
	}
}

func TestProductHasStock(t *testing.T) {
	item := Product{Stock: 10}
	if !item.HasStock() {
		t.Fatal("expected product to have stock")
	}
}

func TestProductHasNoStock(t *testing.T) {
	item := Product{Stock: 0}
	if item.HasStock() {
		t.Fatal("expected product to have no stock")
	}
}
