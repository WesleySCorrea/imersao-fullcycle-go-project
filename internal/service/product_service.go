package service

import (
	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/database"
	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/model"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: productDB,
	}
}

func (ps ProductService) GetProducts() ([]*model.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps ProductService) GetProduct(id string) (*model.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return product, err
}

func (ps ProductService) GetProductByCategoryID(category_id string) ([]*model.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(category_id)
	if err != nil {
		return nil, err
	}

	return products, err
}

func (ps ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*model.Product, error) {
	product := model.NewProduct(name, description, category_id, image_url, price)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
