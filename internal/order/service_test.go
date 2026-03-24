package order

import "testing"

func TestServiceGetOrder(t *testing.T) {
	repo := NewMemoryRepository([]Order{{OrderNo: "O2026002", UserID: 2, Status: StatusPaid}})
	service := NewService(repo)

	entity, err := service.GetOrder("O2026002")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if entity.Status != StatusPaid {
		t.Fatalf("expected status paid, got %d", entity.Status)
	}
}
