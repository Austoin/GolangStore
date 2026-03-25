package product

type Product struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
	Status      uint8  `json:"status"`
	Stock       int    `json:"stock"`
}

func (p Product) IsOnSale() bool {
	return p.Status == 1
}

func (p Product) HasStock() bool {
	return p.Stock > 0
}
