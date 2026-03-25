package product

import "errors"

var ErrProductNotFound = errors.New("product not found")

type Repository interface {
	GetByID(id uint64) (Product, error)
	List() []Product
	Create(product Product) Product
}

type MemoryRepository struct {
	items map[uint64]Product
}

func NewMemoryRepository(items []Product) MemoryRepository {
	indexed := make(map[uint64]Product, len(items))
	for _, item := range items {
		indexed[item.ID] = item
	}

	return MemoryRepository{items: indexed}
}

func (r MemoryRepository) GetByID(id uint64) (Product, error) {
	item, ok := r.items[id]
	if !ok {
		return Product{}, ErrProductNotFound
	}

	return item, nil
}

func (r MemoryRepository) List() []Product {
	items := make([]Product, 0, len(r.items))
	for _, item := range r.items {
		items = append(items, item)
	}

	return items
}

func (r MemoryRepository) Create(product Product) Product {
	product.ID = uint64(len(r.items) + 1)
	r.items[product.ID] = product
	return product
}
