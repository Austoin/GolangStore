package product

import (
	"errors"
	"testing"
)

func TestMemoryRepositoryGetByID(t *testing.T) {
	repo := NewMemoryRepository([]Product{{ID: 1, Name: "phone", Price: 1999, Status: 1, Stock: 10}})

	item, err := repo.GetByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if item.Name != "phone" {
		t.Fatalf("expected phone, got %s", item.Name)
	}
}

func TestMemoryRepositoryGetByIDNotFound(t *testing.T) {
	repo := NewMemoryRepository(nil)

	_, err := repo.GetByID(1)
	if !errors.Is(err, ErrProductNotFound) {
		t.Fatalf("expected ErrProductNotFound, got %v", err)
	}
}
