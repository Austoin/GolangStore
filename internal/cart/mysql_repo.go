package cart

import "gorm.io/gorm"

type mysqlRow struct {
	UserID      uint64 `gorm:"column:user_id"`
	ProductID   uint64 `gorm:"column:product_id"`
	ProductName string `gorm:"column:product_name"`
	Price       uint64 `gorm:"column:price"`
	Quantity    int    `gorm:"column:quantity"`
	Checked     bool   `gorm:"column:checked"`
}

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) MySQLRepository {
	return MySQLRepository{db: db}
}

func (r MySQLRepository) ListByUserID(userID uint64) []Item {
	rows := make([]mysqlRow, 0)
	r.db.Table("cart_items").Where("user_id = ?", userID).Find(&rows)

	items := make([]Item, 0, len(rows))
	for _, row := range rows {
		items = append(items, Item{
			UserID:      row.UserID,
			ProductID:   row.ProductID,
			ProductName: row.ProductName,
			Price:       row.Price,
			Quantity:    row.Quantity,
			Checked:     row.Checked,
		})
	}

	return items
}

func (r MySQLRepository) Save(item Item) Item {
	row := mysqlRow{
		UserID:      item.UserID,
		ProductID:   item.ProductID,
		ProductName: item.ProductName,
		Price:       item.Price,
		Quantity:    item.Quantity,
		Checked:     item.Checked,
	}
	r.db.Table("cart_items").Create(&row)
	return item
}
