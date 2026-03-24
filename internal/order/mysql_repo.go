package order

import "gorm.io/gorm"

type orderRow struct {
	ID          uint64 `gorm:"column:id"`
	OrderNo     string `gorm:"column:order_no"`
	UserID      uint64 `gorm:"column:user_id"`
	TotalAmount uint64 `gorm:"column:total_amount"`
	Status      uint8  `gorm:"column:status"`
}

type orderItemRow struct {
	OrderID     uint64 `gorm:"column:order_id"`
	ProductID   uint64 `gorm:"column:product_id"`
	ProductName string `gorm:"column:product_name"`
	Price       uint64 `gorm:"column:price"`
	Quantity    int    `gorm:"column:quantity"`
}

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) MySQLRepository {
	return MySQLRepository{db: db}
}

func (r MySQLRepository) GetByOrderNo(orderNo string) (Order, error) {
	var row orderRow
	if err := r.db.Table("orders").Where("order_no = ?", orderNo).First(&row).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return Order{}, ErrOrderNotFound
		}
		return Order{}, err
	}

	itemRows := make([]orderItemRow, 0)
	r.db.Table("order_items").Where("order_id = ?", row.ID).Find(&itemRows)

	items := make([]Item, 0, len(itemRows))
	for _, item := range itemRows {
		items = append(items, Item{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}

	return Order{
		ID:          row.ID,
		OrderNo:     row.OrderNo,
		UserID:      row.UserID,
		TotalAmount: row.TotalAmount,
		Status:      row.Status,
		Items:       items,
	}, nil
}

func (r MySQLRepository) Create(order Order) (Order, error) {
	row := orderRow{
		OrderNo:     order.OrderNo,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
	}

	if err := r.db.Table("orders").Create(&row).Error; err != nil {
		return Order{}, err
	}

	for _, item := range order.Items {
		itemRow := orderItemRow{
			OrderID:     row.ID,
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		}
		if err := r.db.Table("order_items").Create(&itemRow).Error; err != nil {
			return Order{}, err
		}
	}

	order.ID = row.ID
	return order, nil
}
