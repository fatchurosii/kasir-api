package handler

import (
	"kasir-api/entity"
	"kasir-api/helper"
	"kasir-api/service"
	"net/http"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/product/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	product, err := service.GetProductByID(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	Success(w, http.StatusOK, "Product found", product)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	Success(w, http.StatusOK, "All products", service.GetAllProduct())
}

func StoreProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	if err := helper.DecodeJSON(r, &product); err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	result := service.CreateProduct(product)
	Success(w, http.StatusCreated, "Product created", result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/product/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	var product entity.Product
	if err := helper.DecodeJSON(r, &product); err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	result, err := service.UpdateProduct(id, product)
	if err != nil {
		Error(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	Success(w, http.StatusOK, "Product updated", result)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/product/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	product, err := service.DeleteProduct(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	Success(w, http.StatusOK, "Product deleted", product)
}
