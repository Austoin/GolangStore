package order

import (
	"errors"
	"fmt"

	"github.com/austoin/GolangStore/internal/cart"
)

var ErrEmptyCheckedCartItems = errors.New("checked cart items cannot be empty")
var ErrInsufficientStock = errors.New("insufficient stock")

type Service struct {
	repo   Repository
	carts  cartQuery
	stocks stockStore
}

type cartQuery interface {
	ListCheckedItems(userID uint64) []cart.Item
}

type stockStore interface {
	HasEnough(productID uint64, quantity int) bool
	Deduct(productID uint64, quantity int)
}

func NewService(repo Repository, carts cartQuery, stocks stockStore) Service {
	return Service{repo: repo, carts: carts, stocks: stocks}
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

	if s.stocks != nil {
		for _, item := range checkedItems {
			if !s.stocks.HasEnough(item.ProductID, item.Quantity) {
				return Order{}, ErrInsufficientStock
			}
		}
		for _, item := range checkedItems {
			s.stocks.Deduct(item.ProductID, item.Quantity)
		}
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
