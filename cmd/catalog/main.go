package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/database"
	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/service"
	"github.com/WesleySCorrea/imersao-fullcycle-go-project/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersaofullcycle")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	WebCategoryHeandler := webserver.NewWebCategoryHeandler(*categoryService)
	WebProductHeandler := webserver.NewWebProductHeandler(*productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category/{id}", WebCategoryHeandler.GetCategory)
	c.Get("/category", WebCategoryHeandler.GetCategories)
	c.Post("/category", WebCategoryHeandler.CreateCategory)

	c.Get("/product/{id}", WebProductHeandler.GetProduct)
	c.Get("/product/category/{categoryID}", WebProductHeandler.GetProductByCategoryID)
	c.Get("/products", WebProductHeandler.GetProducts)
	c.Post("/product", WebProductHeandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
