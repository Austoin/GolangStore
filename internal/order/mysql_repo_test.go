package order

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMySQLRepositoryCreateAndGetByOrderNo(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("expected sqlite db, got %v", err)
	}

	if err := db.Exec(`
		CREATE TABLE orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			order_no TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			total_amount INTEGER NOT NULL,
			status INTEGER NOT NULL
		)
	`).Error; err != nil {
		t.Fatalf("expected orders schema, got %v", err)
	}

	if err := db.Exec(`
		CREATE TABLE order_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			order_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			product_name TEXT NOT NULL,
			price INTEGER NOT NULL,
			quantity INTEGER NOT NULL
		)
	`).Error; err != nil {
		t.Fatalf("expected order_items schema, got %v", err)
	}

	repo := NewMySQLRepository(db)
	created, err := repo.Create(Order{
		OrderNo:     "O9001",
		UserID:      9,
		TotalAmount: 250,
		Status:      StatusPending,
		Items: []Item{
			{ProductID: 1, ProductName: "phone", Price: 100, Quantity: 2},
			{ProductID: 2, ProductName: "cable", Price: 50, Quantity: 1},
		},
	})
	if err != nil {
		t.Fatalf("expected no create error, got %v", err)
	}

	loaded, err := repo.GetByOrderNo("O9001")
	if err != nil {
		t.Fatalf("expected no get error, got %v", err)
	}

	if created.OrderNo != loaded.OrderNo {
		t.Fatalf("expected order no match, got %s and %s", created.OrderNo, loaded.OrderNo)
	}

	if len(loaded.Items) != 2 {
		t.Fatalf("expected 2 order items, got %d", len(loaded.Items))
	}
}
