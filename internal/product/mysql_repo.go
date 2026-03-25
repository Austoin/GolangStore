package product

import "gorm.io/gorm"

type productRow struct {
	ID          uint64 `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Price       uint64 `gorm:"column:price"`
	Status      uint8  `gorm:"column:status"`
}

type stockRow struct {
	ProductID uint64 `gorm:"column:product_id"`
	Stock     int    `gorm:"column:stock"`
}

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) MySQLRepository {
	return MySQLRepository{db: db}
}

func (r MySQLRepository) GetByID(id uint64) (Product, error) {
	var prod productRow
	if err := r.db.Table("products").Where("id = ?", id).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return Product{}, ErrProductNotFound
		}
		return Product{}, err
	}

	var stock stockRow
	r.db.Table("product_stocks").Where("product_id = ?", id).First(&stock)

	return Product{
		ID:          prod.ID,
		Name:        prod.Name,
		Description: prod.Description,
		Price:       prod.Price,
		Status:      prod.Status,
		Stock:       stock.Stock,
	}, nil
}

func (r MySQLRepository) HasEnough(productID uint64, quantity int) bool {
	var stock stockRow
	if err := r.db.Table("product_stocks").Where("product_id = ?", productID).First(&stock).Error; err != nil {
		return false
	}

	return stock.Stock >= quantity
}

func (r MySQLRepository) Deduct(productID uint64, quantity int) {
	r.db.Table("product_stocks").Where("product_id = ?", productID).Update("stock", gorm.Expr("stock - ?", quantity))
}

func (r MySQLRepository) List() []Product {
	rows := make([]productRow, 0)
	r.db.Table("products").Find(&rows)
	items := make([]Product, 0, len(rows))
	for _, row := range rows {
		var stock stockRow
		r.db.Table("product_stocks").Where("product_id = ?", row.ID).First(&stock)
		items = append(items, Product{
			ID: row.ID,
			Name: row.Name,
			Description: row.Description,
			Price: row.Price,
			Status: row.Status,
			Stock: stock.Stock,
		})
	}

	return items
}
