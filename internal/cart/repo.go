package cart

type Repository interface {
	ListByUserID(userID uint64) []Item
	Save(item Item) Item
	Delete(userID uint64, productID uint64)
	SetChecked(userID uint64, productID uint64, checked bool)
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
	for index, existing := range *r.items {
		if existing.UserID == item.UserID && existing.ProductID == item.ProductID {
			(*r.items)[index] = item
			return item
		}
	}
	*r.items = append(*r.items, item)
	return item
}

func (r MemoryRepository) Delete(userID uint64, productID uint64) {
	filtered := make([]Item, 0, len(*r.items))
	for _, item := range *r.items {
		if !(item.UserID == userID && item.ProductID == productID) {
			filtered = append(filtered, item)
		}
	}
	*r.items = filtered
}

func (r MemoryRepository) SetChecked(userID uint64, productID uint64, checked bool) {
	for index, item := range *r.items {
		if item.UserID == userID && item.ProductID == productID {
			item.Checked = checked
			(*r.items)[index] = item
			return
		}
	}
}
