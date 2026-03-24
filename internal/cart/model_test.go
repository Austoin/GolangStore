package cart

import "testing"

func TestItemIsValidQuantity(t *testing.T) {
	item := Item{Quantity: 2}
	if !item.IsValidQuantity() {
		t.Fatal("expected quantity to be valid")
	}
}

func TestItemIsInvalidQuantity(t *testing.T) {
	item := Item{Quantity: 0}
	if item.IsValidQuantity() {
		t.Fatal("expected quantity to be invalid")
	}
}

func TestItemIsChecked(t *testing.T) {
	item := Item{Checked: true}
	if !item.IsChecked() {
		t.Fatal("expected item to be checked")
	}
}
