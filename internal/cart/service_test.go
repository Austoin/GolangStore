package cart

import "testing"

func TestServiceListItems(t *testing.T) {
	repo := NewMemoryRepository([]Item{{UserID: 9, ProductID: 1, Quantity: 1, Checked: true}})
	service := NewService(repo)

	items := service.ListItems(9)
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
}
