package order

import "errors"

var ErrOrderNotFound = errors.New("order not found")

type Repository interface {
	GetByOrderNo(orderNo string) (Order, error)
	List() []Order
	Create(order Order) (Order, error)
}

type MemoryRepository struct {
	items map[string]Order
}

func NewMemoryRepository(items []Order) MemoryRepository {
	indexed := make(map[string]Order, len(items))
	for _, item := range items {
		indexed[item.OrderNo] = item
	}

	return MemoryRepository{items: indexed}
}

func (r MemoryRepository) GetByOrderNo(orderNo string) (Order, error) {
	item, ok := r.items[orderNo]
	if !ok {
		return Order{}, ErrOrderNotFound
	}

	return item, nil
}

func (r MemoryRepository) Create(order Order) (Order, error) {
	r.items[order.OrderNo] = order
	return order, nil
}

func (r MemoryRepository) List() []Order {
	items := make([]Order, 0, len(r.items))
	for _, item := range r.items {
		items = append(items, item)
	}

	return items
}
