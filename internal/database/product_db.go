package database

import (
	"database/sql"

	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/model"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{
		db: db,
	}
}

func (cd *ProductDB) GetProducts() ([]*model.Product, error) {

	rows, err := cd.db.Query("SELECT id, name, description, price, category_id, image_url FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (cd *ProductDB) GetProduct(id string) (*model.Product, error) {
	var product model.Product
	err := cd.db.QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (cd *ProductDB) GetProductByCategoryID(categoryID string) ([]*model.Product, error) {
	rows, err := cd.db.Query("SELECT id, name, description, price, category_id, image_url FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (cd *ProductDB) CreateProduct(product *model.Product) (*model.Product, error) {

	_, err := cd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)",
		product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return nil, err
	}
	return product, nil
}
