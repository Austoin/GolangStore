package cart

type Repository interface {
	ListByUserID(userID uint64) []Item
	Save(item Item) Item
}

type MemoryRepository struct {
	items *[]Item
}

func NewMemoryRepository(items []Item) MemoryRepository {
	cloned := append([]Item(nil), items...)
	return MemoryRepository{items: &cloned}
}

func (r MemoryRepository) ListByUserID(userID uint64) []Item {
	result := make([]Item, 0)
	for _, item := range *r.items {
		if item.UserID == userID {
			result = append(result, item)
		}
	}

	return result
}

func (r MemoryRepository) Save(item Item) Item {
	*r.items = append(*r.items, item)
	return item
}
