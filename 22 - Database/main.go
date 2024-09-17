package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/sys")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// create database test;
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS gotest")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("USE gotest")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id VARCHAR(36) PRIMARY KEY, name VARCHAR(255), price DECIMAL(10,2))")
	if err != nil {
		panic(err)
	}

	product := NewProduct("Product 1", 9.99)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product2, err := getProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	println(product2.ID, product2.Name, product2.Price)
	product2.Name = "Product 2"
	updateProduct(db, product2)

	product3, err := getProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	println(product3.ID, product3.Name, product3.Price)

	products, err := getProducts(db)
	if err != nil {
		panic(err)
	}
	println("Products:")
	for _, product := range products {
		println(product.ID, product.Name, product.Price)
	}

}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	return err
}

func getProduct(db *sql.DB, id string) (*Product, error) {
	row := db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id)

	product := &Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func getProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*Product, 0)
	for rows.Next() {
		product := &Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	return err
}
