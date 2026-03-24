package product

import "testing"

func TestServiceGetProduct(t *testing.T) {
	repo := NewMemoryRepository([]Product{{ID: 2, Name: "keyboard", Price: 299, Status: 1, Stock: 5}})
	service := NewService(repo)

	item, err := service.GetProduct(2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if item.ID != 2 {
		t.Fatalf("expected id 2, got %d", item.ID)
	}
}
