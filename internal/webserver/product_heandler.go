package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/model"
	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/service"
	"github.com/go-chi/chi"
)

type WebProductHeandler struct {
	ProductService *service.ProductService
}

func NewWebProductHeandler(productService service.ProductService) *WebProductHeandler {
	return &WebProductHeandler{
		ProductService: &productService,
	}
}

func (wph *WebProductHeandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	products, err := wph.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHeandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	product, err := wph.ProductService.ProductDB.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)

}

func (wph *WebProductHeandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {

	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	products, err := wph.ProductService.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)

}

func (wph *WebProductHeandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wph.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)

}
