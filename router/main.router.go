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

	// Category (sama polanya, jangan pakai handler global)
}
