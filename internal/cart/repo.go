package cart

type Repository interface {
	ListByUserID(userID uint64) []Item
}

type MemoryRepository struct {
	items []Item
}

func NewMemoryRepository(items []Item) MemoryRepository {
	return MemoryRepository{items: items}
}

func (r MemoryRepository) ListByUserID(userID uint64) []Item {
	result := make([]Item, 0)
	for _, item := range r.items {
		if item.UserID == userID {
			result = append(result, item)
		}
	}

	return result
}
