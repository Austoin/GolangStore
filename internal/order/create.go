package order

import "github.com/austoin/GolangStore/internal/cart"

func convertCartItems(items []cart.Item) []Item {
	result := make([]Item, 0, len(items))
	for _, item := range items {
		result = append(result, Item{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}

	return result
}
