package service

import (
	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/database"
	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/model"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (cs CategoryService) GetCategories() ([]*model.Category, error) {
	categories, err := cs.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs CategoryService) GetCategory(id string) (*model.Category, error) {
	category, err := cs.CategoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}

	return category, err
}

func (cs CategoryService) CreateCategory(name string) (*model.Category, error) {
	category := model.NewCategory(name)
	_, err := cs.CategoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
