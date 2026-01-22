package router

import (
	"kasir-api/handler"
	"net/http"
)

func RegisterRoutes() {

	http.HandleFunc("/health", handler.HealthHandler)

	//Product

	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetAllProducts(w, r)
		case http.MethodPost:
			handler.StoreProduct(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetProductById(w, r)
		case http.MethodPut:
			handler.UpdateProduct(w, r)
		case http.MethodDelete:
			handler.DeleteProduct(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	//end product

	// Category

	http.HandleFunc("/api/category", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetAllCategory(w, r)
		case http.MethodPost:
			handler.StoreCategory(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	http.HandleFunc("/api/category/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCategoryById(w, r)
		case http.MethodPut:
			handler.UpdateCategory(w, r)
		case http.MethodDelete:
			handler.DeleteCategory(w, r)
		default:
			handler.Error(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	// End Category
}
