package product

import (
	"fmt"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMySQLRepositoryGetByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:product_get_%s?mode=memory&cache=shared", t.Name())), &gorm.Config{})
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

func TestMySQLRepositoryDeductStock(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:product_stock_%s?mode=memory&cache=shared", t.Name())), &gorm.Config{})
	if err != nil {
		t.Fatalf("expected sqlite db, got %v", err)
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
		INSERT INTO product_stocks (product_id, stock) VALUES (1, 8)
	`).Error; err != nil {
		t.Fatalf("expected stock seed, got %v", err)
	}

	repo := NewMySQLRepository(db)
	if !repo.HasEnough(1, 5) {
		t.Fatal("expected enough stock")
	}

	repo.Deduct(1, 3)
	item, err := repo.GetByID(1)
	if err == nil {
		_ = item
	}

	var stock int
	if err := db.Table("product_stocks").Select("stock").Where("product_id = ?", 1).Scan(&stock).Error; err != nil {
		t.Fatalf("expected stock query, got %v", err)
	}

	if stock != 5 {
		t.Fatalf("expected stock 5, got %d", stock)
	}
}


func TestMySQLRepositoryList(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:product_list_%s?mode=memory&cache=shared", t.Name())), &gorm.Config{})
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
		INSERT INTO products (id, name, description, price, status) VALUES
		(1, 'phone', 'smart phone', 199900, 1),
		(2, 'cable', 'fast cable', 4900, 1)
	`).Error; err != nil {
		t.Fatalf("expected products seed, got %v", err)
	}

	if err := db.Exec(`
		INSERT INTO product_stocks (product_id, stock) VALUES (1, 8), (2, 25)
	`).Error; err != nil {
		t.Fatalf("expected stock seed, got %v", err)
	}

	repo := NewMySQLRepository(db)
	items := repo.List()

	if len(items) != 2 {
		t.Fatalf("expected 2 products, got %d", len(items))
	}
}
