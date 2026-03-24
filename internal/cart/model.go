package cart

type Item struct {
	UserID      uint64 `json:"user_id"`
	ProductID   uint64 `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       uint64 `json:"price"`
	Quantity    int    `json:"quantity"`
	Checked     bool   `json:"checked"`
}

func (i Item) IsValidQuantity() bool {
	return i.Quantity > 0
}

func (i Item) IsChecked() bool {
	return i.Checked
}
