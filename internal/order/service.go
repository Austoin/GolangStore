package order

import (
	"errors"
	"fmt"

	"github.com/austoin/GolangStore/internal/cart"
)

var ErrEmptyCheckedCartItems = errors.New("checked cart items cannot be empty")

type Service struct {
	repo  Repository
	carts cartQuery
}

type cartQuery interface {
	ListCheckedItems(userID uint64) []cart.Item
}

func NewService(repo Repository, carts cartQuery) Service {
	return Service{repo: repo, carts: carts}
}

func (s Service) GetOrder(orderNo string) (Order, error) {
	return s.repo.GetByOrderNo(orderNo)
}

func (s Service) ConvertCartItems(items []cart.Item) []Item {
	return convertCartItems(items)
}

func (s Service) CreateOrder(req CreateRequest) (Order, error) {
	items := make([]Item, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, Item{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}

	entity := Order{
		OrderNo: fmt.Sprintf("O%d", req.UserID*1000+uint64(len(req.Items))),
		UserID:  req.UserID,
		Status:  StatusPending,
		Items:   items,
	}
	entity.TotalAmount = entity.CalculateTotalAmount()

	return s.repo.Create(entity)
}

func (s Service) CreateOrderFromCheckedCartItems(userID uint64) (Order, error) {
	checkedItems := s.carts.ListCheckedItems(userID)
	if len(checkedItems) == 0 {
		return Order{}, ErrEmptyCheckedCartItems
	}

	items := s.ConvertCartItems(checkedItems)
	entity := Order{
		OrderNo: fmt.Sprintf("O%d", userID*1000+uint64(len(items))),
		UserID:  userID,
		Status:  StatusPending,
		Items:   items,
	}
	entity.TotalAmount = entity.CalculateTotalAmount()

	return s.repo.Create(entity)
}
