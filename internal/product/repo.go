package product

import "errors"

var ErrProductNotFound = errors.New("product not found")

type Repository interface {
	GetByID(id uint64) (Product, error)
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
