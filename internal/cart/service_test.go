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

func TestServiceListCheckedItems(t *testing.T) {
	repo := NewMemoryRepository([]Item{
		{UserID: 9, ProductID: 1, Quantity: 1, Checked: true},
		{UserID: 9, ProductID: 2, Quantity: 1, Checked: false},
		{UserID: 9, ProductID: 3, Quantity: 2, Checked: true},
	})
	service := NewService(repo)

	items := service.ListCheckedItems(9)
	if len(items) != 2 {
		t.Fatalf("expected 2 checked items, got %d", len(items))
	}

	if items[0].Checked != true || items[1].Checked != true {
		t.Fatal("expected all returned items to be checked")
	}
}

func TestServiceAddItem(t *testing.T) {
	repo := NewMemoryRepository(nil)
	service := NewService(repo)

	item := service.AddItem(Item{UserID: 1, ProductID: 11, ProductName: "phone", Price: 100, Quantity: 2, Checked: true})
	if item.Quantity != 2 {
		t.Fatalf("expected quantity 2, got %d", item.Quantity)
	}
}
