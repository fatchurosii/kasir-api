package main

import (
	"fmt"
	"kasir-api/entity"
	"kasir-api/router"
	"net/http"
)

var products = []entity.Product{
	{
		ID:    1,
		Name:  "Apple",
		Price: 1000,
		Stock: 10,
	},
	{
		ID:    2,
		Name:  "Banana",
		Price: 1000,
		Stock: 10,
	},
}

func main() {

	//http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//
	//	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	//	id, err := strconv.Atoi(idStr)
	//	if err != nil {
	//		http.Error(w, "Invalid request ID", http.StatusBadRequest)
	//	}
	//
	//	for _, product := range products {
	//		if product.ID == id {
	//			err := json.NewEncoder(w).Encode(product)
	//			if err != nil {
	//				return
	//			}
	//			return
	//		}
	//	}
	//
	//	http.Error(w, "Product not found", http.StatusNotFound)
	//
	//})
	//
	//http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//
	//	if r.Method == "GET" {
	//		err := json.NewEncoder(w).Encode(products)
	//		if err != nil {
	//			return
	//		}
	//	}
	//
	//	if r.Method == "POST" {
	//		var newProduct entity.Product
	//		err := json.NewDecoder(r.Body).Decode(&newProduct)
	//		if err != nil {
	//			http.Error(w, "Invalid request", http.StatusBadRequest)
	//		}
	//
	//		newProduct.ID = len(products) + 1
	//		products = append(products, newProduct)
	//
	//		w.WriteHeader(http.StatusCreated)
	//
	//		err = json.NewEncoder(w).Encode(newProduct)
	//
	//		if err != nil {
	//			return
	//		}
	//
	//	}
	//
	//})

	router.RegisterRoutes()

	fmt.Println("Running server on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server running failed")
	}

}
