package database

import (
	"app-example/internal/entity"
	pkg_entity "app-example/pkg/entity"
	"database/sql"
	"time"
)

type Product struct {
	db *sql.DB
}

func NewProduct(db *sql.DB) *Product {
	return &Product{db}
}

func (p *Product) Create(product *entity.Product) error {
	product.CreatedAt = time.Now()
	_, err := p.db.Exec("INSERT INTO products (id, name, price, created_at) VALUES (?, ?, ?, ?)", product.ID.String(), product.Name, product.Price, product.CreatedAt)
	return err
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if page < 0 {
		page = 0
	}

	rows, err := p.db.Query("SELECT id, name, price, created_at FROM products ORDER BY created_at "+sort+" LIMIT ? OFFSET ?", limit, page*limit)
	if err != nil {
		return nil, err
	}

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (p *Product) FindByID(id pkg_entity.ID) (*entity.Product, error) {
	var product entity.Product
	err := p.db.QueryRow("SELECT id, name, price, created_at FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.ID.String())
	return err
}

func (p *Product) Delete(id pkg_entity.ID) error {
	_, err := p.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

func CreateProductTables(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE products (id TEXT PRIMARY KEY, name TEXT, price INTEGER, created_at DATETIME)")
	return err
}
