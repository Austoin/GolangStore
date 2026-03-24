package product

type Product struct {
	ID          uint64
	Name        string
	Description string
	Price       uint64
	Status      uint8
	Stock       int
}

func (p Product) IsOnSale() bool {
	return p.Status == 1
}

func (p Product) HasStock() bool {
	return p.Stock > 0
}
