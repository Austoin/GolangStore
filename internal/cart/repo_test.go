package cart

import "testing"

func TestMemoryRepositoryListByUserID(t *testing.T) {
	repo := NewMemoryRepository([]Item{
		{UserID: 1, ProductID: 11, Quantity: 1, Checked: true},
		{UserID: 2, ProductID: 12, Quantity: 2, Checked: true},
		{UserID: 1, ProductID: 13, Quantity: 3, Checked: false},
	})

	items := repo.ListByUserID(1)
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
}
