package cart

type Item struct {
	UserID    uint64
	ProductID uint64
	Quantity  int
	Checked   bool
}

func (i Item) IsValidQuantity() bool {
	return i.Quantity > 0
}

func (i Item) IsChecked() bool {
	return i.Checked
}
