package cart

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMySQLRepositoryListByUserID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("expected sqlite db, got %v", err)
	}

	if err := db.Exec(`
		CREATE TABLE cart_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL,
			checked INTEGER NOT NULL,
			product_name TEXT NOT NULL,
			price INTEGER NOT NULL
		)
	`).Error; err != nil {
		t.Fatalf("expected schema creation, got %v", err)
	}

	if err := db.Exec(`
		INSERT INTO cart_items (user_id, product_id, quantity, checked, product_name, price)
		VALUES (1, 11, 2, 1, 'phone', 100), (1, 12, 1, 0, 'cable', 50), (2, 13, 1, 1, 'mouse', 80)
	`).Error; err != nil {
		t.Fatalf("expected seed insert, got %v", err)
	}

	repo := NewMySQLRepository(db)
	items := repo.ListByUserID(1)

	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}

	if items[0].UserID != 1 {
		t.Fatalf("expected user id 1, got %d", items[0].UserID)
	}
}


func TestMySQLRepositorySaveUpdatesExistingItem(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("expected sqlite db, got %v", err)
	}

	if err := db.Exec(`
		CREATE TABLE cart_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL,
			checked INTEGER NOT NULL,
			product_name TEXT NOT NULL,
			price INTEGER NOT NULL
		)
	`).Error; err != nil {
		t.Fatalf("expected schema creation, got %v", err)
	}

	repo := NewMySQLRepository(db)
	repo.Save(Item{UserID: 1, ProductID: 11, ProductName: "phone", Price: 100, Quantity: 1, Checked: true})
	repo.Save(Item{UserID: 1, ProductID: 11, ProductName: "phone", Price: 100, Quantity: 3, Checked: true})

	items := repo.ListByUserID(1)
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0].Quantity != 3 {
		t.Fatalf("expected quantity 3, got %d", items[0].Quantity)
	}
}
