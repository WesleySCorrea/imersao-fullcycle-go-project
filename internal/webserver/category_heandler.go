package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/model"
	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/service"
	"github.com/go-chi/chi"
)

type WebCategoryHeandler struct {
	CategoryService *service.CategoryService
}

func NewWebCategoryHeandler(categoryService service.CategoryService) *WebCategoryHeandler {
	return &WebCategoryHeandler{
		CategoryService: &categoryService,
	}
}

func (wch *WebCategoryHeandler) GetCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := wch.CategoryService.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func (wch *WebCategoryHeandler) GetCategory(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	category, err := wch.CategoryService.GetCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)

}

func (wch *WebCategoryHeandler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wch.CategoryService.CreateCategory(category.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)

}
