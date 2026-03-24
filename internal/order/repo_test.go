package order

import (
	"errors"
	"testing"
)

func TestMemoryRepositoryGetByOrderNo(t *testing.T) {
	repo := NewMemoryRepository([]Order{{OrderNo: "O2026001", UserID: 1, Status: StatusPending}})

	entity, err := repo.GetByOrderNo("O2026001")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if entity.UserID != 1 {
		t.Fatalf("expected user id 1, got %d", entity.UserID)
	}
}

func TestMemoryRepositoryGetByOrderNoNotFound(t *testing.T) {
	repo := NewMemoryRepository(nil)

	_, err := repo.GetByOrderNo("missing")
	if !errors.Is(err, ErrOrderNotFound) {
		t.Fatalf("expected ErrOrderNotFound, got %v", err)
	}
}

func TestMemoryRepositoryCreate(t *testing.T) {
	repo := NewMemoryRepository(nil)
	entity := Order{OrderNo: "O2027001", UserID: 7, Status: StatusPending}

	created, err := repo.Create(entity)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	loaded, err := repo.GetByOrderNo("O2027001")
	if err != nil {
		t.Fatalf("expected order to be stored, got %v", err)
	}

	if created.OrderNo != loaded.OrderNo {
		t.Fatalf("expected created and loaded order no to match")
	}
}
