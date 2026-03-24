package product

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMySQLRepositoryGetByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("expected sqlite db, got %v", err)
	}

	if err := db.Exec(`
		CREATE TABLE products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			price INTEGER NOT NULL,
			status INTEGER NOT NULL
		)
	`).Error; err != nil {
		t.Fatalf("expected products schema, got %v", err)
	}

	if err := db.Exec(`
		CREATE TABLE product_stocks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			product_id INTEGER NOT NULL,
			stock INTEGER NOT NULL
		)
	`).Error; err != nil {
		t.Fatalf("expected product_stocks schema, got %v", err)
	}

	if err := db.Exec(`
		INSERT INTO products (id, name, description, price, status) VALUES (1, 'phone', 'smart phone', 199900, 1)
	`).Error; err != nil {
		t.Fatalf("expected product seed, got %v", err)
	}

	if err := db.Exec(`
		INSERT INTO product_stocks (product_id, stock) VALUES (1, 8)
	`).Error; err != nil {
		t.Fatalf("expected stock seed, got %v", err)
	}

	repo := NewMySQLRepository(db)
	item, err := repo.GetByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if item.Stock != 8 {
		t.Fatalf("expected stock 8, got %d", item.Stock)
	}
}
