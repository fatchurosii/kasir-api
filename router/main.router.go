package router

import (
	"database/sql"
	"kasir-api/handler"
	"kasir-api/repository"
	"kasir-api/service"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {

	http.HandleFunc("/health", handler.HealthHandler)

	// Product
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetAllProducts(w, r)
		case http.MethodPost:
			productHandler.StoreProduct(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetProductById(w, r)
		case http.MethodPut:
			productHandler.UpdateProduct(w, r)
		case http.MethodDelete:
			productHandler.DeleteProduct(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetAllCategories(w, r)
		case http.MethodPost:
			categoryHandler.StoreCategory(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	http.HandleFunc("/api/category/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetCategoryById(w, r)
		case http.MethodPut:
			categoryHandler.UpdateCategory(w, r)
		case http.MethodDelete:
			categoryHandler.DeleteCategory(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

}
