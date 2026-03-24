package order

const (
	StatusPending uint8 = 1
	StatusPaid    uint8 = 2
	StatusClosed  uint8 = 3
)

type Item struct {
	ProductID   uint64
	ProductName string
	Price       uint64
	Quantity    int
}

type Order struct {
	ID          uint64
	OrderNo     string
	UserID      uint64
	TotalAmount uint64
	Status      uint8
	Items       []Item
}

func (o Order) IsPending() bool {
	return o.Status == StatusPending
}

func (o Order) IsPaid() bool {
	return o.Status == StatusPaid
}

func (o Order) HasItems() bool {
	return len(o.Items) > 0
}
